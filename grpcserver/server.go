package grpcserver

import (
	"fmt"
	"log"
	"net"
	"quickkv/quickkvpb"

	"google.golang.org/grpc"
)

type server struct{}

// StartServer - start grpc server
func StartServer(port uint16) {
	address := fmt.Sprintf("0.0.0.0:%d", port)
	lis, err := net.Listen("tcp", address)
	if err != nil {
		log.Panic(err.Error())
	}

	grpcserver := grpc.NewServer()
	quickkvpb.RegisterIoServiceServer(grpcserver, &server{})
	quickkvpb.RegisterStoreServiceServer(grpcserver, &server{})

	fmt.Printf("grpc Server Listening on %v ... \n", address)
	grpcserver.Serve(lis)
}
