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

	arguments := value.([]interface{})

	firstArgument, secondArgument, err := GetArguments(arguments)

	if err != nil {
		var internalError error

		if firstArgument == nil {
			firstArgument, internalError = InternalConvert(value.([]interface{})[0])
			if internalError != nil {
				return nil, internalError
			}
		}

		if secondArgument == nil {
			secondArgument, internalError = InternalConvert(value.([]interface{})[1])
			if internalError != nil {
				return nil, internalError
			}
		}
	}

	return bson.D{{Key: "$and", Value: bson.A{firstArgument, secondArgument}}}, nil
}
