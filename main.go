// fred/main.go

package main

import (
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"

	pb "github.com/dicurrio/protorepo/fred"
	"google.golang.org/grpc"
)

var hostAddress = os.Getenv("HOST_ADDRESS")

func main() {
	// Setup
	log.SetPrefix("FRED :: ")
	log.Print("Starting up...")

	// Establish TCP Listener
	listener, err := net.Listen("tcp", hostAddress)
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	} else {
		log.Printf("Listening on %v", hostAddress)
	}

	// Start gRPC Server
	server := grpc.NewServer()
	pb.RegisterFredServer(server, &fredServer{})
	go func() {
		log.Print("Listening on " + hostAddress)
		log.Fatal(server.Serve(listener))
	}()

	// Graceful Shutdown
	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, syscall.SIGINT, syscall.SIGTERM)
	<-signalChan // Blocks until SIGINT or SIGTERM received
	log.Print("Shutdown signal received, exiting...")
	server.GracefulStop()
}
