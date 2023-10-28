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

type UsermanagerTask struct {
	Name        string
	Description string
	Done        bool
}

func main() {
	conn, err := grpc.Dial(ADDRESS, grpc.WithInsecure(), grpc.WithBlock())

	if err != nil {
		log.Fatalf("did not connect : %v", err)
	}

	defer conn.Close()

	c := pb.NewUsermanagerServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)

	defer cancel()

	todos := []UsermanagerTask{
		{Name: "Code review", Description: "Review new feature code", Done: false},
		{Name: "Make YouTube Video", Description: "Start Go for beginners series", Done: false},
		{Name: "Go to the gym", Description: "Leg day", Done: false},
		{Name: "Buy groceries", Description: "Buy tomatoes, onions, mangos", Done: false},
		{Name: "Meet with mentor", Description: "Discuss blockers in my project", Done: false},
	}

	for _, todo := range todos {
		res, err := c.GetUser(ctx, &pb.Input{
			UserId: todo.Name,
		})

		if err != nil {
			log.Fatalf("could not create user: %v", err)
		}

		log.Printf(`
           ID : %s
           Name : %s
           Description : %s
           Done : %v,
       `, res.GetUserId(), res.GetName(), res.GetDescription(), res.GetDone())
	}

}
