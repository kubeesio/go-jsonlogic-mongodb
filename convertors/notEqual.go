package convertors

import (
	"errors"
	"fmt"

	"github.com/kubeesio/go-jsonlogic-mongodb/helpers"
	"go.mongodb.org/mongo-driver/bson"
)

func ConvertNotEqual(value interface{}) (bson.D, error) {
	if !helpers.IsSlice(value) {
		return nil, errors.New("value must be a slice with two arguments")
	}

	arguments := value.([]interface{})

	firstArgument, secondArgument, err := GetArguments(arguments)
	if err != nil {
		return nil, err
	}

	return bson.D{{"$ne", bson.D{{fmt.Sprint(firstArgument), secondArgument}}}}, nil
}
