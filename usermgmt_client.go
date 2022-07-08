package main

import (
	"context"
	"log"
	"time"

	pb "example.com/go-usermgmt-grpc/usermgmt"
	"google.golang.org/grpc"
)

const (
	address = "localhost:50001"
)

func main() {
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("didn't connect: %v", err)
	}

	defer conn.Close()

	c := pb.NewUserManagementClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)

	defer cancel()

	var new_users = make(map[string]int32)
	new_users["Alice"] = 30
	new_users["Bob"] = 40

	for name, age := range new_users {
		resp, err := c.CreateUser(ctx, &pb.UserRequest{UserName: name, Age: age})

		if err != nil {
			log.Fatalf("server call failed: %v", err)
		}

		log.Printf(`User Details:
Name: %s
Age: %d
Id: %d `, resp.GetUserName(), resp.GetAge(), resp.GetUserId())
	}
}
