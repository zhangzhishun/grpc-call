package main

import (
	"context"
	"google.golang.org/grpc"
	pb "grpc-call/call_by_proto/protowire/github.com/kaspanet/kaspad/protowire"
	"log"
)

func main() {
	var conn *grpc.ClientConn
	conn, err := grpc.Dial("127.0.0.1:8080", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %s \n", err)
	}
	defer conn.Close()

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
