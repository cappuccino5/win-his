package main

import (
	"flag"
	"fmt"
	"net/http"

	"context"

	"github.com/golang/glog"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"google.golang.org/grpc"
	hello "win-his/gateway"
)

var (
	// the go.micro.srv.greeter address
	endpoint = flag.String("endpoint", "localhost:9000", "http.SagaAdmin")
)

func run() error {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithInsecure()}
	err := hello.RegisterSagaAdminHandlerFromEndpoint(ctx, mux, *endpoint, opts)
	if err != nil {
		return err
	}
	return http.ListenAndServe(":8888", mux)
}

func main() {
	flag.Parse()

	defer glog.Flush()
	fmt.Println("welcome to http gateway..")
	if err := run(); err != nil {
		glog.Fatal(err)
	}
}
