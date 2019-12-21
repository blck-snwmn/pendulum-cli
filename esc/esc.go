package esc

const (
	head      = "\033["
	clearLine = head + "2K"

	graphicReset = head + "0m"

	styleBold = head + "1m"
)
