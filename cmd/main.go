package main

import (
	"fmt"
	"log"
	"net"
	"net/http"
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
	githubCredentials := os.Getenv("GITHUB_ACCESS_TOKEN")

	// ================================================ //
	listener, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("Can't establish listener: %v", err)
	}
	grpcServer := grpc.NewServer()
	// client's settings should not to be changed per the individual infrastructure component that uses client.
	client := http.Client{}

	// ================================================ //
	dbConnection, err := getDbConnection(dbUri)
	defer disconnectFromDb(dbConnection)

	// ================================================ //
	prDomain.RegisterPinnedRepositoryService(
		grpcServer,
		prApplication.GetPinnedRepositoryService(
			GithubService(githubCredentials, &client),
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
