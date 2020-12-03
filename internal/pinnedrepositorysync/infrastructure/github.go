package infrastructure

import pr "github.com/yehudamakarov/personal-site-proto/packages/go/pinnedRepository"

type GithubService struct {
	credentials string
}

func NewGithubService(credentials string) *GithubService {
	return &GithubService{credentials: credentials}
}

func (g GithubService) FetchPinnedRepositories() ([]pr.PinnedRepository, error) {
	return []pr.PinnedRepository{}, nil
}
