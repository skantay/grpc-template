package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/skantay/grpc-template/pkg/note/v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	lis, _ := grpc.Dial(":50051", grpc.WithTransportCredentials(insecure.NewCredentials()))

	c := note.NewNoteV1Client(lis)

	ctx, ca := context.WithTimeout(context.Background(), 5*time.Second)
	defer ca()

	r, err := c.Get(ctx, &note.GetRequest{Id: 123})
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(r)
}
