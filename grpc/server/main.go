package main

import (
	"log"
	"net"

	pb "github.com/diptadas/golang-examples/grpc/proto"
	pr "github.com/diptadas/golang-examples/grpc/proxy"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

const (
	port = ":50051"
)

type server struct{}

func (s *server) SayHello(ctx context.Context, in *pb.MyRequest) (*pb.MyReply, error) {
	log.Println("Client request: " + in.Name)
	return &pb.MyReply{Message: "Hello " + in.Name}, nil
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterMyServiceServer(s, &server{})

	// Register reflection service on gRPC server.
	reflection.Register(s)

	log.Println("Server running at port 50051")

	go pr.Call() //run reverse proxy server in another routine

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

}
