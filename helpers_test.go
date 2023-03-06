package gojsonlogicmongodb

import (
	"reflect"
	"testing"
)

func Test_is(t *testing.T) {
	type args struct {
		obj  interface{}
		kind reflect.Kind
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := is(tt.args.obj, tt.args.kind); got != tt.want {
				t.Errorf("is() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_isBool(t *testing.T) {
	type args struct {
		obj interface{}
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isBool(tt.args.obj); got != tt.want {
				t.Errorf("isBool() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_isString(t *testing.T) {
	type args struct {
		obj interface{}
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isString(tt.args.obj); got != tt.want {
				t.Errorf("isString() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_isNumber(t *testing.T) {
	type args struct {
		obj interface{}
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isNumber(tt.args.obj); got != tt.want {
				t.Errorf("isNumber() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_isPrimitive(t *testing.T) {
	type args struct {
		obj interface{}
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isPrimitive(tt.args.obj); got != tt.want {
				t.Errorf("isPrimitive() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_isMap(t *testing.T) {
	type args struct {
		obj interface{}
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isMap(tt.args.obj); got != tt.want {
				t.Errorf("isMap() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_isSlice(t *testing.T) {
	type args struct {
		obj interface{}
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isSlice(tt.args.obj); got != tt.want {
				t.Errorf("isSlice() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_isTrue(t *testing.T) {
	type args struct {
		obj interface{}
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isTrue(tt.args.obj); got != tt.want {
				t.Errorf("isTrue() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_isVar(t *testing.T) {
	type args struct {
		value interface{}
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isVar(tt.args.value); got != tt.want {
				t.Errorf("isVar() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_toSliceOfNumbers(t *testing.T) {
	type args struct {
		values interface{}
	}
	tests := []struct {
		name string
		args args
		want []float64
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := toSliceOfNumbers(tt.args.values); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("toSliceOfNumbers() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_toNumber(t *testing.T) {
	type args struct {
		value interface{}
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := toNumber(tt.args.value); got != tt.want {
				t.Errorf("toNumber() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_toString(t *testing.T) {
	type args struct {
		value interface{}
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := toString(tt.args.value); got != tt.want {
				t.Errorf("toString() = %v, want %v", got, tt.want)
			}
		})
	}
}
