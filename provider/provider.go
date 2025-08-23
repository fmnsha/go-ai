package provider

import (
	"go-ai/pkg/services/board"
	boardrepo "go-ai/pkg/services/board/repo"
	"go-ai/pkg/services/workspace"
	workspacerepo "go-ai/pkg/services/workspace/repo"

	"github.com/samber/do"
)

func Provide(i *do.Injector) {
	//board
	do.Provide(i, board.NewBoardSvcs)
	do.Provide(i, boardrepo.NewBoardRepo)
	//data
	do.Provide(i, board.NewDataSvcs)
	//group
	do.Provide(i, board.NewGroupSvcs)
	//workspace
	do.Provide(i, workspace.NewWorkspaceSvcs)
	do.Provide(i, workspacerepo.NewWorkspaceRepo)
	//dashboard

}
