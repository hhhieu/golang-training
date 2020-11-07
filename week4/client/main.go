package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/hieu/golang-training/week4/proto"
	"google.golang.org/grpc"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage:", os.Args[0], "NAME")
		return
	}
	personName := os.Args[1]

	conn, err := grpc.Dial("localhost:8080", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Can't connect to server %v", err)
	}
	defer conn.Close()
	c := proto.NewPersonServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.CreatePerson(ctx, &proto.CreatePersonRequest{Name: personName})
	if err != nil {
		log.Fatalf("Could not create person: %v", err)
	}
	log.Printf("Person %s is created", r)
}
