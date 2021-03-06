package pendulumcli

import (
	"fmt"
)

const (
	clearLine = "\033[2K"
)

// Offset is cursor offset from left
type Offset int

func (o Offset) String() (s string) {
	// 0指定は1を指定したことと同じため、空文字にする
	if o == 0 {
		s = ""
	} else {
		s = fmt.Sprintf("\033[%dC", o)
	}
	return
}

func NewDrawnLine(offset Offset, square string) DrawnLine {
	return DrawnLine{
		offset: offset,
		square: square,
	}
}

// DrawnLine is one line
type DrawnLine struct {
	offset Offset
	square string
}

func (w DrawnLine) String() string {
	return w.offset.String() + w.square
}

type MoveParts struct {
	offset Offset
	line   []rune
	len    int
}

func (mp MoveParts) move() {

}

func (mp MoveParts) String() string {
	return mp.offset.String()
}

// NewSpin generate Spin & return it
func NewSpin(offset Offset, stateNum int) Spin {
	return Spin{
		offset:   offset,
		stateNum: stateNum,
	}
}

// Spin spin Every time String() method is called
type Spin struct {
	offset   Offset
	stateNum int
}

func (s *Spin) spin() {
	s.stateNum++
}
func (s *Spin) String() string {
	var spin string
	switch s.stateNum % 4 {
	case 0:
		spin = "\\"
	case 1:
		spin = "|"
	case 2:
		spin = "/"
	case 3:
		spin = "-"
	}
	s.spin()
	return s.offset.String() + "\033[1m" + spin + "\033[0m"
}

//Rain is Rain. move right every time called String()
type Rain struct {
	offset Offset
}

func (r *Rain) String() string {
	r.offset = r.offset + Offset(1)
	return r.offset.String() + "\033[1m" + "\\" + "\033[0m"
}
