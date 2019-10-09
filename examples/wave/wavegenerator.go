package main

import (
	"context"
	"fmt"
	pendulumcli "github.com/blck-snwmn/pendulum-cli"
)

// WaveGenerator generate wave
type WaveGenerator struct {
	width  int
	height int
}

// Generate channels that create wave line
func (w *WaveGenerator) Generate(ctx context.Context, tickNum int) []<-chan fmt.Stringer {
	return w.build(ctx, tickNum)
}

func (w *WaveGenerator) build(ctx context.Context, tickNum int) []<-chan fmt.Stringer {
	lines := make([]<-chan fmt.Stringer, w.height)
	for j := 0; j < w.height; j++ {
		lines[j] = w.buildLine(ctx, tickNum, j, "\033[46m \033[0m")
	}
	return lines
}

func (w *WaveGenerator) buildLine(ctx context.Context, tickNum, initOffset int, square string) <-chan fmt.Stringer {
	ch := make(chan fmt.Stringer, tickNum)

	go func() {
		defer close(ch)

		cv := 1
		offsetLen := pendulumcli.Offset(initOffset) + 2
		for i := 0; i < tickNum; i++ {
			select {
			case <-ctx.Done():
				return
			default:
			}

			ch <- pendulumcli.NewDrawnLine(offsetLen, square)
			switch offsetLen {
			case 2:
				cv = 1
			case pendulumcli.Offset(w.width):
				cv = -1
			}

			offsetLen += pendulumcli.Offset(cv)
		}
	}()
	return ch
}
