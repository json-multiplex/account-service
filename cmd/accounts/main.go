package main

import (
	"context"
	"log"
	"net"
	"net/http"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	pb "github.com/json-multiplex/account-service/generated/v0"
)

type server struct{}

func (s server) GetAccount(ctx context.Context, req *pb.GetAccountRequest) (*pb.Account, error) {
	return &pb.Account{
		Name: "accounts/123",
	}, nil
}

func main() {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	grpcServer := grpc.NewServer()
	pb.RegisterAccountsServer(grpcServer, &server{})
	reflection.Register(grpcServer)

	l, err := net.Listen("tcp", ":3000")
	if err != nil {
		log.Fatalf("failed to listen and serve grpc: %v", err)
	}

	go grpcServer.Serve(l)

	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithInsecure()}
	pb.RegisterAccountsHandlerFromEndpoint(ctx, mux, ":3000", opts)

	if err := http.ListenAndServe(":4000", mux); err != nil {
		log.Fatalf("failed to listen and serve http: %v", err)
	}
}
