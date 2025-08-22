package repo

import (
	"go-ai/pkg/db"
	"go-ai/pkg/services/board/models"

	"github.com/samber/do"
	"go.mongodb.org/mongo-driver/mongo"
)

type BoardRepo interface {
	db.MainRepo[models.Board]
}

type boardrepo struct {
	db.MainRepoImpl[models.Board]
}

func NewBoardRepo(i *do.Injector) (BoardRepo, error) {
	return &boardrepo{
		MainRepoImpl: db.MainRepoImpl[models.Board]{
			Db:       do.MustInvoke[*mongo.Client](i),
			CollName: "boards",
		},
	}, nil
}
