package main

import (
	"context"
	"io"
	"log"

	pb "murali.dev.grpc/proto"
)

func callSayHelloServerStream(client pb.GreetServiceClient, names *pb.NamesList) {
	log.Printf("The streaming as started")

	stream, err := client.SayHelloServerStreaming(context.Background(), names)

	if err != nil {
		log.Fatalf("could not send names: %v", err)
	}

	for {
		message, err := stream.Recv()
		if err == io.EOF {
			break
		}

		if err != nil {
			log.Fatalf("err while streaming %v", err)
		}

		log.Println(message)

	}
	log.Printf("Streaming finished")
}
