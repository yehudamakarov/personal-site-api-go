package main

import (
	"context"
	"log"
	"net"
	"os"
	"time"

	prDomain "github.com/yehudamakarov/personal-site-proto/packages/go/pinnedRepository"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"google.golang.org/grpc"
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
	client, err := getDbConnection(dbUri)
	defer disconnectFromDb(client)

	// ================================================ //
	prDomain.RegisterPinnedRepositoryService(
		grpcServer,
		prApplication.GetPinnedRepositoryService(
			GetServiceGithub(githubCredentials),
			GetDbPinnedRepository(client),
		),
	)

	// ================================================ //
	err = grpcServer.Serve(listener)
	if err != nil {
		log.Fatalf("gRpc server won't spin up: %v", err)
	}
}

func disconnectFromDb(client *mongo.Client) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	if err := client.Disconnect(ctx); err != nil {
		log.Fatalf("There was a problem disconnecting from the DB: ")
	}
}

func getDbConnection(dbUri string) (*mongo.Client, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(dbUri))
	if err != nil {
		log.Fatalf("There was a problem connecting to the DB: %v", err)
	}
	return client, err
}
