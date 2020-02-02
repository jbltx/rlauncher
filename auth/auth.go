package auth

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"os/exec"
	"strconv"

	"github.com/jbltx/rlauncher/auth/provider"
	"github.com/jbltx/rlauncher/cfg"
	user "github.com/jbltx/rlauncher/user/model"
	"golang.org/x/oauth2"
)

type authProvider interface {
	OAuth2Config() *oauth2.Config
	UserInfoURL() string
	ParseUserInfo(content []byte) (*user.User, error)
}

// Service ...
type Service struct {
	authenticatedUser *user.User
	provider          authProvider
	server            *http.Server
	ctx               context.Context
}

// NewService ...
func NewService(appConfig *cfg.Config) *Service {

	var chosenProvider authProvider

	switch appConfig.Auth.Provider {
	case cfg.GoogleAP:
		chosenProvider = provider.NewGoogleProvider()
	default:
		panic(fmt.Sprintf("Invalid auth provider: '%s'", appConfig.Auth.Provider))
	}

	return &Service{
		authenticatedUser: nil,
		provider:          chosenProvider,
	}
}

// User ...
func (svc *Service) User() *user.User {
	return svc.authenticatedUser
}

// State ...
func (svc *Service) State() string {
	state, _ := os.Hostname()
	state += "-" + strconv.Itoa(os.Getpid())
	return state
}

func (svc *Service) callbackHandler(w http.ResponseWriter, r *http.Request) {

	svc.ctx = context.Background()
	defer svc.server.Shutdown(svc.ctx)

	code := r.FormValue("code")
	state := r.FormValue("state")
	if state != svc.State() {
		fmt.Fprintf(w, "Invalid request, state is unknown\n")
		return
	}
	token, err := svc.provider.OAuth2Config().Exchange(svc.ctx, code)
	if err != nil {
		fmt.Fprintf(w, "Unable to exchange code to access token\n")
		return
	}
	res, err := svc.provider.OAuth2Config().Client(svc.ctx, token).Get(svc.provider.UserInfoURL())
	if err != nil {
		fmt.Fprintf(w, "Unable to process user info request\n")
		return
	}
	defer res.Body.Close()
	content, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Fprintf(w, "Unable to fetch response body from user info request\n")
		return
	}
	svc.authenticatedUser, err = svc.provider.ParseUserInfo(content)
	if err != nil {
		fmt.Fprintf(w, "Unable to parse response body from user info request\n")
		return
	}
	fmt.Fprintf(w, "User Data : %s\n", content)
}

// Login ...
func (svc *Service) Login() {
	// launch Sign-In web page
	url := svc.provider.OAuth2Config().AuthCodeURL(svc.State(), oauth2.AccessTypeOffline)
	exec.Command(url).Start()

	// run a server to handle response
	mux := http.NewServeMux()
	mux.HandleFunc("/auth/callback", svc.callbackHandler)
	svc.server = &http.Server{
		Addr:    fmt.Sprintf(":55789"),
		Handler: mux,
	}
	svc.server.ListenAndServe()
}
