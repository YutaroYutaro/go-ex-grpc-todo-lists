package client

import (
	"context"
	"github.com/golang/mock/gomock"
	"github.com/golang/protobuf/ptypes/empty"
	mock_proto "go-ex-grpc-todo-lists/mock"
	todolists "go-ex-grpc-todo-lists/proto"
	"io"
	"testing"
)

var todo1 = &todolists.Todo{
	Id:       1,
	Assignee: "Tester",
	Title:    "Test Title",
	Body:     "Test Body",
}

var todo2 = &todolists.Todo{
	Id:       2,
	Assignee: "Tester2",
	Title:    "Test Title2",
	Body:     "Test Body2",
}

func TestGetTodo(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockTodoManagerClient := mock_proto.NewMockTodoManagerClient(ctrl)
	req := &todolists.GetTodoRequest{Id: 1}
	mockTodoManagerClient.EXPECT().GetTodo(
		gomock.Any(),
		req,
	).Return(todo1, nil)

	testGetTodo(t, mockTodoManagerClient)
}

func testGetTodo(t *testing.T, client todolists.TodoManagerClient) {
	t.Helper()
	res, err := client.GetTodo(context.Background(), &todolists.GetTodoRequest{Id: 1})
	if err != nil || res.Title != todo1.Title {
		t.Errorf("mocking failed")
	}
}

func TestGetTodoLists(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	stream := mock_proto.NewMockTodoManager_GetTodoListsClient(ctrl)
	stream.EXPECT().Recv().Return(todo1, nil)

	stream.EXPECT().Recv().Return(todo2, nil)

	stream.EXPECT().Recv().Return(nil, io.EOF)

	mockTodoManagerClient := mock_proto.NewMockTodoManagerClient(ctrl)
	mockTodoManagerClient.EXPECT().GetTodoLists(
		gomock.Any(),
		gomock.Any(),
	).Return(stream, nil)

	testGetTodoLists(t, mockTodoManagerClient)
}

func testGetTodoLists(t *testing.T, client todolists.TodoManagerClient) {
	t.Helper()

	stream, err := client.GetTodoLists(context.Background(), &empty.Empty{})
	res, err := stream.Recv()
	if err != nil || res.Title != todo1.Title {
		t.Errorf("not todo1")
	}

	res, err = stream.Recv()
	if err != nil || res.Title != todo2.Title {
		t.Errorf("not todo2")
	}

	_, eof := stream.Recv()
	if eof != io.EOF {
		t.Error(eof)
	}
}
