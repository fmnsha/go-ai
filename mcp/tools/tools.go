package tools

import (
	mcp "github.com/metoro-io/mcp-golang"
	"github.com/samber/do"
)

func Register(i *do.Injector, server *mcp.Server) {
	NewBoardTools(i, server)
	//NewWorkspaceTools(i, server)
}
