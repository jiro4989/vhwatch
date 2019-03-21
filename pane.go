package main

import (
	"strings"
	"time"

	termbox "github.com/nsf/termbox-go"
)

type Pane struct {
	Name    string
	X       int
	Y       int
	Width   int
	Height  int
	Command string
}

type Panes []Pane

type Offset struct {
	X int
	Y int
}

func NewPanes(c, tw, th int, cmds []string) (ret Panes) {
	cmdLen := len(cmds)
	for i, cmd := range cmds {
		pw := paneWidth(c, tw)          // 1ペインあたりの幅
		ph := paneHeight(c, th, cmdLen) // 1ペインあたりの幅

		x := paneX(c, pw, i)
		y := paneY(c, ph, i)
		// ペインの下のペインが空いていたら縦幅を拡張
		h := ph
		if cmdLen-1 < i+c && (cmdLen-1)/c != i/c {
			h += ph
		}
		p := Pane{
			Name:    cmd,
			X:       x,
			Y:       y,
			Width:   pw,
			Height:  h,
			Command: cmd,
		}
		ret = append(ret, p)
	}
	return
}

func paneWidth(c, w int) int {
	return w / c
}

func paneHeight(c, h, cnt int) int {
	mod := cnt % c
	div := cnt / c
	if 0 < mod {
		div++
	}
	return h / div
}

func paneX(c, w, i int) int {
	return w * (i % c)
}

func paneY(c, h, i int) int {
	return h * (i / c)
}

func (p *Pane) DrawHeader() {
	// ヘッダの背景色を変更
	w, _ := termbox.Size()
	bgline := strings.Repeat(" ", w)
	p.DrawLineText(bgline, p.Y, Offset{}, termbox.ColorBlack, termbox.ColorWhite)

	// 上書きでテキストをセット
	now := time.Now().Format("2006/01/02 15:04:05")
	line := p.Name + " " + now
	p.DrawLineText(line, p.Y, Offset{}, termbox.ColorBlack, termbox.ColorWhite)
}

// DrawText はテキストをペインにセットする。
// セット対象のテキストがペインの表示領域を超過しそうな場合は
// 超過しないように切り落とす。
// termbox.Flushしないので、別途Flushが必要
func (p *Pane) DrawText(b []byte, offset Offset, fc, bc termbox.Attribute) {
	s := string(b)
	lines := strings.Split(s, "\n")
	for i, line := range lines {
		i += offset.Y
		if p.Height < i {
			break
		}
		y := p.Y + i
		p.DrawLineText(line, y, offset, fc, bc)
	}
}

func (p *Pane) DrawLineText(line string, y int, offset Offset, fc, bc termbox.Attribute) {
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
