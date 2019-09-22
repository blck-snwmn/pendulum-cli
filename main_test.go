package main

import "testing"

func Test_offsetEmpty(t *testing.T) {
	type args struct {
		width int
		i     int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "start", args: args{width: 4, i: 0}, want: 0},
		{name: "space max", args: args{width: 4, i: 3}, want: 3},
		{name: "decrease space", args: args{width: 4, i: 4}, want: 2},
		{name: "return start", args: args{width: 4, i: 6}, want: 0},
		{name: "after return start", args: args{width: 4, i: 7}, want: 1},
		{name: "space max at other width", args: args{width: 10, i: 9}, want: 9},
		{name: "return start at other width", args: args{width: 10, i: 18}, want: 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := offsetEmpty(tt.args.width, tt.args.i); len(got) != tt.want {
				t.Errorf("len(offsetEmpty()) = %v, want %v", len(got), tt.want)
			}
		})
	}
}
