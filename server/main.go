package main

import (
	"context"
	"fmt"
	"github.com/pkg/errors"
	proto "github.com/rubyconvict/lunchnlearn/proto/lunchnlearn"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/reflection"
	"google.golang.org/grpc/status"
	"net"
)

const (
	// ErrCode500 - unhandled error
	ErrCode500 Code = "500"
	// ErrCode400 - invalid request
	ErrCode400 Code = "400"
)

var clientErrors = map[Code]Error{
	ErrCode400: {Code: ErrCode400, Message: "Bad request"},
}

type Code string

type Error struct {
	Code    Code   `json:"code,omitempty"`
	Message string `json:"title,omitempty"`
}

type server struct {
	proto.UnimplementedLunchnlearnServiceServer
}

// Error - implements `error` interface
func (e *Error) Error() (err string) {
	return e.Message
}

func (s *server) PlaceOrder(ctx context.Context, request *proto.LunchRequest) (*proto.LunchResponse, error) {
	order := request.GetOrder()

	switch order {
	case "pizza":
		// unhandled internal error
		return nil, errors.Wrap(errors.New("we don't deliver pizza"), "failed to place order")
	case "pasta":
		// handled internal error
		someError := &Error{
			Code:    ErrCode400,
			Message: "we don't deliver pasta",
		}

		if value, found := clientErrors[someError.Code]; found {
			// https://grpc.github.io/grpc/core/md_doc_statuscodes.html
			return nil, status.Error(codes.InvalidArgument, value.Message)
		}

		return nil, &Error{Code: ErrCode500, Message: "Internal Server Error"}
	default:
		lunch := &proto.Lunch{
			State: proto.Lunch_ORDERED,
			Name:  fmt.Sprintf("Delicious %s!", order),
		}

		return &proto.LunchResponse{Lunch: lunch}, nil
	}
}

func main() {
	listener, err := net.Listen("tcp", ":4040")
	if err != nil {
		panic(err)
	}

	srv := grpc.NewServer()
	proto.RegisterLunchnlearnServiceServer(srv, &server{})
	reflection.Register(srv)

	if e := srv.Serve(listener); e != nil {
		panic(e)
	}

}
