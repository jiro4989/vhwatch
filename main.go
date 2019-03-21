package main

import (
	"os/exec"
	"strings"
	"time"

	termbox "github.com/nsf/termbox-go"
)

type Pane struct {
	X      int
	Y      int
	Width  int
	Height int
}

func main() {
	// termboxの初期化
	if err := termbox.Init(); err != nil {
		panic(err)
	}
	defer termbox.Close()
	termbox.SetInputMode(termbox.InputEsc)
	termbox.Flush()

	w, h := termbox.Size()
	p1 := Pane{
		X:      0,
		Y:      0,
		Width:  w,
		Height: h / 2,
	}
	p2 := Pane{
		X:      0,
		Y:      h / 2,
		Width:  w,
		Height: h / 2,
	}

	out, err := exec.Command("ls", "-l").Output()
	if err != nil {
		panic(err)
	}
	p1.SetText(out)

	out, err = exec.Command("ls", "-la").Output()
	if err != nil {
		panic(err)
	}
	p2.SetText(out)

	termbox.Flush()

	time.Sleep(5 * time.Second)
}

// Flushはしない
func (p *Pane) SetText(b []byte) {
	s := string(b)
	lines := strings.Split(s, "\n")
	for i, line := range lines {
		if p.Height < i {
			break
		}
		y := p.Y + i
		for j, c := range line {
			if p.Width < j {
				break
			}
			x := p.X + j
			termbox.SetCell(x, y, c, termbox.ColorDefault, termbox.ColorDefault)
		}
	}
}
