package server

import (
	"context"
	pb "google.golang.org/grpc/examples/helloworld/helloworld"
)

type server struct{}

func (s *server) SayHelloAgain(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	return &pb.HelloReply{Message: "Hello again " + in.Name}, nil
}

func main() {

}