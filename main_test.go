package main

import (
	"fmt"
	"testing"
)

func Test_offsetEmpty(t *testing.T) {
	offsetFun := func(i int) string {
		return fmt.Sprintf("\033[%dC", i)
	}
	type args struct {
		width int
		i     int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{name: "start", args: args{width: 4, i: 0}, want: ""},
		{name: "space max", args: args{width: 4, i: 3}, want: offsetFun(3)},
		{name: "decrease space", args: args{width: 4, i: 4}, want: offsetFun(2)},
		{name: "return start", args: args{width: 4, i: 6}, want: ""},
		{name: "after return start", args: args{width: 4, i: 7}, want: offsetFun(1)},
		{name: "space max at other width", args: args{width: 10, i: 9}, want: offsetFun(9)},
		{name: "return start at other width", args: args{width: 10, i: 18}, want: ""},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := offsetEmpty(tt.args.width, tt.args.i); got != tt.want {
				t.Errorf("len(offsetEmpty()) = %v, want %v", got, tt.want)
			}
		})
	}
}
