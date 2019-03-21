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

func (p *Pane) SetHeader() {

}

// SetText はテキストをペインにセットする。
// セット対象のテキストがペインの表示領域を超過しそうな場合は
// 超過しないように切り落とす。
// termbox.Flushしないので、別途Flushが必要
func (p *Pane) SetText(b []byte) {
	s := string(b)
	lines := strings.Split(s, "\n")
	for i, line := range lines {
		if p.Height < i {
			break
		}
		y := p.Y + i
		p.SetLineText(y, line)
	}
}

func (p *Pane) SetLineText(y int, line string) {
	for j, c := range line {
		if p.Width < j {
			// はみ出してしまっていたときは
			// テキストが切り落とされていることを明示する
			termbox.SetCell(p.X+j-1, y, '.', termbox.ColorDefault, termbox.ColorDefault)
			termbox.SetCell(p.X+j-2, y, '.', termbox.ColorDefault, termbox.ColorDefault)
			break
		}
		x := p.X + j
		termbox.SetCell(x, y, c, termbox.ColorDefault, termbox.ColorDefault)
	}
}
