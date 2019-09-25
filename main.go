package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"time"
)

const (
	offset                  = "  "
	delayTime time.Duration = 100
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

func offsetEmpty(width, i int) Offset {
	spaceMax := width - 1

	// 周期は2*width-1
	// それを0始まりに補正するために-1
	i = i % ((2*width - 1) - 1)

	spaceNum := 0
	if i < spaceMax {
		spaceNum = i
	} else {
		spaceNum = spaceMax - (i - spaceMax)
	}
	return Offset(spaceNum)
}

func clearLine(w *bufio.Writer) {
	fmt.Fprint(w, "\033[2K")
}

func writeLine(w *bufio.Writer, width, i int) {
	fmt.Fprint(w, offset)
	fmt.Fprint(w, offsetEmpty(width, i))
	fmt.Fprint(w, "\033[46m")
	fmt.Fprint(w, " ")
	fmt.Fprint(w, "\033[0m")
	fmt.Fprint(w, "\r\n")
}

func writeHeader(w *bufio.Writer, width int) {
	fmt.Fprint(w, offset)
	for i := 0; i < width; i++ {
		fmt.Fprint(w, "|")
	}
	fmt.Fprint(w, "\n")

	w.Flush()
}

func writePendulum(w *bufio.Writer, width, height, count int) {
	for i := 0; i < count; i++ {
		for j := 0; j < height; j++ {
			writeLine(w, width, i+j)
		}
		w.Flush()
		time.Sleep(delayTime * time.Millisecond)
		for j := 0; j < height; j++ {
			fmt.Fprint(w, "\033[1F")
			clearLine(w)
		}
	}
	// delete header
	fmt.Fprint(w, "\033[1F")
	clearLine(w)
	w.Flush()
}

func main() {
	var (
		count  int
		width  int
		height int
	)
	flag.IntVar(&count, "c", 100, "how long repeat count")
	flag.IntVar(&width, "w", 40, "how long width")
	flag.IntVar(&height, "h", 20, "how long height")

	flag.Parse()

	w := bufio.NewWriter(os.Stdout)

	writeHeader(w, width)

	writePendulum(w, width, height, count)
}
