package main

import (
	"context"
	"fmt"
	"math/rand"
	"time"
)

// Generator define how to generate line
// Generate method return channels that send one line
type Generator interface {
	Generate(ctx context.Context, tickNum int) []<-chan fmt.Stringer
}

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
		offsetLen := Offset(initOffset) + 2
		for i := 0; i < tickNum; i++ {
			select {
			case <-ctx.Done():
				return
			default:
			}

			ch <- DrawnLine{offset: offsetLen, square: square}
			switch offsetLen {
			case 2:
				cv = 1
			case Offset(w.width):
				cv = -1
			}

			offsetLen += Offset(cv)
		}
	}()
	return ch
}

// FallnGenerator generate Falln line
type FallnGenerator struct {
	width  int
	height int
}

// Generate channels that create Falln line
func (f *FallnGenerator) Generate(ctx context.Context, tickNum int) []<-chan fmt.Stringer {
	lines := make([]<-chan fmt.Stringer, f.height)
	ffirst := func() <-chan fmt.Stringer {
		ch := make(chan fmt.Stringer, tickNum)
		go func() {
			defer close(ch)
			seed := rand.NewSource(time.Now().UnixNano())
			r := rand.New(seed)
			for i := 0; i < tickNum; i++ {
				select {
				case <-ctx.Done():
					return
				case ch <- &Spin{offset: Offset(r.Intn(f.width)), stateNum: 0}:
					// case ch <- &Rain{offset: Offset(r.Intn(f.width))}:
				}

			}
		}()
		return ch
	}

	fdouble := func(in <-chan fmt.Stringer) (<-chan fmt.Stringer, <-chan fmt.Stringer) {
		chl := make(chan fmt.Stringer, tickNum)
		chr := make(chan fmt.Stringer, tickNum)
		go func() {
			defer close(chl)
			defer close(chr)

			for v := range in {
				select {
				case <-ctx.Done():
					return
				case chl <- v:
					select {
					case <-ctx.Done():
						return
					case chr <- v:
					}
				}
			}
		}()
		return chl, chr
	}

	chr := ffirst()
	for j := 0; j < f.height; j++ {
		chl, tmp := fdouble(f.buildLine(ctx, tickNum, j, chr))
		lines[j] = chl
		chr = tmp
	}
	return lines
}

func (f *FallnGenerator) buildLine(ctx context.Context, tickNum, ignoreNum int, before <-chan fmt.Stringer) <-chan fmt.Stringer {
	ch := make(chan fmt.Stringer, tickNum)
	go func() {
		defer close(ch)

		buf := make(chan fmt.Stringer, tickNum)
		defer close(buf)

		for i := 0; i < tickNum; i++ {
			select {
			case <-ctx.Done():
				return
			default:
			}
			// 無視しなくなる1つ前までのものは破棄
			t := <-before
			if ignoreNum <= 1 {
				buf <- t
			}
			var ss fmt.Stringer
			if ignoreNum <= 0 {
				ss = <-buf
			} else {
				ignoreNum--
				ss = DrawnLine{
					offset: Offset(0),
					square: "\033[47m\033[30m \033[0m",
				}
			}
			ch <- ss
		}
	}()
	return ch
}
