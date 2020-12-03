package infrastructure

import (
	pr "github.com/yehudamakarov/personal-site-proto/packages/go/pinnedRepository"
	"go.mongodb.org/mongo-driver/mongo"
)

type Db struct {
	Coll *mongo.Collection
}

func (d Db) UpsertMany([]pr.PinnedRepository) ([]pr.PinnedRepository, error) {
	return []pr.PinnedRepository{}, nil
}
