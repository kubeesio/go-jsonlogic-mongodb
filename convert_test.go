package gojsonlogicmongodb

import (
	"io"
	"reflect"
	"strings"
	"testing"

	"go.mongodb.org/mongo-driver/bson"
)

func TestConvert(t *testing.T) {
	type args struct {
		rules io.Reader
	}

	tests := []struct {
		name    string
		args    args
		want    bson.D
		wantErr bool
	}{
		{name: "convert basic example", args: args{rules: strings.NewReader(`{"==": [1, 1]}`)}, want: bson.D{{"$match", bson.D{{"1", "1"}}}}, wantErr: false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Convert(tt.args.rules)
			if (err != nil) != tt.wantErr {
				t.Errorf("Convert() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			t.Errorf("%v", got)

			// if !reflect.DeepEqual(got, tt.want) {
			// 	t.Errorf("Convert() = %v, want %v", got, tt.want)
			// }
		})
	}
}

func Test_internalConvert(t *testing.T) {
	type args struct {
		rules interface{}
	}
	tests := []struct {
		name    string
		args    args
		want    interface{}
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := internalConvert(tt.args.rules)
			if (err != nil) != tt.wantErr {
				t.Errorf("internalConvert() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("internalConvert() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_convertEqual(t *testing.T) {
	type args struct {
		value interface{}
	}
	tests := []struct {
		name    string
		args    args
		want    bson.D
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := convertEqual(tt.args.value)
			if (err != nil) != tt.wantErr {
				t.Errorf("convertEqual() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("convertEqual() = %v, want %v", got, tt.want)
			}
		})
	}
}
