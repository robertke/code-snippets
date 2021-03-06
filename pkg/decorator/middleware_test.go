package decorator

import (
	"log"
	"testing"
)

func TestLogger(t *testing.T) {
	logger := log.Logger{}

	type args struct {
		l *log.Logger
	}

	tests := []struct {
		name string
		args args
	}{
		{"base-case", args{&logger}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Logger(tt.args.l)
		})
	}
}
