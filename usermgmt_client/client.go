package main

import (
	"context"
	"fmt"
	"log"
	"time"

	pb "github.com/Revanth-99/usermgmt-grpc/usermgmt"
	"google.golang.org/grpc"
)

const (
	ADDRESS = "localhost:50051"
)

func main() {
	conn, err := grpc.Dial(ADDRESS, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("Unable to create connection due to error: %v", err)
	}
	defer conn.Close()

	client := pb.NewUserManagementClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	var users = make(map[string]string)
	users["bob"] = "bob.builder@gmail.com"
	users["oswold"] = "oswold.cn@yahoo.com"

	for username, email := range users {
		response, err := client.CreateNewUser(ctx, &pb.CreateNewUserRequest{User: &pb.User{Name: username, Email: email}})
		if err != nil {
			fmt.Printf("Could not create user: %s due to error: %v\n", username, err)
		}

		fmt.Printf("User created successfully with the following details \n username: %s, email: %s, id: %d \n", response.GetName(), response.GetEmail(), response.GetId())
	}
}
