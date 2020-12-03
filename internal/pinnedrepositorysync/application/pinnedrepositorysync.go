package application

import (
	"context"

	pr "github.com/yehudamakarov/personal-site-proto/packages/go/pinnedRepository"
)

func GetPinnedRepositoryService(githubService IGithubService, db IPinnedRepositoryData) *pr.PinnedRepositoryService {
	prEnv := &env{
		githubService: githubService,
		persistence:   db,
	}
	return &pr.PinnedRepositoryService{Sync: prEnv.sync}
}

type env struct {
	githubService IGithubService
	persistence   IPinnedRepositoryData
}

func (e env) sync(c context.Context, req *pr.SyncPinnedRepositoryRequest) (*pr.SyncPinnedRepositoryResponse, error) {
	repos, _ := e.githubService.FetchPinnedRepositories()
	_, err := e.persistence.UpsertMany(repos)
	if err != nil {
		return nil, err
	}
	return &pr.SyncPinnedRepositoryResponse{Success: true}, nil
}
