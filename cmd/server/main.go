package main

import (
	"context"
	"log"
	"net"

	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	validator "github.com/grpc-ecosystem/go-grpc-middleware/validator"
	"github.com/skantay/grpc-template/pkg/note/v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type server struct {
	note.UnimplementedNoteV1Server
}

func (s *server) Get(ctx context.Context, g *note.GetRequest) (*note.GetResponse, error) {

	return &note.GetResponse{
		Note: &note.Note{
			Id: g.Id,
			Info: &note.NoteInfo{
				Title: "sdf",
			},
		},
	}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatal(err)
	}

	s := grpc.NewServer(
		grpc.UnaryInterceptor(
			grpc_middleware.ChainUnaryServer(
				validator.UnaryServerInterceptor(),
			),
		))
	reflection.Register(s)

	note.RegisterNoteV1Server(s, &server{})

	s.Serve(lis)
}
