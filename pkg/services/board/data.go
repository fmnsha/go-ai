package board

import (
	"context"
	"fmt"
	"go-ai/pkg/db"
	"go-ai/pkg/services/board/models"
	"maps"

	"github.com/samber/do"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type DataSvcs interface {
	GetById(ctx context.Context, id primitive.ObjectID) (*models.Data, error)
	AddData(ctx context.Context, data *models.Data) error
	GetAll(ctx context.Context, id primitive.ObjectID) ([]models.Data, error)
	Update(ctx context.Context, data *models.Data) (*models.Data, error)
	//Update(ctx context.Context, data *models.Data) (*models.Data, error)
}

type datasvcs struct {
	repo db.MainRepo[models.Data]
}

type DataSvcsConstructor func(collName string) DataSvcs

func NewDataSvcs(i *do.Injector) (DataSvcsConstructor, error) {
	return func(collName string) DataSvcs {
		return &datasvcs{
			repo: &db.MainRepoImpl[models.Data]{
				Db:       do.MustInvoke[*mongo.Client](i),
				CollName: collName,
			},
		}
	}, nil
}

func (d *datasvcs) GetById(ctx context.Context, id primitive.ObjectID) (*models.Data, error) {
	return d.repo.GetByFilter(ctx, bson.M{"_id": id, "trash": false})

}

func (d *datasvcs) GetAll(ctx context.Context, id primitive.ObjectID) ([]models.Data, error) {

	filter := bson.M{"trash": false}

	pipeline := []bson.M{
		{"$match": filter},
		{"$sort": bson.M{"createdAt": -1}},
	}

	var items []models.Data
	d.repo.Aggregate(ctx, pipeline, func(cur *mongo.Cursor) error {
		return cur.All(ctx, &items)
	})

	fmt.Println(items)
	return items, nil
}

func (d *datasvcs) AddData(ctx context.Context, data *models.Data) error {
	return d.repo.Add(ctx, data)
}

func (d *datasvcs) Update(ctx context.Context, data *models.Data) (*models.Data, error) {
	oldData, err := d.GetById(ctx, data.Id)
	if err != nil {
		return nil, err
	}

	maps.Copy(oldData.Data, data.Data)

	filter := bson.M{"_id": data.Id}
	update := bson.M{"$set": bson.M{"data": oldData.Data}}
	updatedData, err := d.repo.Patch(ctx, filter, update)
	if err != nil {
		return nil, err
	}

	return updatedData, nil
}
