package main

import (
	"bufio"
	"context"
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

func writePendulum(ctx context.Context, w *bufio.Writer, width, height, count int) {
	// build
	lines := build(ctx, width, height, count)
	// write
	ticker := time.NewTicker(delayTime * time.Millisecond)
	defer ticker.Stop()

	for i := 0; i < count; i++ {
		select {
		case <-ticker.C:
			for _, ch := range lines {
				fmt.Fprint(w, "\033[2K"+<-ch+"\r\n")
			}
			// すべて書き込んだ後は先頭へ
			fmt.Fprintf(w, "\033[%dF", len(lines))
			w.Flush()
		case <-ctx.Done():
			return
		}
	}
	// clean up
	fmt.Fprint(w, "\033[1F")
	lenWithH := len(lines) + 1
	for i := 0; i < lenWithH; i++ {
		fmt.Fprint(w, "\033[2K\033[1E")
	}
	fmt.Fprintf(w, "\033[%dF", lenWithH)
	w.Flush()
}

func build(ctx context.Context, width, height, count int) []<-chan string {
	lines := make([]<-chan string, height)
	for j := 0; j < height; j++ {
		lines[j] = buildLine(ctx, width, count, j, "\033[46m \033[0m")
	}
	return lines
}

func buildLine(ctx context.Context, width, count, initOffset int, square string) <-chan string {
	ch := make(chan string, count)

	go func() {
		defer close(ch)

		cv := 1
		offsetLen := Offset(initOffset) + 2
		for i := 0; i < count; i++ {
			select {
			case <-ctx.Done():
				return
			default:
			}

			ch <- offsetLen.String() + square

			if offsetLen == 2 {
				cv = 1
			} else if offsetLen == Offset(width) {
				cv = -1
			}
			offsetLen += Offset(cv)
		}
	}()
	return ch
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
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	w := bufio.NewWriter(os.Stdout)

	var drawer Drawer
	drawer = Drawer{
		generator: &WaveGenerator{
			width:  width,
			height: height,
		},
		delayTime: delayTime,
		offset:    Offset(0),
		w:         w,
	}
	drawer.Draw(ctx, count)

	drawer = Drawer{
		generator: &FallnGenerator{
			width:  width,
			height: height,
		},
		delayTime: delayTime,
		offset:    Offset(0),
		w:         w,
	}
	drawer.Draw(ctx, count)
}
