package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Data struct {
	Id        primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	BoardId   primitive.ObjectID `bson:"boardId" json:"boardId"`
	Data      map[string]any     `bson:"data" json:"data"`
	CreatedAt time.Time          `bson:"createdAt" json:"createdAt"`
	UpdatedAt time.Time          `bson:"updatedAt" json:"updatedAt"`
	Trash     bool               `bson:"trash" json:"trash"`
}

type FileData struct {
	Id           primitive.ObjectID `bson:"_id" json:"_id"`
	OriginalName string             `bson:"originalName" json:"originalName"`
	Path         string             `bson:"path" json:"path"`
	Service      string             `bson:"service" json:"service"`
	Expire       string             `bson:"expire" json:"expire"`
	Variants     []string           `bson:"variants" json:"variants"`
}
