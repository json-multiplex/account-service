bin/accounts:
	go build -o bin/accounts cmd/accounts/main.go

generated/v0: v0/*
	protoc \
		--proto_path=proto_ext \
		--proto_path=v0 \
		--go_out=plugins=grpc:generated/v0 \
		v0/accounts.proto

	protoc \
		--proto_path=proto_ext \
		--proto_path=v0 \
		--grpc-gateway_out=logtostderr=true:generated/v0 \
		v0/accounts.proto
