package gojsonlogicmongodb

import (
	"encoding/json"
	"errors"
	"io"
	"strings"

	"github.com/diegoholiveira/jsonlogic/v3"
	"github.com/kubeesio/go-jsonlogic-mongodb/convertors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func Convert(rules io.Reader) (bson.D, error) {
	rulesByte, err := io.ReadAll(rules)
	if err != nil {
		return nil, err
	}

	r1 := strings.NewReader(string(rulesByte))
	if !jsonlogic.IsValid(r1) {
		return nil, errors.New("invalid jsonlogic")
	}

	var _rules interface{}

	r2 := strings.NewReader(string(rulesByte))
	decoderRule := json.NewDecoder(r2)
	err = decoderRule.Decode(&_rules)
	if err != nil {
		return nil, err
	}

	output, err := convertors.InternalConvert(_rules)
	if err != nil {
		return nil, err
	}

	return output.(bson.D), nil
}

func AddConvertor(name string, function func(interface{}) (primitive.D, error)) {
	convertors.CustomConvertors = append(convertors.CustomConvertors, convertors.CustomConvertor{
		Name:     name,
		Function: function,
	})

	// add fake jsonlogic function so it validates the operator name
	jsonlogic.AddOperator(name, func(values interface{}, data interface{}) (result interface{}) { return })
}

func GetArguments(arguments []interface{}) (interface{}, interface{}, error) {
	return convertors.GetArguments(arguments)
}

func InternalConvert(rules interface{}) (interface{}, error) {
	return convertors.InternalConvert(rules)
}
