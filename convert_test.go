package gojsonlogicmongodb

import (
	"io"
	"strings"
	"testing"

	"github.com/google/go-cmp/cmp"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func isKeyValue(value interface{}) (primitive.D, error) {
	parsed, _ := value.([]interface{})

	key := parsed[1].(string)
	val := parsed[2].(string)

	firstArgument, internalError := InternalConvert(value.([]interface{})[0])
	if internalError != nil {
		return nil, internalError
	}

	return bson.D{{
		"$match",
		bson.D{{
			"$expr", bson.D{{
				"$eq", bson.A{
					bson.D{{
						"$getField", bson.D{
							{"field", bson.D{{"$literal", key}}},
							{"input", firstArgument},
						},
					}},
					val,
				},
			}},
		}},
	}}, nil
}

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
		{name: "convert equal example", args: args{rules: strings.NewReader(`{"==": [1, 1]}`)}, want: bson.D{{Key: "$eq", Value: bson.A{1.0, 1.0}}}, wantErr: false},
		{name: "invalid jsonlogic", args: args{rules: strings.NewReader(`{"==": [1, ]}`)}, wantErr: true},
		{name: "convert not equal jsonlogic", args: args{rules: strings.NewReader(`{"!=": [1, 0]}`)}, want: bson.D{{Key: "$ne", Value: bson.A{1.0, 0.0}}}, wantErr: false},
		{name: "convert not jsonlogic", args: args{rules: strings.NewReader(`{"!": true}`)}, want: bson.D{{Key: "$not", Value: true}}, wantErr: false},
		{name: "convert complex not jsonlogic", args: args{rules: strings.NewReader(`{"!": {"==": [1, 1]}}`)}, want: bson.D{{Key: "$not", Value: bson.D{{Key: "$eq", Value: bson.A{1.0, 1.0}}}}}, wantErr: false},
		{name: "convert 2 comparisons joined by and example", args: args{rules: strings.NewReader(`{"and": [{"!=": [1, 2]}, {"==": [1, 1]}]}`)}, want: bson.D{{Key: "$and", Value: bson.A{bson.D{{Key: "$ne", Value: bson.A{1.0, 2.0}}}, bson.D{{Key: "$eq", Value: bson.A{1.0, 1.0}}}}}}, wantErr: false},
		{name: "convert 2 comparisons joined by or example", args: args{rules: strings.NewReader(`{"or": [{"!=": [1, 2]}, {"==": [1, 1]}]}`)}, want: bson.D{{Key: "$or", Value: bson.A{bson.D{{Key: "$ne", Value: bson.A{1.0, 2.0}}}, bson.D{{Key: "$eq", Value: bson.A{1.0, 1.0}}}}}}, wantErr: false},
		{name: "convert complex and & or", args: args{rules: strings.NewReader(`{"and": [{"or": [{"!=": [1, 2]}, {"==": [1, 1]}]}, {"and": [{"!=": [1, 2]}, {"==": [1, 1]}]}]}`)}, want: bson.D{{Key: "$and", Value: bson.A{bson.D{{Key: "$or", Value: bson.A{bson.D{{Key: "$ne", Value: bson.A{1.0, 2.0}}}, bson.D{{Key: "$eq", Value: bson.A{1.0, 1.0}}}}}}, bson.D{{Key: "$and", Value: bson.A{bson.D{{Key: "$ne", Value: bson.A{1.0, 2.0}}}, bson.D{{Key: "$eq", Value: bson.A{1.0, 1.0}}}}}}}}}, wantErr: false},
		{name: "convert simple example with var", args: args{rules: strings.NewReader(`{"==": ["observability", {"var": ".metadata.namespace"}]}`)}, want: bson.D{{Key: "$eq", Value: bson.A{"observability", "$metadata.namespace"}}}, wantErr: false},
		{name: "convert custom operator with is_key_value", args: args{rules: strings.NewReader(`{"is_key_value": [{"var":".metadata.labels"},"app.kubernetes.io/name","kubees"]}`)}, want: bson.D{{Key: "$match", Value: bson.D{{Key: "$expr", Value: bson.D{{Key: "$eq", Value: bson.A{bson.D{{"$getField", bson.D{{"field", bson.D{{"$literal", "app.kubernetes.io/name"}}}, {"input", "$metadata.labels"}}}}, "kubees"}}}}}}}, wantErr: false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.name == "convert custom operator with is_key_value" {
				AddConvertor("is_key_value", isKeyValue)
			}

			got, err := Convert(tt.args.rules)
			if (err != nil) != tt.wantErr {
				t.Errorf("Convert() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if !cmp.Equal(got, tt.want) {
				t.Errorf("Convert() = %v, want %v", got, tt.want)
			}
		})
	}
}
