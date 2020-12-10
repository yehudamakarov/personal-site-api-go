package application

import (
	"context"

	pr "github.com/yehudamakarov/personal-site-proto/packages/go/pinnedRepository"
)

func GetPinnedRepositoryService(githubService IGithubService, pinnedRepositoryDataAccess IPinnedRepositoryDataAccess) *pr.PinnedRepositoryService {
	prEnv := &env{
		githubService:              githubService,
		pinnedRepositoryDataAccess: pinnedRepositoryDataAccess,
	}
	return &pr.PinnedRepositoryService{Sync: prEnv.sync}
}

type env struct {
	githubService              IGithubService
	pinnedRepositoryDataAccess IPinnedRepositoryDataAccess
}

func (e env) sync(c context.Context, req *pr.SyncPinnedRepositoryRequest) (*pr.SyncPinnedRepositoryResponse, error) {
	repos, _ := e.githubService.FetchPinnedRepositories()
	_, err := e.pinnedRepositoryDataAccess.UpsertMany(repos)
	if err != nil {
		return nil, err
	}
	return &pr.SyncPinnedRepositoryResponse{Success: true}, nil
}
