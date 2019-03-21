package main

import (
	"os/exec"
	"strings"
	"time"

	termbox "github.com/nsf/termbox-go"
)

type Pane struct {
	Name   string
	X      int
	Y      int
	Width  int
	Height int
}

type Offset struct {
	X int
	Y int
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
		Name:   "Pane1",
		X:      0,
		Y:      0,
		Width:  w,
		Height: h / 2,
	}
	p2 := Pane{
		Name:   "Pane2",
		X:      0,
		Y:      h / 2,
		Width:  w,
		Height: h / 2,
	}

	const fc = termbox.ColorDefault
	const bc = termbox.ColorDefault

	offset := Offset{X: 0, Y: 1}
	out, err := exec.Command("ls", "-l").Output()
	if err != nil {
		panic(err)
	}
	p1.SetHeader()
	p1.SetText(out, offset, fc, bc)

	out, err = exec.Command("ls", "-la").Output()
	if err != nil {
		panic(err)
	}
	p2.SetHeader()
	p2.SetText(out, offset, fc, bc)

	termbox.Flush()

	time.Sleep(5 * time.Second)
}

func (p *Pane) SetHeader() {
	// ヘッダの背景色を変更
	w, _ := termbox.Size()
	bgline := strings.Repeat(" ", w)
	p.SetLineText(bgline, p.Y, Offset{}, termbox.ColorBlack, termbox.ColorWhite)

	// 上書きでテキストをセット
	now := time.Now().Format("2006/01/02 03:04:05")
	line := p.Name + " " + now
	p.SetLineText(line, p.Y, Offset{}, termbox.ColorBlack, termbox.ColorWhite)
}

// SetText はテキストをペインにセットする。
// セット対象のテキストがペインの表示領域を超過しそうな場合は
// 超過しないように切り落とす。
// termbox.Flushしないので、別途Flushが必要
func (p *Pane) SetText(b []byte, offset Offset, fc, bc termbox.Attribute) {
	s := string(b)
	lines := strings.Split(s, "\n")
	for i, line := range lines {
		i += offset.Y
		if p.Height < i {
			break
		}
		y := p.Y + i
		p.SetLineText(line, y, offset, fc, bc)
	}
}

func (p *Pane) SetLineText(line string, y int, offset Offset, fc, bc termbox.Attribute) {
	for j, c := range line {
		j += offset.X
		if p.Width < j {
			// はみ出してしまっていたときは
			// テキストが切り落とされていることを明示する
			termbox.SetCell(p.X+j-1, y, '.', fc, bc)
			termbox.SetCell(p.X+j-2, y, '.', fc, bc)
			break
		}
		x := p.X + j
		termbox.SetCell(x, y, c, fc, bc)
	}
}
