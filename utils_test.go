package main

import (
	"os"
	"testing"
)

func TestValidMustGetenv(t *testing.T) {
	type args struct {
		key string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"valid", args{"key"}, "value"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_ = os.Setenv(tt.args.key, tt.want)
			if got, _ := MustGetenv(tt.args.key); got != tt.want {
				t.Errorf("MustGetenv() = %v, want %v", got, tt.want)
			}
			os.Clearenv()
		})
	}
}
func TestInvalidMustGetenv(t *testing.T) {
	type args struct {
		key string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"invalid", args{"key"}, "missing environment variable 'key'"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if _, err := MustGetenv(tt.args.key); err.Error() != tt.want {
				t.Errorf("Variable wanted empty '%s'", tt.args.key)
			}
		})
	}
}
