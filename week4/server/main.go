package main

import (
	"context"
	"log"
	"net"

	"google.golang.org/grpc"

	"github.com/hieu/golang-training/week4/proto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type PersonService struct {
	PersonList []proto.Person
	CurId      int32
}

func (S PersonService) FindPerson(name string) bool {
	for _, p := range S.PersonList {
		if p.Name == name {
			return true
		}
	}
	return false
}

func (S PersonService) PrintPersonList() {
	log.Printf("List of person")
	for _, p := range S.PersonList {
		log.Printf("%v", p)
	}
}

func (S *PersonService) CreatePerson(c context.Context, req *proto.CreatePersonRequest) (*proto.Person, error) {
	if S.FindPerson(req.GetName()) {
		return nil, status.Errorf(codes.AlreadyExists, "Person name is existed")
	}
	person := proto.Person{
		Id:   S.CurId,
		Name: req.GetName(),
	}
	S.PersonList = append(S.PersonList, person)
	S.CurId++
	S.PrintPersonList()
	return &person, nil
}

func main() {
	socket, _ := net.Listen("tcp", "localhost:8080")
	server := grpc.NewServer()
	proto.RegisterPersonServiceServer(server, &PersonService{CurId: 0})
	server.Serve(socket)
}
