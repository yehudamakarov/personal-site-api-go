package application

import (
	"context"

	pr "github.com/yehudamakarov/personal-site-proto/packages/go/pinnedRepository"
	"personal-site-api-go/internal/pinnedrepositorysync/infrastructure"
)

func GetPinnedRepositoryService(githubCredentials string, db infrastructure.Db) *pr.PinnedRepositoryService {
	prEnv := &env{
		githubService: infrastructure.NewGithubService(githubCredentials),
		persistence:   db,
	}
	return &pr.PinnedRepositoryService{Sync: prEnv.sync}

}

type iPinnedRepositoryData interface {
	UpsertMany([]pr.PinnedRepository) ([]pr.PinnedRepository, error)
}

type iGithubService interface {
	FetchPinnedRepositories() ([]pr.PinnedRepository, error)
}

type env struct {
	githubService iGithubService
	persistence   iPinnedRepositoryData
}

func (e env) sync(c context.Context, req *pr.SyncPinnedRepositoryRequest) (*pr.SyncPinnedRepositoryResponse, error) {
	repos, _ := e.githubService.FetchPinnedRepositories()
	_, err := e.persistence.UpsertMany(repos)
	if err != nil {
		return nil, err
	}
	return &pr.SyncPinnedRepositoryResponse{Success: true}, nil
}
