package main

import (
	"google.golang.org/grpc"
	"log"
	"net"
	"test_tablelink/config"
	"test_tablelink/repository"
	"test_tablelink/server"
	"test_tablelink/usecase"
	"test_tablelink/user/proto"
)

func main() {
	lis, err := net.Listen("tcp", "127.0.0.1:5100")
	if err != nil {
		log.Fatalf("Failed to listen %v ", err)
	}
	db := config.InitDB()
	repo := repository.NewRepository(db)
	uc := usecase.NewUsecase(repo)

	grpcServer := grpc.NewServer()
	proto.RegisterUserServer(grpcServer, &server.UserServer{Uc: uc})

	log.Print("listening to port 5100")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve %v", err)
	}

}
