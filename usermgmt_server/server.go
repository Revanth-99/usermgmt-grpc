package main

import (
	"context"
	"fmt"
	"log"
	"math/rand"
	"net"

	pb "github.com/Revanth-99/usermgmt-grpc/usermgmt"

	"google.golang.org/grpc"
)

const (
	PORT = ":50051"
)

type UserManangementServer struct {
	pb.UnimplementedUserManagementServer
}

func (s *UserManangementServer) CreateNewUser(ctx context.Context, req *pb.CreateNewUserRequest) (*pb.CreateNewUserResponse, error) {
	fmt.Printf("Received user name is : %v\n", req.GetUser().GetName())
	var user_id int32 = int32(rand.Intn(1000))
	return &pb.CreateNewUserResponse{Name: req.User.Name, Email: req.User.Email, Id: user_id}, nil
}

func main() {
	listener, err := net.Listen("tcp", PORT)
	if err != nil {
		log.Fatalf("Unable to listen on %s due to error %v", PORT, err)
	}

	s := grpc.NewServer()
	pb.RegisterUserManagementServer(s, &UserManangementServer{})
	fmt.Printf("Server is listening on %v\n", listener.Addr().String())

	if err := s.Serve(listener); err != nil {
		log.Fatalf("Failed to serve due to error: %v", err)
	}
}
