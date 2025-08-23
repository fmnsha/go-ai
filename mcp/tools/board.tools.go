package tools

import (
	"context"
	"encoding/json"
	"go-ai/pkg/services/board"
	"go-ai/pkg/services/board/models"
	"log"

	mcp "github.com/metoro-io/mcp-golang"
	"github.com/samber/do"
)

type Empty struct{}

type BoardTools struct {
	boardsvcs board.BoardSvcs
}

type AddRecordArgs struct {
	BoardId string                     `json:"boardId"`
	Data    map[string]json.RawMessage `json:"data"`
}

type UpdateRecordArgs struct {
	BoardId string                     `json:"boardId"`
	ItemId  string                     `json:"itemId"`
	Data    map[string]json.RawMessage `json:"data"`
}

func NewBoardTools(i *do.Injector, server *mcp.Server) {
	tools := &BoardTools{
		boardsvcs: do.MustInvoke[board.BoardSvcs](i),
	}

	if err := server.RegisterTool("add-board", "add new board", func(ctx context.Context, args models.BoardDto) (*mcp.ToolResponse, error) {
		board, err := tools.boardsvcs.AddBoard(ctx, &args)
		if err != nil {
			return nil, err
		}
		return mcp.NewToolResponse(mcp.NewTextContent(board.Id.Hex())), nil
	}); err != nil {
		log.Fatal(err)
	}

	if err := server.RegisterTool("get-all-boards", "get all boards", func(ctx context.Context, args Empty) (*mcp.ToolResponse, error) {
		boards, err := tools.boardsvcs.GetAll(ctx)
		if err != nil {
			return nil, err
		}

		j, _ := json.Marshal(boards)

		return mcp.NewToolResponse(mcp.NewTextContent(string(j))), nil
	}); err != nil {
		log.Fatal(err)
	}

	if err := server.RegisterTool("add-record", "add data to board", func(ctx context.Context, args AddRecordArgs) (*mcp.ToolResponse, error) {
		result, err := tools.boardsvcs.AddRecord(ctx, args.BoardId, args.Data)
		if err != nil {
			return nil, err
		}

		j, _ := json.Marshal(result)

		return mcp.NewToolResponse(mcp.NewTextContent(string(j))), nil
	}); err != nil {
		log.Fatal(err)
	}

	if err := server.RegisterTool("update-record", "update record in specific board", func(ctx context.Context, args UpdateRecordArgs) (*mcp.ToolResponse, error) {

		result, err := tools.boardsvcs.UpdateRecord(ctx, args.BoardId, args.ItemId, args.Data)
		if err != nil {
			return nil, err
		}

		j, _ := json.Marshal(result)

		return mcp.NewToolResponse(mcp.NewTextContent(string(j))), nil
	}); err != nil {
		log.Fatal(err)
	}

}
