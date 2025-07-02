package main

import (
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	pb "murali.dev.grpc/proto"
)

const (
	port = ":8080"
)

func main() {
	conn, err := grpc.NewClient("localhost"+port, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect to the client %v", err)
	}

	defer conn.Close()
	client := pb.NewGreetServiceClient(conn)

	names := &pb.NamesList{
		Names: []string{"Murali", "Alice", "Bob"},
	}

	//callSayHello(client)
	//callSayHelloServerStream(client, names)
	//callSayHelloClientStream(client, names)
	callSayHelloBidirectionalStream(client, names)
}
