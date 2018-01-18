// fred/server.go

package main

import (
	"log"

	"golang.org/x/net/context"

	pb "github.com/dicurrio/protorepo/fred"
)

type fredServer struct {
}

func (fred *fredServer) GetIndex(ctx context.Context, req *pb.Request) (*pb.Response, error) {
	log.Printf("GetIndex %v", req)
	return &pb.Response{Message: "Hello to" + req.GetName() + " from Fred"}, nil
}
