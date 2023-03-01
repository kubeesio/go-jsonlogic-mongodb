package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"strings"

	"github.com/diegoholiveira/jsonlogic/v3"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
)

// sample code
func main() {
	logic := strings.NewReader(`{"==": [1, 1]}`)

	if !jsonlogic.IsValid(logic) {
		logrus.Errorln("invalid jsonlogic")
		return
	}

	var rules interface{}

	err := json.Unmarshal([]byte(`{"==": [1, 1]}`), &rules)
	if err != nil {
		logrus.WithError(err).Errorln("unable to unmarshal")
		return
	}

	// not working, why ?
	// decoderRule := json.NewDecoder(logic)
	// err = decoderRule.Decode(&rules)
	// if err != nil {
	// 	logrus.WithError(err).Errorln("unable to decode")

	// 	return
	// }

	logrus.WithField("interface", rules).Infoln("rules")

	convert(rules)
	// converted, err := convert(logic)
	// if err != nil {
	// 	logrus.Errorln("unable to convert")
	// }

	// logrus.WithField("data", converted).Infoln("data successfully converted")
}

func convert(rules interface{}) {
	if isVar(rules) {
		logrus.Infoln("isvar")
		// return true
	}

	if isMap(rules) {
		logrus.Infoln("ismap")
		for operator, value := range rules.(map[string]interface{}) {
			switch operator {
			case "==":
				res, err := convertEqual(value)
				if err != nil {
					logrus.WithError(err).Errorln("unable to convert equal operator")
				}

				logrus.WithField("converted-data", res).Infoln("yes")
			}

			// return convert(rules)
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

	logrus.WithField("yes", isPrimitive(rules)).Infoln("isprimitive")

	// return isPrimitive(rules)
}

func convertEqual(value interface{}) (bson.D, error) {
	if !isSlice(value) {
		return nil, errors.New("value must be a slice with two arguments")
	}

	arguments := value.([]interface{})

	if !isVar(arguments[0]) && !isPrimitive(arguments[0]) && !isVar(arguments[1]) && !isPrimitive(arguments[1]) {
		return nil, errors.New("arguments must be a primitive or a var")
	}

	// bson.D needs a string in the first argument
	return bson.D{{"$match", bson.D{{fmt.Sprint(arguments[0]), arguments[1]}}}}, nil
}
