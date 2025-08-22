package tools

import (
	"context"
	"go-ai/pkg/services/workspace"
	"go-ai/pkg/services/workspace/models"
	"log"

	mcp "github.com/metoro-io/mcp-golang"
	"github.com/samber/do"
)

type WorkspaceTools struct {
	workspacesvcs workspace.WorkspaceSvcs
}

func NewWorkspaceTools(i *do.Injector, server *mcp.Server) {
	// tools := &WorkspaceTools{
	// 	workspacesvcs: do.MustInvoke[workspace.WorkspaceSvcs](i),
	// }

	if err := server.RegisterTool("add workspace", "add new workspace", func(ctx context.Context, args models.WorkspaceDto) (*mcp.ToolResponse, error) {
		// workspace, err := tools.workspacesvcs.AddWorkspace(ctx, &args)
		// if err != nil {
		// 	return nil, err
		// }
		return mcp.NewToolResponse(mcp.NewTextContent("ok")), nil
	}); err != nil {
		log.Fatal(err)
	}

}
