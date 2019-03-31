package main

import (
	pb "api-grpc/proto/output"
	"context"
	"google.golang.org/grpc"
	"log"
	"net"
	"sync"
)


const (
	port = ":50051"
)

// server is used to implement helloworld.GreeterServer.
type server struct{}

func worker(wg *sync.WaitGroup, number int, lista pb.ServicioTareas_ListTasksServer) error {
	for i:=0;i<5;i++ {
		if err := lista.Send(&pb.Task{Id:uint32(i + number),Title:"pepe"}); err != nil {
			return err
		}
	}
	wg.Done()
	return nil
}

func (s *server) ListTasks(user *pb.User, lista pb.ServicioTareas_ListTasksServer) error {
	var wg sync.WaitGroup
	wg.Add(2)

	go worker(&wg,3, lista)
	go worker(&wg,22, lista)

	wg.Wait()
	return nil
}

func (s *server) AddTask(context.Context, *pb.Task) (*pb.Id, error) {
	panic("implement me")
}

func (s *server) GetTask(ctx context.Context, in *pb.Id) (*pb.Task, error) {

	log.Printf("Received: %v", in.Id)
	return &pb.Task{Id: 1,Title: "Paco"}, nil
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterServicioTareasServer(s,&server{})

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
