package gojsonlogicmongodb

import (
	"encoding/json"
	"errors"
	"io"
	"strings"

	"github.com/diegoholiveira/jsonlogic/v3"
	"github.com/kubeesio/go-jsonlogic-mongodb/convertors"
	"go.mongodb.org/mongo-driver/bson"
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

func AddOperator(jsonlogicKeyword string, fn func(value interface{}) (bson.D, error)) {
	convertors.CustomOperators[jsonlogicKeyword] = fn

	// register custom operator to jsonLogic library in order to not throw an error during the jsonlogic validation
	jsonlogic.AddOperator(jsonlogicKeyword, func(values interface{}, data interface{}) (result interface{}) { return "" })
}
