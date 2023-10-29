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

type MultipleUserRequest struct {
	Ids []int32
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
	// fetch user details
	users := []UserRequest{
		{Id: 1},
		{Id: 2},
		{Id: 3},
	}
	for _, user := range users {
		res, err := client.GetUser(ctx, &pb.GetUserInput{
			Id: int32(user.Id),
		})
		if err != nil {
			log.Printf("could not find user: %v", err)
		} else {
			log.Printf(`
				id : %d
				fname : %s
				city : %s
				phone : %v,
			`, res.GetId(), res.GetFname(), res.GetCity(), res.GetPhone())
		}
	}
	// fetch multiple users
	multiple_users := []MultipleUserRequest{
		{Ids: []int32{1, 2}},
		{Ids: []int32{1, 3}},
	}
	for _, users := range multiple_users {
		res, err := client.GetUsers(ctx, &pb.GetUsersInput{
			Ids: users.Ids,
		})
		if err != nil {
			log.Printf("could not find user: %v", err)
		} else {
			log.Println(res)
			for _, user := range res.GetUsers() {
				log.Printf(`
					id : %d
					fname : %s
					city : %s
					phone : %v,
				`, user.GetId(), user.GetFname(), user.GetCity(), user.GetPhone())
			}
		}
	}
}
