package main

import (
	"context"
	"google.golang.org/grpc"
	"io"
	"log"
	pb "mod/assignment/proto"
)

func main() {
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewUserServiceClient(conn)

	user, err := c.AddUser(context.Background(), &pb.User{Name: "Alice", Email: "alice@example.com"})
	if err != nil {
		log.Fatalf("could not add user: %v", err)
	}
	log.Printf("Added user ID: %d", user.Id)

	gotUser, err := c.GetUser(context.Background(), &pb.User{Id: user.Id})
	if err != nil {
		log.Fatalf("could not get user: %v", err)
	}
	log.Printf("Got user: %s", gotUser.Name)

	stream, err := c.ListUsers(context.Background(), &pb.Empty{})
	if err != nil {
		log.Fatalf("could not list users: %v", err)
	}
	for {
		user, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("list users error: %v", err)
		}
		log.Printf("User: %s", user.Name)
	}
}
