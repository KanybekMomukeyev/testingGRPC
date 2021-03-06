package main

import (
	"fmt"
	"io"
	pb "github.com/KanybekMomukeyev/testingGRPC/protolocation"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"github.com/mattn/sc"
	"strconv"
)

func add(name string, age int) error {
	address := "localhost:11111"
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	//conn, err := grpc.Dial("127.0.0.1:11111")
	if err != nil {
		return err
	}
	defer conn.Close()
	client := pb.NewCustomerServiceClient(conn)

	person := &pb.Person{
		Name: name,
		Age:  int32(age),
	}

	fmt.Print("func add(name string, age int) called\n")

	_, err = client.AddPerson(context.Background(), person)
	return err
}

func list() error {

	fmt.Print("func list() error called\n")

	address := "localhost:11111"
	conn, err := grpc.Dial(address, grpc.WithInsecure())

	//conn, err := grpc.Dial("127.0.0.1:11111")
	if err != nil {
		return err
	}
	defer conn.Close()
	client := pb.NewCustomerServiceClient(conn)

	stream, err := client.ListPerson(context.Background(), new(pb.RequestType))
	if err != nil {
		return err
	}
	for {
		person, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}
		fmt.Println(person)
	}
	return nil
}

func main() {

	(&sc.Cmds{
		{
			Name: "list",
			Desc: "list: listing person",
			Run: func(c *sc.C, args []string) error {
				return list()
			},
		},
		{
			Name: "add",
			Desc: "add [name] [age]: add person",
			Run: func(c *sc.C, args []string) error {
				if len(args) != 2 {
					return sc.UsageError
				}
				name := args[0]
				age, err := strconv.Atoi(args[1])
				if err != nil {
					return err
				}
				return add(name, age)
			},
		},
	}).Run(&sc.C{})
}