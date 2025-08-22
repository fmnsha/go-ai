package models

import (
	"encoding/json"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type FieldType string

const (
	FieldTypeText     FieldType = "text"
	FieldTypeNumber   FieldType = "number"
	FieldTypePhone    FieldType = "phone"
	FieldTypeFile     FieldType = "file"
	FieldTypeDate     FieldType = "date"
	FieldTypeEmail    FieldType = "email"
	FieldTypeSelect   FieldType = "select"
	FieldTypeBoolean  FieldType = "boolean"
	FieldTypeObjectId FieldType = "objectId"
)

func (ft FieldType) String() string {
	return string(ft)
}

func (ft FieldType) IsValid() bool {
	switch ft {
	case FieldTypeText, FieldTypeNumber, FieldTypePhone, FieldTypeFile,
		FieldTypeDate, FieldTypeEmail,
		FieldTypeBoolean:
		return true
	}
	return false
}

type Field struct {
	Id      string         `bson:"id" json:"id"`
	Label   string         `bson:"label" json:"label"`
	Type    FieldType      `bson:"type" json:"type"`
	IsMulti bool           `bson:"isMulti" json:"isMulti"`
	Meta    map[string]any `bson:"meta" json:"meta"`
	Desc    string         `bson:"desc" json:"desc"`
}

type FileField struct {
	Id           primitive.ObjectID `bson:"_id" json:"_id"`
	OriginalName string             `bson:"originalName" json:"originalName"`
	Path         string             `bson:"path" json:"path"`
	Service      string             `bson:"service" json:"service"`
	Expire       string             `bson:"expire" json:"expire"`
	Variants     []string           `bson:"variants" json:"variants"`
}

func (f *Field) ParseFieldValue(value json.RawMessage) (any, error) {

	switch f.Type {
	case FieldTypeText, FieldTypeEmail, FieldTypePhone, FieldTypeSelect:
		{
			if f.IsMulti {
				var val []string
				err := json.Unmarshal(value, &val)
				if err != nil {
					return nil, err
				}
				return val, nil
			} else {
				var val string
				err := json.Unmarshal(value, &val)
				if err != nil {
					return nil, fmt.Errorf("invalid value for %s field '%s': expected string", f.Type, f.Label)
				}
				return val, nil
			}
		}
	case FieldTypeNumber:
		{
			if f.IsMulti {
				var val []float64
				err := json.Unmarshal(value, &val)
				if err != nil {
					return nil, fmt.Errorf("invalid value for %s field '%s': expected array of numbers", f.Type, f.Label)
				}
				return val, nil
			} else {
				var val float64
				err := json.Unmarshal(value, &val)
				if err != nil {
					return nil, fmt.Errorf("invalid value for %s field '%s': expected number", f.Type, f.Label)
				}
				return val, nil
			}
		}

	case FieldTypeFile:
		{
			if f.IsMulti {
				var val []FileData
				err := json.Unmarshal(value, &val)
				if err != nil {
					return nil, fmt.Errorf("invalid value for %s field '%s': expected array of file objects", f.Type, f.Label)
				}
				return val, nil
			} else {
				var val FileData
				err := json.Unmarshal(value, &val)
				if err != nil {
					return nil, fmt.Errorf("invalid value for %s field '%s': expected file object", f.Type, f.Label)
				}
				return val, nil
			}
		}
	case FieldTypeDate:
		{
			if f.IsMulti {
				var val []time.Time
				err := json.Unmarshal(value, &val)
				if err != nil {
					return nil, fmt.Errorf("invalid value for %s field '%s': expected array of dates", f.Type, f.Label)
				}
				return val, nil
			} else {
				var val time.Time
				err := json.Unmarshal(value, &val)
				if err != nil {
					return nil, fmt.Errorf("invalid value for %s field '%s': expected date", f.Type, f.Label)
				}
				return val, nil
			}
		}
	case FieldTypeBoolean:
		{
			var val bool
			err := json.Unmarshal(value, &val)
			if err != nil {
				return nil, fmt.Errorf("invalid value for %s field '%s': expected boolean (true/false)", f.Type, f.Label)
			}
			return val, nil
		}
	case FieldTypeObjectId:
		{
			var val primitive.ObjectID
			err := json.Unmarshal(value, &val)
			if err != nil {
				return nil, fmt.Errorf("invalid value for %s field '%s': expected string", f.Type, f.Label)
			}

			return val, nil
		}
	default:
		{
			return nil, fmt.Errorf("unsupported field type '%s' for field '%s'", f.Type, f.Label)
		}
	}
}
