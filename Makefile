proto:
	protoc user/delivery/user.proto -I. --go_out=plugins=grpc:./../../..
	protoc agent/delivery/agent.proto -I. --go_out=plugins=grpc:./../../..