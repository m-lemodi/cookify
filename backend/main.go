package cookify

import (
	"context"
	"fmt"
	"log"
	"net"

	pb "github.com/m-lemodi/cookify/backend/proto"

	"google.golang.org/grpc"
)

type Cookify struct{}

func (x *Cookify) GetRecipe(ctx context.Context, request *pb.Ingredients) (*pb.Recipe, error) {
	return &pb.Recipe{
		Name: "Sean",
	}, nil
}

func main() {
	fmt.Println("Hello, World! This is a server.")
	listen, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcserver := grpc.NewServer()
	pb.RegisterCookifyServer(grpcserver, &Cookify{})

	err = grpcserver.Serve(listen)
	if err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
