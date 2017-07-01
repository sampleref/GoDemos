package main

import (
	"HelloDemo/addition"
	"fmt"
	"os"
	"path"
	"path/filepath"
	"HelloDemo/HelloServiceRPC"
	"google.golang.org/grpc"
	"net"
	"strconv"
	"google.golang.org/grpc/reflection"
	"HelloDemo/proto"
)

func main() {
	printExecutablePath()
	result, err := addition.Add(10, 20)
	if err != nil {
		fmt.Println("Invalid result %v", err)
	} else {
		fmt.Println("Added to ", result)
	}
	startGrpcServices()
}

func startGrpcServices() {

	fmt.Println("Starting grpc services")
	helloService := HelloServiceRPC.HelloServiceServer{NamePrefix:"Mr/Mrs"}
	addr := ":" + strconv.Itoa(7008)
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		fmt.Println("Failed to listen %v", err)
	}
	//Start gRPC server
	server := grpc.NewServer()
	service_Hello.RegisterHelloServiceServer(server, helloService)
	reflection.Register(server)

	if err := server.Serve(lis); err != nil {
		fmt.Println("Failed to start gRPC server %v", err)
	}
	fmt.Println("Server listening on the port", addr)

}

func printExecutablePath() {
	ex, err := os.Executable()
	if err != nil {
		fmt.Println("Error in printExecutablePath %v", err)
	}
	dir := path.Dir(ex)
	fmt.Println("Executable at ", dir)
	path, exception := filepath.Abs("./")
	if exception != nil {
		fmt.Println("Exception in printExecutablePath %v", exception)
	}
	fmt.Println("Executable at ", path)
	addition.PrintNoOfCores()
}
