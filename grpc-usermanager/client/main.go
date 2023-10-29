package main

import (
	"context"
	"log"
	"time"

	pb "github.com/grpc-usermanager/proto"

	"google.golang.org/grpc"
)

const (
	ADDRESS = "localhost:50051"
)

type UserRequest struct {
	Id int
}

func main() {
	// create connection
	conn, err := grpc.Dial(ADDRESS, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect : %v", err)
	}
	defer conn.Close()

	// create client
	client := pb.NewUsermanagerServiceClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	users := []UserRequest{
		{Id: 1},
		{Id: 2},
		{Id: 3},
	}

	for _, user := range users {
		res, err := client.GetUser(ctx, &pb.Input{
			Id: int32(user.Id),
		})

		if err != nil {
			log.Fatalf("could not find user: %v", err)
		}

		log.Printf(`
			id : %d
			fname : %s
			city : %s
			phone : %v,
		`, res.GetId(), res.GetFname(), res.GetCity(), res.GetPhone())
	}
}
