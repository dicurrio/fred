// fred/main.go

package main

import (
	"log"
	"net"

	pb "github.com/dicurrio/protorepo/fred"
	"google.golang.org/grpc"
)

const port = ":50051"

func main() {
	log.SetPrefix("FRED :: ")

	log.Print("Starting up...")
	listener, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	} else {
		log.Printf("Listening on %v", port)
	}

	s := grpc.NewServer()
	pb.RegisterFredServer(s, &fredServer{})

	if err := s.Serve(listener); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
