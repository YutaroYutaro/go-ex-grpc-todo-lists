package main

import (
	"context"
	"github.com/golang/protobuf/ptypes/empty"
	todolists "go-ex-grpc-todo-lists/proto"
	"google.golang.org/grpc"
	"io"
	"log"
	"time"
)

const (
	address = "localhost:50051"
)

func main() {
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	c := todolists.NewTodoManagerClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	r, err := c.GetTodo(ctx, &todolists.GetTodoRequest{Id: 1})
	if err != nil {
		log.Fatalf("could not get: %v", err)
	}

	log.Println("GetTodo Id:1")
	log.Printf("\nId: %v\nAssignee: %v\nTitle: %v\nBody: %v\n\n", r.Id, r.Assignee, r.Title, r.Body)

	stream, err := c.GetTodoLists(context.Background(), &empty.Empty{})
	if err != nil {
		log.Fatalf("could not get stream: %v", err)
	}

	log.Println("GetTodoLists")
	for {
		todo, err := stream.Recv()
		if err != nil {
			log.Fatalf("could not get stream: %v", err)
		}
		if err == io.EOF {
			break
		}
		log.Printf("\nId: %v\nAssignee: %v\nTitle: %v\nBody: %v\n", todo.Id, todo.Assignee, todo.Title, todo.Body)
	}
}
