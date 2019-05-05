package main

import (
	"context"
	"errors"
	"github.com/golang/protobuf/ptypes/empty"
	todolists "go-ex-grpc-todo-lists/proto"
	"google.golang.org/grpc"
	"log"
	"net"
)

const (
	port = ":50051"
)

type server struct{}

func (s *server) GetTodo(ctx context.Context, in *todolists.GetTodoRequest) (*todolists.Todo, error) {
	log.Println("GetTodo in gRPC server")

	todo := &todolists.Todo{
		Id:       1,
		Assignee: "yutaro",
		Title:    "create Todolists",
		Body:     "practice gRPC",
	}

	if in.Id != todo.Id {
		return nil, errors.New("not find")
	}

	return todo, nil
}

func (s *server) GetTodoLists(_ *empty.Empty, stream todolists.TodoManager_GetTodoListServer) error {
	log.Println("GetTodoLists in gRPC server")

	todos := []*todolists.Todo{
		{
			Id:       1,
			Assignee: "yutaro",
			Title:    "create Todolists",
			Body:     "practice gRPC",
		},
		{
			Id:       2,
			Assignee: "yutaro",
			Title:    "learn gRPC",
			Body:     "read web sites",
		},
		{
			Id:       3,
			Assignee: "yutaro",
			Title:    "learn docker",
			Body:     "read a book",
		},
	}

	for _, todo := range todos {
		if err := stream.Send(todo); err != nil {
			return err
		}
	}

	return nil
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	todolists.RegisterTodoManagerServer(s, &server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
