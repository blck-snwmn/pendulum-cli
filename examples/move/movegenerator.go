package main

import (
	"context"
	"fmt"
	"strings"

	pendulumcli "github.com/blck-snwmn/pendulum-cli"
)

// NewGenerator return MoveGenerator
func NewGenerator(s string) *MoveGenerator {
	sr := strings.Split(s, "\n")
	max := 0
	for _, l := range sr {
		if max < len(l) {
			max = len(l)
		}
	}
	return &MoveGenerator{sr, max}
}

// MoveGenerator generate move char
type MoveGenerator struct {
	targets []string
	max     int
}

// Generate channels that create wave line
func (mg *MoveGenerator) Generate(ctx context.Context, tickNum int) []<-chan fmt.Stringer {
	return mg.build(ctx, tickNum)
}

func (mg *MoveGenerator) build(ctx context.Context, tickNum int) []<-chan fmt.Stringer {
	lines := make([]<-chan fmt.Stringer, len(mg.targets))
	for i, l := range mg.targets {
		length := len(l)

		if mg.max > length {
			// 右側スペース埋め
			l += strings.Repeat(" ", mg.max-length)
		}
		lines[i] = mg.buildLine(ctx, tickNum, []rune(l))
	}
	return lines
}

func (mg *MoveGenerator) buildLine(ctx context.Context, tickNum int, target []rune) <-chan fmt.Stringer {
	ch := make(chan fmt.Stringer, tickNum)
	go func() {
		defer close(ch)
		drawnLen := 0

		for i := 0; i < tickNum; i++ {
			select {
			case <-ctx.Done():
				return
			default:
			}
			// すべて描画している場合は、offsetによる移動のみを行う
			length := len(target) - 1 - drawnLen
			if length < 0 {
				ch <- pendulumcli.NewDrawnLine(pendulumcli.Offset(-length), string(target))
			} else {
				ch <- pendulumcli.NewDrawnLine(0, string(target[length:]))
			}
			drawnLen++
		}
	}()
	return ch
}
