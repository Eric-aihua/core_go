package main

import (
	"context"
	"log"
	"time"

	pb "github.com/Eric-aihua/rpc/cal"
	"google.golang.org/grpc"
)



func main() {
	conn, err := grpc.Dial("localhost:12345", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewCalClient(conn)


	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.Add(ctx, &pb.InputElement{Input1:1,Input2:2})
	r2, err := c.Sub(ctx, &pb.InputElement{Input1:10,Input2:2})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("Add Result: %s", r.Output)
	log.Printf("Sub Result: %s", r2.Output)
}







