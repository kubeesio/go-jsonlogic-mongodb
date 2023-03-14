package gojsonlogicmongodb

import (
	"io"
	"strings"
	"testing"

	"github.com/google/go-cmp/cmp"
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
		{name: "convert equal example", args: args{rules: strings.NewReader(`{"==": [1, 1]}`)}, want: bson.D{{"$match", bson.D{{"1", 1.0}}}}, wantErr: false},
		{name: "invalid jsonlogic", args: args{rules: strings.NewReader(`{"==": [1, ]}`)}, wantErr: true},
		{name: "convert not equal jsonlogic", args: args{rules: strings.NewReader(`{"!=": [1, 0]}`)}, want: bson.D{{"$ne", bson.D{{"1", 0.0}}}}, wantErr: false},
		{name: "convert not jsonlogic", args: args{rules: strings.NewReader(`{"!": true}`)}, want: bson.D{{"$not", true}}, wantErr: false},
		// {name: "convert not array jsonlogic", args: args{rules: strings.NewReader(`{"!": [true, true]}`)}, want: bson.D{{"$not", ?}}, wantErr: false},
		{name: "convert complex not jsonlogic", args: args{rules: strings.NewReader(`{"!": {"==": [1, 1]}}`)}, want: bson.D{{"$not", bson.D{{"$match", bson.D{{"1", 1.0}}}}}}, wantErr: false},
		// {name: "convert 2 equal joined by and example", args: args{rules: strings.NewReader(`{"and": [{"!=": [1, 0]}, {"==": [1, 1]}]}`)}, want: bson.D{{"$and", bson.A{bson.D{{"$ne", bson.D{{"1", 0.0}}}}, bson.D{{"$match", bson.D{{"1", 1.0}}}}}}}, wantErr: false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
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
