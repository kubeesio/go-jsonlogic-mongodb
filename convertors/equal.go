package convertors

import (
	"errors"

	"github.com/kubeesio/go-jsonlogic-mongodb/helpers"
	"go.mongodb.org/mongo-driver/bson"
)

func ConvertEqual(value interface{}) (bson.D, error) {
	if !helpers.IsSlice(value) {
		return nil, errors.New("value must be a slice with two arguments")
	}

	arguments := value.([]interface{})

	firstArgument, secondArgument, err := GetArguments(arguments)
	if err != nil {
		return nil, err
	}

	// bson.D needs a string in the first argument and accept string or float for the second
	return bson.D{{Key: "$eq", Value: bson.A{firstArgument, secondArgument}}}, nil
}
