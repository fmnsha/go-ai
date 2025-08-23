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
	AddRecord(ctx context.Context, boardId string, data map[string]json.RawMessage) (*models.Data, error)
	UpdateRecord(ctx context.Context, boardId, itemId string, data map[string]json.RawMessage) (*models.Data, error)
	GetAllRecords(ctx context.Context, boardId string) ([]models.Data, error)
}

type boardsvcs struct {
	repo           repo.BoardRepo
	datasvcsConstr DataSvcsConstructor
}

func NewBoardSvcs(i *do.Injector) (BoardSvcs, error) {
	return &boardsvcs{
		repo:           do.MustInvoke[repo.BoardRepo](i),
		datasvcsConstr: do.MustInvoke[DataSvcsConstructor](i),
	}, nil
}

// board
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

// data
func (b *boardsvcs) AddRecord(ctx context.Context, boardId string, data map[string]json.RawMessage) (*models.Data, error) {
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

	if err := b.datasvcsConstr(board.Name).AddData(ctx, d); err != nil {
		return nil, err
	}

	return d, nil
}

func (b *boardsvcs) UpdateRecord(ctx context.Context, boardId, itemId string, data map[string]json.RawMessage) (*models.Data, error) {
	_id, err := primitive.ObjectIDFromHex(itemId)
	if err != nil {
		return nil, err
	}

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
		Id:        _id,
		BoardId:   board.Id,
		Data:      _data,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	updated, err := b.datasvcsConstr(board.Name).Update(ctx, d)
	if err != nil {
		return nil, err
	}

	return updated, nil
}

func (b *boardsvcs) GetAllRecords(ctx context.Context, boardId string) ([]models.Data, error) {
	board, err := b.GetBoard(ctx, boardId)
	if err != nil {
		return nil, errors.New("board not found")
	}

	return b.datasvcsConstr(board.Name).GetAll(ctx, board.Id)
}
