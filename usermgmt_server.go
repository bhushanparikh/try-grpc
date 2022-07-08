package main

import (
	"context"
	"log"
	"math/rand"
	"net"

	pb "example.com/go-usermgmt-grpc/usermgmt"
	"google.golang.org/grpc"
)

const (
	port = ":50001"
)

type UserManagementServer struct {
	pb.UnimplementedUserManagementServer
}

func (s *UserManagementServer) CreateUser(ctx context.Context, in *pb.UserRequest) (*pb.UserResponse, error) {
	log.Printf("User Name: %v", in.GetUserName())
	var user_id int32 = int32(rand.Intn(1000))

	return &pb.UserResponse{UserId: user_id, UserName: in.GetUserName(), Age: in.GetAge()}, nil
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()

	pb.RegisterUserManagementServer(s, &UserManagementServer{})

	log.Printf("listening on address: %v", lis.Addr())

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
