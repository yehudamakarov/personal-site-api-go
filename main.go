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
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(dbUri))
	defer func() {
		if err = client.Disconnect(ctx); err != nil {
			panic(err)
		}
	}()

	// ================================================ //
	listener, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("Can't establish listener: %v", err)
	}
	grpcServer := grpc.NewServer()

	// ================================================ //
	prDomain.RegisterPinnedRepositoryService(
		grpcServer,
		prApplication.GetPinnedRepositoryService(githubCredentials, client),
	)

	// ================================================ //
	err = grpcServer.Serve(listener)
	if err != nil {
		log.Fatalf("gRpc server won't spin up: %v", err)
	}
}
