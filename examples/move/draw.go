package main

import (
	"bufio"
	"context"
	"flag"
	pendulumcli "github.com/blck-snwmn/pendulum-cli"
	"os"
)

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

	s := `
aaaaaaaazaaaaaaaaz
bbbbbz
      cccccZcccccZ
  eeeZ
`
	drawer := pendulumcli.NewDrawer(
		NewGenerator(s),
		1000,
		pendulumcli.Offset(0),
		w,
	)
	drawer.Draw(ctx, count)
}
