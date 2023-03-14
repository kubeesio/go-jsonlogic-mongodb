package helpers

import (
	"reflect"
	"strconv"
)

func Is(obj interface{}, kind reflect.Kind) bool {
	return obj != nil && reflect.TypeOf(obj).Kind() == kind
}

func IsBool(obj interface{}) bool {
	return Is(obj, reflect.Bool)
}

func IsString(obj interface{}) bool {
	return Is(obj, reflect.String)
}

func IsNumber(obj interface{}) bool {
	return Is(obj, reflect.Float64)
}

func IsPrimitive(obj interface{}) bool {
	return IsBool(obj) || IsString(obj) || IsNumber(obj)
}

func IsMap(obj interface{}) bool {
	return Is(obj, reflect.Map)
}

func IsSlice(obj interface{}) bool {
	return Is(obj, reflect.Slice)
}

func IsTrue(obj interface{}) bool {
	if IsBool(obj) {
		return obj.(bool)
	}

	if IsNumber(obj) {
		n := ToNumber(obj)
		return n != 0
	}

	if IsString(obj) || IsSlice(obj) || IsMap(obj) {
		length := reflect.ValueOf(obj).Len()
		return length > 0
	}

	return false
}

func IsVar(value interface{}) bool {
	if !IsMap(value) {
		return false
	}

	_var, ok := value.(map[string]interface{})["var"]
	if !ok {
		return false
	}

	return IsString(_var) || IsNumber(_var) || _var == nil
}

func ToSliceOfNumbers(values interface{}) []float64 {
	_values := values.([]interface{})

	numbers := make([]float64, len(_values))
	for i, n := range _values {
		numbers[i] = ToNumber(n)
	}
	return numbers
}

func ToNumber(value interface{}) float64 {
	if IsString(value) {
		w, _ := strconv.ParseFloat(value.(string), 64)

		return w
	}

	return value.(float64)
}

func ToString(value interface{}) string {
	if IsNumber(value) {
		return strconv.FormatFloat(value.(float64), 'f', -1, 64)
	}

	if value == nil {
		return ""
	}

	return value.(string)
}
