package main

import (
	"gorm.io/gorm"
	"personal-site-api-go/internal/pinnedrepositorysync/application"
	"personal-site-api-go/internal/pinnedrepositorysync/infrastructure"
)

func GithubService(githubCredentials string) application.IGithubService {
	return infrastructure.NewGithubService(githubCredentials)
}

func PinnedRepositoryDataAccess(client *gorm.DB) application.IPinnedRepositoryDataAccess {
	return infrastructure.Init(client)
}
