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

func FetchUser(id int32) (*pb.User, error) {
	// check if the user is present in in-memory DB
	user, exists := USER_MAP[int(id)]
	if exists {
		user := &pb.User{
			Id:      id,
			Fname:   user.fname,
			City:    user.city,
			Phone:   user.phone,
			Height:  user.height,
			Married: user.Married,
		}
		return user, nil
	}
	// return with error if user is not found
	return nil, http_error.Error(404, "user not found")
}

func (server *UsermanagerServer) GetUser(ctx context.Context, input *pb.GetUserInput) (*pb.User, error) {
	log.Printf("Received request for user id: %v", input.GetId())
	return FetchUser(input.GetId())
}

func (server *UsermanagerServer) GetUsers(ctx context.Context, input *pb.GetUsersInput) (*pb.Users, error) {
	ids := input.GetIds()
	log.Printf("Received request for user ids: %v", ids)
	var users []*pb.User
	// validate input
	if len(ids) == 0 {
		return nil, http_error.Error(400, "bad request: ids not provided")
	}
	// traverse through ids passed in request
	// and fetch user details from in-memory DB
	for _, id := range ids {
		user, http_error := FetchUser(id)
		if http_error == nil {
			users = append(users, user)
		}
	}
	response := &pb.Users{Users: users}
	return response, nil
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
	log.Printf("server listening at: %v", listener.Addr())
	// log error, if any
	if err := server.Serve(listener); err != nil {
		log.Fatalf("failed to serve with error: %v", err)
	}
}
