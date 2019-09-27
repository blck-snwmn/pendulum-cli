package main

// DrawnLine is one line
type DrawnLine struct {
	offset Offset
	square string
}

func (w DrawnLine) String() string {
	return w.offset.String() + w.square
}

// Spin spin Every time String() method is called
type Spin struct {
	offset   Offset
	stateNum int
}

func (w *Spin) String() string {
	var spin string
	switch w.stateNum % 4 {
	case 0:
		spin = "\\"
	case 1:
		spin = "|"
	case 2:
		spin = "/"
	case 3:
		spin = "-"
	}
	w.stateNum = (w.stateNum + 1) % 4
	return w.offset.String() + "\033[1m" + spin + "\033[0m"
}

//Rain is Rain. move right every time called String()
type Rain struct {
	offset Offset
}

func (r *Rain) String() string {
	r.offset = r.offset + Offset(1)
	return r.offset.String() + "\033[1m" + "\\" + "\033[0m"
}
