package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
)

var (
	count  = 39
	width  = 20
	height = 20
)

const (
	offset                  = "  "
	delayTime time.Duration = 100
)

func offsetEmpty(width, i int) string {
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
	return strings.Repeat(" ", spaceNum)
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

	w.Flush()
}

func main() {
	w := bufio.NewWriter(os.Stdout)

	writeHeader(w, width)

	writePendulum(w, width, height, count)
}
