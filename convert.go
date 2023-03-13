package gojsonlogicmongodb

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"strings"

	"github.com/diegoholiveira/jsonlogic/v3"
	"github.com/sirupsen/logrus"
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

	output, err := internalConvert(_rules)
	if err != nil {
		return nil, err
	}

	return output.(bson.D), nil
}

func internalConvert(rules interface{}) (interface{}, error) {
	if isVar(rules) {
		return bson.D{}, nil
	}

	if isMap(rules) {
		for operator, value := range rules.(map[string]interface{}) {
			switch operator {
			case "==":
				res, err := convertEqual(value)
				if err != nil {
					return nil, err
				}

				return res, nil
			case "!=":
				res, err := convertNotEqual(value)
				if err != nil {
					return nil, err
				}

				return res, nil
			}
		}
	}

	if isSlice(rules) {
		logrus.Infoln("isslice")
		// for _, value := range rules.([]interface{}) {
		// 	if isSlice(value) || isMap(value) {
		// 		if convert(value) {
		// 			continue
		// 		}

		// 		return false
		// 	}

		// 	if isVar(value) || isPrimitive(value) {
		// 		continue
		// 	}
		// }

		// return true
	}

	if isPrimitive(rules) {
		return rules, nil
	}

	// handle custom operator
	return bson.D{}, nil
}

func convertEqual(value interface{}) (bson.D, error) {
	if !isSlice(value) {
		return nil, errors.New("value must be a slice with two arguments")
	}

	arguments := value.([]interface{})

	firstArgument, secondArgument, err := getArguments(arguments)
	if err != nil {
		return nil, err
	}

	// bson.D needs a string in the first argument and accept string or float for the second
	return bson.D{{"$match", bson.D{{fmt.Sprint(firstArgument), secondArgument}}}}, nil
}

func convertNotEqual(value interface{}) (bson.D, error) {
	if !isSlice(value) {
		return nil, errors.New("value must be a slice with two arguments")
	}

	arguments := value.([]interface{})

	firstArgument, secondArgument, err := getArguments(arguments)
	if err != nil {
		return nil, err
	}

	return bson.D{{"$ne", bson.D{{fmt.Sprint(firstArgument), secondArgument}}}}, nil
}

func getArguments(arguments []interface{}) (interface{}, interface{}, error) {
	if !isVar(arguments[0]) && !isPrimitive(arguments[0]) && !isVar(arguments[1]) && !isPrimitive(arguments[1]) {
		return nil, nil, errors.New("arguments must be a primitive or a var")
	}

	firstArgument, err := internalConvert(arguments[0])
	if err != nil {
		return nil, nil, err
	}
	
	secondArgument, err := internalConvert(arguments[1])
	if err != nil {
		return nil, nil, err
	}

	return firstArgument, secondArgument, err
}