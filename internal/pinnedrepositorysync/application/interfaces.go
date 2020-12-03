package application

import "github.com/yehudamakarov/personal-site-proto/packages/go/pinnedRepository"

type IPinnedRepositoryData interface {
	UpsertMany([]pinnedRepository.PinnedRepository) ([]pinnedRepository.PinnedRepository, error)
}

type IGithubService interface {
	FetchPinnedRepositories() ([]pinnedRepository.PinnedRepository, error)
}
