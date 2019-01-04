package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	pb "github.com/philipsahli/gontador.test/service"
	"google.golang.org/grpc"
)

var c pb.CounterClient
var ctx context.Context
var cancel context.CancelFunc

func inc() {
	r, err := c.Count(ctx, &pb.Empty{})
	if err != nil {
		log.Fatalf("could not increment counter: %v", err)
	}
	log.Printf("Counter: %s", fmt.Sprint(r.Counter))
}

func get() {
	r, err := c.GetCounter(ctx, &pb.Empty{})
	if err != nil {
		log.Fatalf("could not get counter: %v", err)
	}
	log.Printf("Counter: %s", fmt.Sprint(r.Counter))
}

func main() {
	// Set up a connection to the server.
	address := os.Getenv("SERVER_ADDRESS")
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c = pb.NewCounterClient(conn)

	// Contact the server and print out its response.
	// name := defaultName
	// if len(os.Args) > 1 {
	// 	name = os.Args[1]
	// }
	for {
		ctx, cancel = context.WithTimeout(context.Background(), 4*time.Second)
		inc()
		time.Sleep(2 * time.Second)
		get()
		defer cancel()
	}

}
