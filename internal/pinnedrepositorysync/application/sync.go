package application

import (
	"context"

	pr "github.com/yehudamakarov/personal-site-proto/packages/go/pinnedRepository"
	"go.mongodb.org/mongo-driver/mongo"
	"personal-site-api-go/internal/pinnedrepositorysync/infrastructure"
)

func GetPinnedRepositoryService(githubCredentials string, client *mongo.Client) *pr.PinnedRepositoryService {
	prEnv := &env{
		githubService: infrastructure.NewGithubService(githubCredentials),
		persistence:   infrastructure.Db{Coll: client.Database("personal-site").Collection("pinned-repository")},
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
