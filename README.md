# rLauncher - Job Scheduler over gRPC ![Go](https://github.com/jbltx/rlauncher/workflows/Go/badge.svg)

> This is a work in progress

## Features

 - [ ] Manage Jobs, Tasks and Agents from the built-in TUI Monitor
 - [ ] Easy to setup with simple command prompts
 - [ ] Agent discovery using [Zeroconf](https://fr.wikipedia.org/wiki/Zeroconf)
 - [ ] Scheduler ⇄ Agent communication with [gRPC](https://grpc.io/) and [protocol buffers](https://developers.google.com/protocol-buffers)
 - [ ] Spread a single Job in several smart Tasks
 - [ ] Group similar agents using Pools
 - [ ] Create Pools based on relational sets of others Pools
 - [ ] User authentication via [Google Sign-In](https://developers.google.com/identity)
 - [ ] Support of SQLite/MySQL/PostgreSQL as persistent storage
 - [ ] REST API to manage remotely your farm or create a web-based clients
 - [ ] Agent running as a service on slave machines
 - [ ] All included in one executable : Scheduler, Agent, Service and TUI Monitor
  
#### Future features

- [ ] Webhooks
- [ ] [MongoDB](https://www.mongodb.com/) support
- [ ] [RBAC](https://en.wikipedia.org/wiki/Role-based_access_control) for Users
- [ ] Others authentication providers (LDAP/ActiveDirectory)

## How it works

## Installation

You can pick the latest release for each supported platforms (Windows/MacOS/Linux)
on the [Github's releases page](https://github.com/jbltx/rlauncher/releases).

> As the application is using SSH protocol to communicate between Scheduler and Agents, please make sure to generate and share **RSA** **certificates** across your nodes for security purposes. You can use `ssh-keygen -t rsa` command for this.

 1. Install the agent service on all slave nodes you want to use to run processes.
    ```sh
    sudo rLauncher agent install
    sudo rLauncher agent start
    ```
 2. On the master node, you run the application without options
    ```sh
    rLauncher
    ```
    This will launch the terminal-based user-interface.

### Available *Scheduler* commands
 * **list** : Print a list of available agents from discovery.

### Available *Agent* commands

 * **install** : Install the Agent OS-specific service.
 * **uninstall** : Uninstall the service.
 * **start** : Start the service, if needed.
 * **stop** : Stop the service, if needed.
 * **run** : Special command to run the agent directly without installing
 a service. Please be aware it will cause conflict if the service is listening in background on same ports as the run command.
   * `--no-discovery` : will run without discovery services.
   * `--port <p>` : define a port to listen to (default: 22000).


## Get the code for Go projects

```sh
go get github.com/jbltx/rlauncher
```

## Feedback

Add your issue here on GitHub. Feel free to get in touch if you have any questions.


## Contribution

[TODO]

## Dependencies

##### Configuration
* [viper](https://github.com/spf13/viper)
* [cobra](https://github.com/spf13/cobra)

##### OS (Services, Filesystem...)
* [afero](https://github.com/spf13/afero)
* [service](https://github.com/kardianos/service)

##### Communication
* [crypto/ssh](https://golang.org/x/crypto/ssh)
* [zeroconf](https://github.com/grandcat/zeroconf)

##### Database
* [gorm](https://github.com/jinzhu/gorm)
* [go.uuid](https://github.com/satori/go.uuid)

##### TUI
* [tcell](https://github.com/gdamore/tcell)
* [tview](https://github.com/rivo/tview)
* [termui](https://github.com/gizak/termui)
* [promptui](https://github.com/manifoldco/promptui)

##### Security
* [auth](https://github.com/qor/auth)