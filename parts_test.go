package pendulumcli

import (
	"testing"
)

func TestOffset_String(t *testing.T) {
	tests := []struct {
		name  string
		o     Offset
		wantS string
	}{
		{name: "empty", o: Offset(0), wantS: ""},
		{name: "non empty", o: Offset(9), wantS: "\033[9C"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotS := tt.o.String(); gotS != tt.wantS {
				t.Errorf("Offset.String() = %v, want %v", gotS, tt.wantS)
			}
		})
	}
}

func TestDrawnLine_String(t *testing.T) {
	type fields struct {
		offset Offset
		square string
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{"empty", fields{0, "s"}, "s"},
		{"non empty", fields{2, "s"}, "\033[2Cs"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := DrawnLine{
				offset: tt.fields.offset,
				square: tt.fields.square,
			}
			if got := w.String(); got != tt.want {
				t.Errorf("String() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSpin_String(t *testing.T) {
	type fields struct {
		offset   Offset
		stateNum int
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{"stateNum = 0 & offset is non empty",
			fields{3, 0}, "\033[3C\033[1m\\\033[0m"},
		{"stateNum = 1 & offset is non empty",
			fields{2, 1}, "\033[2C\033[1m|\033[0m"},
		{"stateNum = 2 & offset is non empty",
			fields{1, 2}, "\033[1C\033[1m/\033[0m"},
		{"stateNum = 3 & offset is empty",
			fields{0, 3}, "\033[1m-\033[0m"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := &Spin{
				offset:   tt.fields.offset,
				stateNum: tt.fields.stateNum,
			}
			if got := w.String(); got != tt.want {
				t.Errorf("String() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRain_String(t *testing.T) {
	type fields struct {
		offset Offset
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{"show", fields{3}, "\033[4C\033[1m\\\033[0m"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &Rain{
				offset: tt.fields.offset,
			}
			if got := r.String(); got != tt.want {
				t.Errorf("String() = %v, want %v", got, tt.want)
			}
		})
	}
}
