package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"sync"
	"time"
)

var (
	count = 37
	width = 10
)

const (
	offset = "  "
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

func clearLine() {
	fmt.Print("\033[2K")
}

func writeLine(w *bufio.Writer, width, i int) {
	fmt.Fprint(w, offset)
	fmt.Fprint(w, offsetEmpty(width, i))
	fmt.Fprint(w, "\033[46m")
	fmt.Fprint(w, " ")
	fmt.Fprint(w, "\033[0m")
	fmt.Fprint(w, "\r")
}

func main() {
	w := bufio.NewWriter(os.Stdout)
	fmt.Fprint(w, offset)
	for i := 0; i < width; i++ {
		fmt.Fprint(w, "|")
	}
	fmt.Fprint(w, "\n")

	w.Flush()

	var sw sync.WaitGroup
	sw.Add(1)
	printf := func() {
		defer sw.Done()
		for i := 0; i < count; i++ {
			writeLine(w, width, i)
			w.Flush()
			time.Sleep(100 * time.Millisecond)
			clearLine()
		}
	}

	go printf()

	sw.Wait()

	w.Flush()
}
