package convertors

import (
	"errors"
	"strings"

	"github.com/kubeesio/go-jsonlogic-mongodb/helpers"
	"go.mongodb.org/mongo-driver/bson"
)

var customOperators = make(map[string]func(value interface{}) (bson.D, error))

func AddOperator(jsonlogicKeyword string, fn func(value interface{}) (bson.D, error)) {
	customOperators[jsonlogicKeyword] = fn
}

func InternalConvert(rules interface{}) (interface{}, error) {
	if helpers.IsVar(rules) {
		arguments := rules.(map[string]interface{})
		varData := strings.Replace(arguments["var"].(string), ".", "$", 1)
		return varData, nil
	}

	if helpers.IsMap(rules) {
		for operator, value := range rules.(map[string]interface{}) {

			for index, customOperation := range customOperators {
				if operator == index {
					return customOperation(value)
				}
			}

			switch operator {
			case "==":
				res, err := ConvertEqual(value)
				if err != nil {
					return nil, err
				}

				return res, nil
			case "!=":
				res, err := ConvertNotEqual(value)
				if err != nil {
					return nil, err
				}

				return res, nil
			case "!":
				res, err := ConvertNot(value)
				if err != nil {
					return nil, err
				}

				return res, nil
			case "and":
				res, err := ConvertAnd(value)
				if err != nil {
					return nil, err
				}

				return res, nil
			case "or":
				res, err := ConvertOr(value)
				if err != nil {
					return nil, err
				}

				return res, nil
			}
		}
	}

	if helpers.IsPrimitive(rules) {
		return rules, nil
	}

	// handle custom operator
	return bson.D{}, nil
}

func GetArguments(arguments []interface{}) (interface{}, interface{}, error) {
	if !helpers.IsVar(arguments[0]) && !helpers.IsPrimitive(arguments[0]) && !helpers.IsVar(arguments[1]) && !helpers.IsPrimitive(arguments[1]) {
		return nil, nil, errors.New("arguments must be a primitive or a var")
	}

	firstArgument, err := InternalConvert(arguments[0])
	if err != nil {
		return nil, nil, err
	}

	secondArgument, err := InternalConvert(arguments[1])
	if err != nil {
		return nil, nil, err
	}

	return firstArgument, secondArgument, err
}
