package board

import (
	"context"
	"encoding/json"
	"errors"
	"go-ai/pkg/services/board/models"
	"go-ai/pkg/services/board/repo"
	"time"

	"github.com/samber/do"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type BoardSvcs interface {
	GetAll(ctx context.Context) ([]models.Board, error)
	AddBoard(ctx context.Context, data *models.BoardDto) (*models.Board, error)
	AddRecord(ctx context.Context, boardId string, data map[string]json.RawMessage) (any, error)
}

type boardsvcs struct {
	repo     repo.BoardRepo
	datasvcs DataSvcs
}

func NewBoardSvcs(i *do.Injector) (BoardSvcs, error) {
	return &boardsvcs{
		repo:     do.MustInvoke[repo.BoardRepo](i),
		datasvcs: do.MustInvoke[DataSvcs](i),
	}, nil
}

func (b *boardsvcs) GetBoard(ctx context.Context, boardId string) (*models.Board, error) {
	_id, err := primitive.ObjectIDFromHex(boardId)
	if err != nil {
		return nil, err
	}
	board, err := b.repo.GetByFilter(ctx, bson.M{"_id": _id})
	if err != nil {
		return nil, err
	}
	return board, nil
}

func (b *boardsvcs) GetAll(ctx context.Context) ([]models.Board, error) {
	match := bson.M{"$match": bson.M{}}
	pipeline := []bson.M{
		match,
	}

	var result []models.Board
	if err := b.repo.Aggregate(ctx, pipeline, func(cur *mongo.Cursor) error {
		return cur.All(ctx, &result)
	}); err != nil {
		return nil, err
	}

	return result, nil
}

func (b *boardsvcs) AddBoard(ctx context.Context, data *models.BoardDto) (*models.Board, error) {
	board := &models.Board{
		Id:        primitive.NewObjectID(),
		BoardDto:  *data,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	if err := b.repo.Add(ctx, board); err != nil {
		return nil, err
	}

	return board, nil
}

func (b *boardsvcs) AddRecord(ctx context.Context, boardId string, data map[string]json.RawMessage) (any, error) {
	board, err := b.GetBoard(ctx, boardId)
	if err != nil {
		return nil, errors.New("board not found")
	}

	fields := board.Fields

	_data := make(map[string]any)

	for _, field := range fields {
		if val, ok := data[field.Id]; ok {
			actualVal, err := field.ParseFieldValue(val)
			if err != nil {
				return nil, err
			}
			_data[field.Id] = actualVal
		}
	}

	d := &models.Data{
		Id:        primitive.NewObjectID(),
		BoardId:   board.Id,
		Data:      _data,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	if err := b.datasvcs.AddData(ctx, d); err != nil {
		return nil, err
	}

	return data, nil
}
