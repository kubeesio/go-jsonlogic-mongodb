package convertors

import (
	"errors"
	"fmt"

	"github.com/kubeesio/go-jsonlogic-mongodb/helpers"
	"go.mongodb.org/mongo-driver/bson"
)

func ConvertNot(value interface{}) (bson.D, error) {
	// if the value is a map, we still need to recursively convert it
	if helpers.IsMap(value) {
		value, err := InternalConvert(value)
		if err != nil {
			return nil, err
		}

		return bson.D{{Key: "$not", Value: value}}, nil
	}

	if !helpers.IsSlice(value) {
		if helpers.IsBool(value) {
			return bson.D{{Key: "$not", Value: value}}, nil
		}

		return nil, errors.New("value must be a slice with two arguments")
	}

	arguments := value.([]interface{})

	firstArgument, secondArgument, err := GetArguments(arguments)
	if err != nil {
		return nil, err
	}

	return bson.D{{Key: "$not", Value: bson.D{{Key: "$eq", Value: bson.D{{Key: fmt.Sprint(firstArgument), Value: secondArgument}}}}}}, nil
}
