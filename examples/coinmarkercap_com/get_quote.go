package main

import (
	"context"
	pb "github.com/panda-next-team/poolrank-proto/agent"
	"google.golang.org/grpc"
	"log"
	"time"
)

const (
	address string = "localhost:80"
)

func main() {
	// Set up a connection to the server.
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	c := pb.NewCoinMaraketCapServiceClient(conn)
	// Contact the server and print out its response.
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	r, err := c.GetQuote(ctx, &pb.GetQuoteRequest{EnTag:"BTC"})
	if err != nil {
		log.Fatalf("could not : %v", err)
	}
	log.Printf("entity: %v", r)

}
