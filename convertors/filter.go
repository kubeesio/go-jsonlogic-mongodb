package convertors

import (
	"encoding/json"
	"errors"
	"fmt"
	"strings"

	"github.com/kubeesio/go-jsonlogic-mongodb/helpers"
	"go.mongodb.org/mongo-driver/bson"
)

func ConvertFilter(value interface{}) (bson.D, error) {
	if !helpers.IsSlice(value) {
		return nil, errors.New("value must be a slice with two arguments")
	}

	var internalError error

	firstArgument, internalError := InternalConvert(value.([]interface{})[0])
	if internalError != nil {
		return nil, internalError
	}

	secondArgumentMarshal, internalError := json.Marshal(value.([]interface{})[1])
	if internalError != nil {
		return nil, internalError
	}

	secondArgumentEdited := strings.ReplaceAll(string(secondArgumentMarshal), "\"var\":", "\"$VAR_FILTER1\":")

	var _cond interface{}

	r2 := strings.NewReader(secondArgumentEdited)
	decoderRule := json.NewDecoder(r2)
	err := decoderRule.Decode(&_cond)
	if err != nil {
		return nil, err
	}

	secondArgument, internalError := InternalConvert(_cond)
	if internalError != nil {
		return nil, internalError
	}

	return bson.D{{
		Key: "$addFields",
		Value: bson.D{{
			Key: strings.Replace(fmt.Sprint(firstArgument), "$", "", 1),
			Value: bson.D{{
				Key: "$filter",
				Value: bson.D{
					{
						Key:   "input",
						Value: firstArgument,
					},
					{
						Key:   "cond",
						Value: secondArgument,
					},
				},
			}},
		}},
	}}, nil
}
