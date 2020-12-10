package infrastructure

import (
	pr "github.com/yehudamakarov/personal-site-proto/packages/go/pinnedRepository"
	"gorm.io/gorm"
)

func Init(db *gorm.DB) *PinnedRepositoryDataAccess {
	return &PinnedRepositoryDataAccess{
		db: db,
	}
}

type PinnedRepositoryDataAccess struct {
	db *gorm.DB
}

func (d PinnedRepositoryDataAccess) UpsertMany([]pr.PinnedRepository) ([]pr.PinnedRepository, error) {
	return []pr.PinnedRepository{}, nil
}
