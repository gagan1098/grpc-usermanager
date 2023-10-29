package main

import (
	"context"
	"log"
	"net"

	pb "github.com/grpc-usermanager/proto"
	"google.golang.org/grpc"
	http_error "google.golang.org/grpc/status"
)

const (
	// Port for gRPC server to listen to
	PORT = ":50051"
)

type UsermanagerServer struct {
	pb.UnimplementedUsermanagerServiceServer
}

type User struct {
	fname   string
	city    string
	phone   int64
	height  float32
	Married bool
}

var USER_MAP = map[int]*User{
	1: {"Steve", "LA", 1234567890, 5.8, true},
	2: {"Gagan", "New Delhi", 9876543210, 5.9, false},
}

func (s *UsermanagerServer) GetUser(ctx context.Context, input *pb.Input) (*pb.Output, error) {
	log.Printf("Received request for user id: %v", input.GetId())
	// check if the user is present in in-memory DB
	user, exists := USER_MAP[int(input.GetId())]
	if exists {
		user := &pb.Output{
			Id:      input.GetId(),
			Fname:   user.fname,
			City:    user.city,
			Phone:   user.phone,
			Height:  user.height,
			Married: user.Married,
		}
		return user, nil
	}
	// return with error if user is not found
	return nil, http_error.Error(400, "user not found")

}

func main() {
	// start listening on port
	listener, err := net.Listen("tcp", PORT)
	if err != nil {
		log.Fatalf("failed connection: %v", err)
	}
	server := grpc.NewServer()

	// register usermanager service
	pb.RegisterUsermanagerServiceServer(server, &UsermanagerServer{})
	log.Printf("server listening at - %v", listener.Addr())

	// log error, if any
	if err := server.Serve(listener); err != nil {
		log.Fatalf("failed to connect to server: %v", err)
	}
}
