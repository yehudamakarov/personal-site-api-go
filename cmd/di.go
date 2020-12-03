package main

import (
	"go.mongodb.org/mongo-driver/mongo"
	"personal-site-api-go/internal/pinnedrepositorysync/application"
	"personal-site-api-go/internal/pinnedrepositorysync/infrastructure"
)

func GetServiceGithub(githubCredentials string) application.IGithubService {
	return infrastructure.NewGithubService(githubCredentials)
}

func GetDbPinnedRepository(client *mongo.Client) application.IPinnedRepositoryData {
	return infrastructure.Init(client)
}
