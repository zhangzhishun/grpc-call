package main

import (
	"context"
	pb "github.com/kaspanet/kaspad/infrastructure/network/netadapter/server/grpcserver/protowire"
	"google.golang.org/grpc"
	"log"
)

func main() {
	// Set up a connection to the server.
	conn, err := grpc.Dial("127.0.0.1:8080", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	// Create a new client
	client := pb.NewRPCClient(conn)

	// Call a RPC method
	stream, err := client.MessageStream(context.Background())
	if err != nil {
		log.Fatalf("error while calling MessageStream: %v", err)
	}

	// You can now use the stream to send and receive messages
	// For example:
	err = stream.Send(&pb.KaspadMessage{
		Payload: &pb.KaspadMessage_GetBlockCountRequest{
			GetBlockCountRequest: &pb.GetBlockCountRequestMessage{},
		},
	})
	if err != nil {
		log.Fatalf("error while sending to stream: %v", err)
	}

	// Receive a message from the stream
	response, err := stream.Recv()
	if err != nil {
		log.Fatalf("error while receiving from stream: %v", err)
	}
	// Process the received message
	log.Printf("Received: %v", response.String())
}
