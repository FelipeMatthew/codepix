package cmd

import (
	"os"

	"github.com/felipematthew/codepix/application/grpc"
	"github.com/felipematthew/codepix/infrastructure/db"
	"github.com/jinzhu/gorm"
)

var database *gorm.DB

// Vai rodar as linhas de comando 
func main() {
	database = db.ConnectDB(os.Getenv("env"))

	// Start gRPC
	// Default port of gRPC server - 50051
	grpc.StartGrpcServer(database, 50051)

}

