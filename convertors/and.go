package convertors

import (
	"errors"

	"github.com/kubeesio/go-jsonlogic-mongodb/helpers"
	"go.mongodb.org/mongo-driver/bson"
)

func ConvertAnd(value interface{}) (bson.D, error) {
	if !helpers.IsSlice(value) {
		return nil, errors.New("value must be a slice with two arguments")
	}

	var internalError error

	firstArgument, internalError := InternalConvert(value.([]interface{})[0])
	if internalError != nil {
		return nil, internalError
	}

	secondArgument, internalError := InternalConvert(value.([]interface{})[1])
	if internalError != nil {
		return nil, internalError
	}

	return bson.D{{Key: "$and", Value: bson.A{firstArgument, secondArgument}}}, nil
}
