// fred/main.go

package main

import (
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"

	"google.golang.org/grpc/credentials"

	pb "github.com/dicurrio/protorepo/fred"
	"google.golang.org/grpc"
)

const (
	crt         = "./tls/fred-cert.pem"
	key         = "./tls/fred-key.pem"
	hostAddress = "localhost:3001"
)

// var hostAddress = os.Getenv("HOST_ADDRESS")

func main() {
	// Setup
	log.SetPrefix("FRED :: ")
	log.Print("Starting up...")

	// Establish TCP Listener
	listener, err := net.Listen("tcp", hostAddress)
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	} else {
		log.Printf("Listening on %v", listener.Addr().String())
	}

	// Create TLS credentials
	creds, err := credentials.NewServerTLSFromFile(crt, key)
	if err != nil {
		log.Fatalf("Failed to load TLS files: %v", err)
	}

	// Start gRPC Server
	server := grpc.NewServer(grpc.Creds(creds))
	pb.RegisterFredServer(server, &fredServer{})
	go func() {
		log.Print("Serving...")
		log.Fatal(server.Serve(listener))
	}()

	// Graceful Shutdown
	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, syscall.SIGINT, syscall.SIGTERM)
	<-signalChan // Blocks until SIGINT or SIGTERM received
	log.Print("Shutdown signal received, exiting...")
	server.GracefulStop()
}
