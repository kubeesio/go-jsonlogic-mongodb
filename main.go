package gojsonlogicmongodb

import (
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
	}

	converted, err := convert(logic)
	if err != nil {
		logrus.Errorln("unable to convert")
	}

	logrus.WithField("data", converted).Infoln("data successfully converted")
}

func convert(rules interface{}) (result bson.D, err error) {
	if isVar(rules) {
		return true
	}

	if isMap(rules) {
		for operator, value := range rules.(map[string]interface{}) {
			if !isOperator(operator) {
				return false
			}

			return convert(value)
		}
	}

	if isSlice(rules) {
		for _, value := range rules.([]interface{}) {
			if isSlice(value) || isMap(value) {
				if convert(value) {
					continue
				}

				return false
			}

			if isVar(value) || isPrimitive(value) {
				continue
			}
		}

		return true
	}

	return isPrimitive(rules)
}
