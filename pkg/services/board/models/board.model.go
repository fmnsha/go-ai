package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Board struct {
	Id        primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	BoardDto  `bson:",inline"`
	CreatedAt time.Time `bson:"created_at" json:"created_at"`
	UpdatedAt time.Time `bson:"updated_at" json:"updated_at"`
	Trash     bool      `bson:"trash" json:"trash"`
}

type BoardDto struct {
	Name   string   `bson:"name" json:"name"`
	Title  string   `bson:"title" json:"title"`
	Fields []Field  `bson:"fields" json:"fields"`
	Views  []string `bson:"views" json:"views"`
}
