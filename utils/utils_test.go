/*
Copyright © 2025 Duncan de Boer <duncan.d.boer@gmail.com>
This file is part of CLI application slack-luxafor.
*/
package utils

import (
	"os"
	"testing"
)

func Test_color(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name  string
		args  args
		want  uint8
		want1 uint8
		want2 uint8
	}{
		{"red", args{"red"}, 255, 0, 0},
		{"yellow", args{"yellow"}, 255, 255, 0},
		{"green", args{"green"}, 0, 255, 0},
		{"off", args{"off"}, 0, 0, 0},
		{"other", args{"some"}, 0, 0, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1, got2 := convertColor(tt.args.s)
			if got != tt.want {
				t.Errorf("color() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("color() got1 = %v, want %v", got1, tt.want1)
			}
			if got2 != tt.want2 {
				t.Errorf("color() got2 = %v, want %v", got2, tt.want2)
			}
		})
	}
}
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
