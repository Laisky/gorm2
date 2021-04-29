package gorm

import (
	"reflect"
	"testing"
	"time"
)

type leak string

func (l leak) String() string {
	panic("yo")
}

func TestLogFormatter(t *testing.T) {
	type args struct {
		values []interface{}
	}
	tests := []struct {
		name         string
		args         args
		wantMessages []interface{}
	}{
		{"0", args{[]interface{}{"sql", "", time.Duration(123), "", []interface{}{leak("yo")}, int64(12)}}, []interface{}{"yo"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotMessages := LogFormatter(tt.args.values...); !reflect.DeepEqual(gotMessages, tt.wantMessages) {
				t.Errorf("LogFormatter() = %v, want %v", gotMessages, tt.wantMessages)
			}
		})
	}
}
