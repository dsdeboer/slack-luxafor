package main

import "testing"

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
		{"Test - red", args{"red"}, 255, 0, 0},
		{"Test - yellow", args{"yellow"}, 255, 255, 0},
		{"Test - green", args{"green"}, 0, 255, 0},
		{"Test - off", args{"off"}, 0, 0, 0},
		{"Test - other", args{"some"}, 0, 0, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1, got2 := color(tt.args.s)
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
