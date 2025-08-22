package workspace

import (
	"go-ai/pkg/services/workspace/repo"

	"github.com/samber/do"
)

type WorkspaceSvcs interface {
	//AddWorkspace(ctx context.Context, data *models.WorkspaceDto) (*models.Workspace, error)
}

type workspacesvcs struct {
	repo repo.WorkspaceRepo
}

func NewWorkspaceSvcs(i *do.Injector) (WorkspaceSvcs, error) {
	return &workspacesvcs{
		repo: do.MustInvoke[repo.WorkspaceRepo](i),
	}, nil
}
