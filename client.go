package main


import (
	pb "api-grpc/proto/output"
	"context"
	"google.golang.org/grpc"
	"io"
	"log"
	"time"
)

const (
	address     = "localhost:50051"
)

func main() {
	// Set up a connection to the server.
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	client := pb.NewServicioTareasClient(conn)


	ctx, cancel := context.WithTimeout(context.Background(), time.Hour)
	defer cancel()

	r,err := client.GetTask(ctx, &pb.Id{Id:1})

	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}

	log.Printf("Greeting: %s", r)

	stream, err := client.ListTasks(ctx,&pb.User{Name:"paco"})

	if err != nil {
		log.Fatalf("%v.ListTasks(_) = _, %v", client, err)
	}

	for {
		feature, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("%v.ListFeatures(_) = _, %v", client, err)
		}
		log.Println(feature)
	}
}