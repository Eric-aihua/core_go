package main

import
(
	"context"
	pb "github.com/Eric-aihua/rpc/cal"
	"google.golang.org/grpc"
	"log"
	"net"
)

type server struct {}

func (p *server) Add(ctx context.Context,request *pb.InputElement) (*pb.OutputElement,error){
	sum :=request.Input1+request.Input2
	return  &pb.OutputElement{Output:sum},nil
}

func (p *server) Sub(ctx context.Context,request *pb.InputElement) (*pb.OutputElement,error){
	sum :=request.Input1-request.Input2
	return  &pb.OutputElement{Output:sum},nil
}

func main() {
	lis, err := net.Listen("tcp", "12345")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterCalServer(s, &server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
