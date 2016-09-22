package main

import (
	"log"
	"net"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	pb "github.com/KanybekMomukeyev/testingGRPC/protolocation"

)

const (
	port = ":50051"
)

// server is used to implement helloworld.GreeterServer.
type kanoserver struct{}

// SayHello implements helloworld.GreeterServer
func (s *kanoserver) RpcMethod(ctx context.Context, in *pb.RpcRequest) (*pb.RpcResponse, error) {
	return &pb.RpcResponse{ResponseParam: "Hello KOKE" + in.RequestParam}, nil
}


func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterRpcGreeterServer(s, &kanoserver{})
	s.Serve(lis)
}


