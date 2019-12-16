package pendulumcli

import (
	"bufio"
	"bytes"
	"context"
	"fmt"
	"strings"
	"testing"
)

const (
	downOne = "\033[1E"
)

func move(i int) string {
	return fmt.Sprintf("\033[%dF", i)
}

type DummyLine struct{}

func (DummyLine) String() string {
	return "_"
}
func TestDrawer_draw(t *testing.T) {
	var buf bytes.Buffer
	buffer := bufio.NewWriter(&buf)
	offset := Offset(1)
	drawer := Drawer{
		offset:    offset,
		w:         buffer,
		delayTime: 1,
	}
	ctx := context.Background()
	t.Run("tick is 0", func(t *testing.T) {
		ch := make(chan fmt.Stringer, 1)
		ch <- DummyLine{}
		chs := []<-chan fmt.Stringer{ch}
		want := 0
		if got := drawer.draw(ctx, 0, chs); got != want {
			t.Errorf("Drawer.draw() = %v, want %v", got, want)
		}
		if buf.String() != "" {
			t.Errorf("written = %v, want %v", buf.String(), "")
		}
	})
	buf.Reset()

	t.Run("tick is 2", func(t *testing.T) {
		tickNum := 2
		dum := DummyLine{}
		ch := make(chan fmt.Stringer, tickNum)
		ch <- dum
		ch <- dum

		chs := []<-chan fmt.Stringer{ch}
		want := len(chs)

		if got := drawer.draw(ctx, tickNum, chs); got != want {
			t.Errorf("Drawer.draw() = %v, want %v", got, want)
		}
		wantStr := strings.Repeat(fmt.Sprintf("%v%v%v%v%v", clearLine, offset.String(), dum.String(), "\r\n", move(want)),
			tickNum)
		if buf.String() != wantStr {
			t.Errorf("written = %v, want %v", buf.String(), wantStr)
		}
	})
}

func TestDrawer_cleanUp(t *testing.T) {

	var buf bytes.Buffer
	buffer := bufio.NewWriter(&buf)
	drawer := Drawer{
		w:         buffer,
		delayTime: 1,
	}
	ctx := context.Background()

	t.Run("height is 0", func(t *testing.T) {
		wl := 0
		drawer.cleanUp(ctx, wl)
		if buf.String() != move(wl) {
			t.Errorf("error")
		}
	})
	buf.Reset()
	t.Run("height is 2", func(t *testing.T) {
		wl := 2
		drawer.cleanUp(ctx, wl)
		expected := clearLine + downOne +
			clearLine + downOne +
			move(wl)
		if buf.String() != expected {
			t.Errorf("written: %v, want: %v", buf.String(), expected)
		}
	})

}
