package main

import (
	"net/http"

	"gorm.io/gorm"
	"personal-site-api-go/internal/pinnedrepositorysync/application"
	"personal-site-api-go/internal/pinnedrepositorysync/infrastructure"
)

func GithubService(githubCredentials string, client *http.Client) application.IGithubService {
	return infrastructure.NewGithubService(githubCredentials, client)
}

func PinnedRepositoryDataAccess(client *gorm.DB) application.IPinnedRepositoryDataAccess {
	return infrastructure.Init(client)
}
