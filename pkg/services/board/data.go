package board

import (
	"context"
	"go-ai/pkg/services/board/models"
	"go-ai/pkg/services/board/repo"

	"github.com/samber/do"
)

type DataSvcs interface {
	AddData(ctx context.Context, data *models.Data) error
}

type datasvcs struct {
	repo repo.DataRepo
}

func NewDataSvcs(i *do.Injector) (DataSvcs, error) {
	return &datasvcs{
		repo: do.MustInvoke[repo.DataRepo](i),
	}, nil
}

func (d *datasvcs) AddData(ctx context.Context, data *models.Data) error {
	return d.repo.Add(ctx, data)
}
