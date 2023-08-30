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

	var arguments bson.A

	for _, argument := range value.([]interface{}) {
		newArgument, err := InternalConvert(argument)
		if err != nil {
			return nil, err
		}

		arguments = append(arguments, newArgument)
	}

	return bson.D{{Key: "$and", Value: arguments}}, nil
}
