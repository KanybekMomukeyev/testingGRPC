package main

import (
	"golang.org/x/net/context"
	"log"
	"net"
	"sync"

	"google.golang.org/grpc"
	pb "github.com/KanybekMomukeyev/testingGRPC/protolocation"
	"fmt"
)

type customerService struct {
	customers []*pb.Person
	m         sync.Mutex
}

func (cs *customerService) ListPerson(p *pb.RequestType, stream pb.CustomerService_ListPersonServer) error {

	fmt.Print("ListPerson called from server")

	cs.m.Lock()
	defer cs.m.Unlock()
	for _, p := range cs.customers {
		if err := stream.Send(p); err != nil {
			return err
		}
	}
	return nil
}

func (cs *customerService) AddPerson(c context.Context, p *pb.Person) (*pb.ResponseType, error) {

	fmt.Print("AddPerson called from server")

	cs.m.Lock()
	defer cs.m.Unlock()
	cs.customers = append(cs.customers, p)
	return new(pb.ResponseType), nil
}

func main() {
	lis, err := net.Listen("tcp", ":11111")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	server := grpc.NewServer()

	pb.RegisterCustomerServiceServer(server, new(customerService))
	server.Serve(lis)
}
