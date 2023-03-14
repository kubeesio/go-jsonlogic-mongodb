package convertors

import (
	"errors"

	"github.com/kubeesio/go-jsonlogic-mongodb/helpers"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
)

func InternalConvert(rules interface{}) (interface{}, error) {
	if helpers.IsVar(rules) {
		return bson.D{}, nil
	}

	if helpers.IsMap(rules) {
		for operator, value := range rules.(map[string]interface{}) {
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

	if helpers.IsSlice(rules) {
		logrus.Infoln("isslice")
		// for _, value := range rules.([]interface{}) {
		// 	if IsSlice(value) || isMap(value) {
		// 		if convert(value) {
		// 			continue
		// 		}

		// 		return false
		// 	}

		// 	if IsVar(value) || isPrimitive(value) {
		// 		continue
		// 	}
		// }

		// return true
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
