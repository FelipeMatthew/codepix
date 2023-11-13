package grpc

import (
	"fmt"
	"log"
	"net"

	"github.com/jinzhu/gorm"
	"google.golang.org/grpc"
)

// O inicio igual express
func StartGrpcServer(database *gorm.DB, port int) {
	grpcServer := grpc.NewServer()

	// Endereço
	address := fmt.Sprintf("0.0.0.0:%d", port)
	listener, err := net.Listen("tcp", address)

	if err != nil {
		log.Fatal("CAnnot start grpc server", err)
	}

	log.Printf("gRPC server has been started on port %d", port)

	// Iniciar o server
	err = grpcServer.Serve(listener)

	if err != nil {
		log.Fatal("Cannot start gRPC server", err)
	}

}
