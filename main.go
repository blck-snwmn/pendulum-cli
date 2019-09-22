package main

import (
	"fmt"
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

func main() {
	fmt.Print(offset)
	for i := 0; i < width; i++ {
		fmt.Print("|")
	}
	fmt.Print("\n")

	var sw sync.WaitGroup
	sw.Add(1)
	printf := func() {
		defer sw.Done()
		for i := 0; i < count; i++ {
			//カーソルは先頭に
			fmt.Print(offset)
			fmt.Print(offsetEmpty(width, i))
			fmt.Print("\033[46m")
			fmt.Print(" ")
			fmt.Print("\033[0m")
			fmt.Print("\r")
			time.Sleep(100 * time.Millisecond)
			clearLine()
		}
	}

	go printf()

	sw.Wait()
}
