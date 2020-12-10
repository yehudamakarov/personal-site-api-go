package main

import (
	"fmt"
	"log"
	"net"
	"os"

	prDomain "github.com/yehudamakarov/personal-site-proto/packages/go/pinnedRepository"
	"google.golang.org/grpc"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	prApplication "personal-site-api-go/internal/pinnedrepositorysync/application"
)

func main() {
	port := os.Getenv("PORT")
	dbUri := os.Getenv("DB_URI")
	githubCredentials := os.Getenv("GITHUB_CREDENTIALS")

	// ================================================ //
	listener, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("Can't establish listener: %v", err)
	}
	grpcServer := grpc.NewServer()

	// ================================================ //
	dbConnection, err := getDbConnection(dbUri)
	defer disconnectFromDb(dbConnection)

	// ================================================ //
	prDomain.RegisterPinnedRepositoryService(
		grpcServer,
		prApplication.GetPinnedRepositoryService(
			GithubService(githubCredentials),
			PinnedRepositoryDataAccess(dbConnection),
		),
	)

	// ================================================ //
	fmt.Printf("Server seems to be listening at %s", port)
	err = grpcServer.Serve(listener)
	if err != nil {
		log.Fatalf("gRpc server won't spin up: %v", err)
	}
}

func disconnectFromDb(client *gorm.DB) {

}

func getDbConnection(dbUri string) (*gorm.DB, error) {
	return gorm.Open(postgres.Open(dbUri), &gorm.Config{})
}
