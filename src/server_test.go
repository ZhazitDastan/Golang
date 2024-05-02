package main

import (
	"context"
	"log"
	"net"
	"testing"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/grpc/test/bufconn"
	pb "mod/assignment/proto"
)

const bufSize = 1024 * 1024

var lis *bufconn.Listener

func init() {
	lis = bufconn.Listen(bufSize)
	s := grpc.NewServer()
	pb.RegisterUserServiceServer(s, &server{})

	go func() {
		if err := s.Serve(lis); err != nil {
			log.Fatalf("Server exited with error: %v", err)
		}
	}()
}

func bufDialer(context.Context, string) (net.Conn, error) {
	return lis.Dial()
}

func TestAddAndGetUser(t *testing.T) {
	ctx := context.Background()
	conn, err := grpc.DialContext(ctx, "bufnet", grpc.WithContextDialer(bufDialer), grpc.WithInsecure())
	if err != nil {
		t.Fatalf("Failed to dial bufnet: %v", err)
	}
	defer conn.Close()

	client := pb.NewUserServiceClient(conn)

	newUser, err := client.AddUser(ctx, &pb.User{Name: "John Doe", Email: "john@example.com"})
	if err != nil {
		t.Errorf("AddUser failed: %v", err)
	}

	gotUser, err := client.GetUser(ctx, &pb.User{Id: newUser.Id})
	if err != nil {
		t.Errorf("GetUser failed: %v", err)
	}
	if gotUser.Name != "John Doe" || gotUser.Email != "john@example.com" {
		t.Errorf("GetUser returned wrong user: got %v, want %v", gotUser, newUser)
	}
}

func TestUserNotFound(t *testing.T) {
	ctx := context.Background()
	conn, err := grpc.DialContext(ctx, "bufnet", grpc.WithContextDialer(bufDialer), grpc.WithInsecure())
	if err != nil {
		t.Fatalf("Failed to dial bufnet: %v", err)
	}
	defer conn.Close()

	client := pb.NewUserServiceClient(conn)

	_, err = client.GetUser(ctx, &pb.User{Id: 999}) // Assuming 999 is an invalid ID.
	if err == nil {
		t.Error("Expected error for user not found, got none")
	}
	if st, ok := status.FromError(err); !ok || st.Code() != codes.NotFound {
		t.Errorf("Expected NotFound error for user not found, got %v", err)
	}
}
