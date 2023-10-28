package main

import (
	"context"
	"log"
	"net"

	"github.com/google/uuid"
	pb "github.com/grpc-usermanager/proto"
	"google.golang.org/grpc"
)

const (
	// Port for gRPC server to listen to
	PORT = ":50051"
)

type UsermanagerServer struct {
	pb.UnimplementedUsermanagerServiceServer
}

func (s *UsermanagerServer) GetUser(ctx context.Context, in *pb.Input) (*pb.Output, error) {
	log.Printf("Received: %v", in.GetUserId())
	todo := &pb.Output{
		UserId:      in.GetUserId(),
		Description: "desc",
		Done:        false,
		Name:        uuid.New().String(),
	}

	return todo, nil

}

func main() {
	lis, err := net.Listen("tcp", PORT)

	if err != nil {
		log.Fatalf("failed connection: %v", err)
	}

	s := grpc.NewServer()

	pb.RegisterUsermanagerServiceServer(s, &UsermanagerServer{})

	log.Printf("server listening at %v", lis.Addr())

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to server: %v", err)
	}
}
