package main

import (
	"context"
	"demo-grpc/invoicer"
	"fmt"
	"net"

	"google.golang.org/grpc"
)

type myInvoicerServer struct {
	invoicer.UnimplementedInvoicerServer
}

func (s *myInvoicerServer) Create(ctx context.Context, req *invoicer.CreateRequest) (*invoicer.CreateResponse, error) {
	fmt.Println("Create invoicer")

	return &invoicer.CreateResponse{
		Pdf:  []byte("Hello World"),
		Docx: []byte("Hello World"),
	}, nil
}

func main() {
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		panic(err)
	}

	serverRegistrar := grpc.NewServer()
	service := &myInvoicerServer{}

	invoicer.RegisterInvoicerServer(serverRegistrar, service)
	err = serverRegistrar.Serve(listener)
	if err != nil {
		panic(err)
	}
}
