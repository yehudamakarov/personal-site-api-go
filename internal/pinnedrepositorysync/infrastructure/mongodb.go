package infrastructure

import (
	pr "github.com/yehudamakarov/personal-site-proto/packages/go/pinnedRepository"
	"go.mongodb.org/mongo-driver/mongo"
)

const database = "personal-site"
const collection = "pinned-repository"

type Db struct {
	coll *mongo.Collection
}

func Init(client *mongo.Client) *Db {
	return &Db{
		coll: client.Database(database).Collection(collection),
	}
}

func (d Db) UpsertMany([]pr.PinnedRepository) ([]pr.PinnedRepository, error) {
	return []pr.PinnedRepository{}, nil
}
