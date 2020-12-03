package infrastructure

import pr "github.com/yehudamakarov/personal-site-proto/packages/go/pinnedRepository"

type Github struct {
	credentials string

	// url & credentials
}

func NewGithubService(credentials string) *Github {
	return &Github{credentials: credentials}
}

func (g Github) FetchPinnedRepositories() ([]pr.PinnedRepository, error) {
	return []pr.PinnedRepository{}, nil
}
