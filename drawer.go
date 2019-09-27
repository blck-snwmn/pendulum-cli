package main

import (
	"bufio"
	"context"
	"fmt"
	"time"
)

// Drawer draw lines
type Drawer struct {
	generator Generator
	delayTime time.Duration
	offset    Offset
	w         *bufio.Writer
}

// Draw line and cleanup
func (d *Drawer) Draw(ctx context.Context, tickNum int) {
	lines := d.generate(ctx, tickNum)
	wl := d.draw(ctx, tickNum, lines)
	d.cleanUp(ctx, wl)
}
func (d *Drawer) generate(ctx context.Context, tickNum int) []<-chan fmt.Stringer {
	return d.generator.Generate(ctx, tickNum)
}
func (d *Drawer) draw(ctx context.Context, tickNum int, lines []<-chan fmt.Stringer) int {
	ticker := time.NewTicker(d.delayTime * time.Millisecond)
	defer ticker.Stop()

	len := len(lines)
	isWrite := false
loop:
	for i := 0; i < tickNum; i++ {
		select {
		case <-ticker.C:
			for _, ch := range lines {
				fmt.Fprint(d.w, "\033[2K")
				fmt.Fprint(d.w, d.offset)
				fmt.Fprint(d.w, <-ch)
				fmt.Fprint(d.w, "\r\n")
			}
			// すべて書き込んだ後は先頭へ
			fmt.Fprintf(d.w, "\033[%dF", len)
			d.w.Flush()
			// 一度書き込んだら行数分削除処理を入れる
			isWrite = true
		case <-ctx.Done():
			break loop
		}
	}

	var writedLine int
	if isWrite {
		writedLine = len
	} else {
		writedLine = 0
	}
	return writedLine
}
func (d *Drawer) cleanUp(ctx context.Context, wl int) {
	for i := 0; i < wl; i++ {
		fmt.Fprint(d.w, "\033[2K\033[1E")
	}
	fmt.Fprintf(d.w, "\033[%dF", wl)
	d.w.Flush()
}
