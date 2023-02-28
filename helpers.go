package gojsonlogicmongodb

import (
	"reflect"
	"strconv"
)

func is(obj interface{}, kind reflect.Kind) bool {
	return obj != nil && reflect.TypeOf(obj).Kind() == kind
}

func isBool(obj interface{}) bool {
	return is(obj, reflect.Bool)
}

func isString(obj interface{}) bool {
	return is(obj, reflect.String)
}

func isNumber(obj interface{}) bool {
	return is(obj, reflect.Float64)
}

func isPrimitive(obj interface{}) bool {
	return isBool(obj) || isString(obj) || isNumber(obj)
}

func isMap(obj interface{}) bool {
	return is(obj, reflect.Map)
}

func isSlice(obj interface{}) bool {
	return is(obj, reflect.Slice)
}

func isTrue(obj interface{}) bool {
	if isBool(obj) {
		return obj.(bool)
	}

	if isNumber(obj) {
		n := toNumber(obj)
		return n != 0
	}

	if isString(obj) || isSlice(obj) || isMap(obj) {
		length := reflect.ValueOf(obj).Len()
		return length > 0
	}

	return false
}

func isVar(value interface{}) bool {
	if !isMap(value) {
		return false
	}

	_var, ok := value.(map[string]interface{})["var"]
	if !ok {
		return false
	}

	return isString(_var) || isNumber(_var) || _var == nil
}

func toSliceOfNumbers(values interface{}) []float64 {
	_values := values.([]interface{})

	numbers := make([]float64, len(_values))
	for i, n := range _values {
		numbers[i] = toNumber(n)
	}
	return numbers
}

func toNumber(value interface{}) float64 {
	if isString(value) {
		w, _ := strconv.ParseFloat(value.(string), 64)

		return w
	}

	return value.(float64)
}

func toString(value interface{}) string {
	if isNumber(value) {
		return strconv.FormatFloat(value.(float64), 'f', -1, 64)
	}

	if value == nil {
		return ""
	}

	return value.(string)
}

func isOperator(op string) bool {
	operators := []string{
		"==",
		"!=",
		"!",
		"or",
		"and",
		"filter",
		"var",
	}

	for customOperator := range customOperators {
		operators = append(operators, customOperator)
	}

	for _, operator := range operators {
		if operator == op {
			return true
		}
	}

	return false
}