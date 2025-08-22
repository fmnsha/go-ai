package repo

import (
	"go-ai/pkg/db"
	"go-ai/pkg/services/workspace/models"

	"github.com/samber/do"
	"go.mongodb.org/mongo-driver/mongo"
)

type WorkspaceRepo interface {
	db.MainRepo[models.Workspace]
}

type workspacerepo struct {
	db.MainRepoImpl[models.Workspace]
}

func NewWorkspaceRepo(i *do.Injector) (WorkspaceRepo, error) {
	return &workspacerepo{
		MainRepoImpl: db.MainRepoImpl[models.Workspace]{
			Db:       do.MustInvoke[*mongo.Client](i),
			CollName: "workspaces",
		},
	}, nil
}
