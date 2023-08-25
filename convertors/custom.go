package convertors

import "go.mongodb.org/mongo-driver/bson/primitive"

type CustomConvertor struct {
	Name     string
	Function func(interface{}) (primitive.D, error)
}

var CustomConvertors []CustomConvertor
