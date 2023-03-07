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
