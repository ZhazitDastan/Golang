package main

import (
	"context"
	"log"
	"net"
	"sync"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	pb "mod/assignment/proto"
)

type server struct {
	pb.UnimplementedUserServiceServer
	users []*pb.User
	mu    sync.Mutex
}

func (s *server) AddUser(ctx context.Context, user *pb.User) (*pb.User, error) {
	s.mu.Lock()
	user.Id = int32(len(s.users) + 1)
	s.users = append(s.users, user)
	s.mu.Unlock()
	return user, nil
}

func (s *server) GetUser(ctx context.Context, user *pb.User) (*pb.User, error) {
	s.mu.Lock()
	defer s.mu.Unlock()
	for _, u := range s.users {
		if u.Id == user.Id {
			return u, nil
		}
	}
	return nil, status.Errorf(codes.NotFound, "User not found")
}

func (s *server) ListUsers(_ *pb.Empty, stream pb.UserService_ListUsersServer) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	for _, user := range s.users {
		if err := stream.Send(user); err != nil {
			return err
		}
	}
	return nil
}

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterUserServiceServer(s, &server{})
	log.Println("Server listening at", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
