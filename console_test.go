package console

import (
	"testing"
)

func TestLog(t *testing.T) {
	type args struct {
		pattern string
		vars    []interface{}
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"stringtest", args{"%s", []interface{}{"stringtest"}}, "stringtest"},
		{"stringstest", args{"%s %s", []interface{}{"string1", "string2"}}, "string1 string2"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Log(tt.args.pattern, tt.args.vars...); got != tt.want {
				t.Errorf("Log() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestError(t *testing.T) {
	type args struct {
		pattern string
		vars    []interface{}
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{"stringtest", args{"%s", []interface{}{"stringtest"}}, true},
		{"stringstest", args{"%s %s", []interface{}{"string1", "string2"}}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := Error(tt.args.pattern, tt.args.vars...); (err != nil) != tt.wantErr {
				t.Errorf("Error() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
