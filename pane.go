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

// DrawHeader はヘッダ情報を描画する。
// ヘッダはコマンド名を表示し、背景色を変更する。
func (p *Pane) DrawHeader() {
	const fc = termbox.ColorBlack
	const bc = termbox.ColorWhite

	// ヘッダの背景色を変更
	w, _ := termbox.Size()
	bgline := strings.Repeat(" ", w)
	p.DrawText(0, 0, []byte(bgline), fc, bc, true)

	// 上書きでテキストをセット
	now := time.Now().Format("2006/01/02 15:04:05")
	line := p.Name + " " + now
	p.DrawText(0, 0, []byte(line), fc, bc, true)
}

// DrawText はテキストを描画する。
// termbox.Flushしないので、別途Flushが必要
func (p *Pane) DrawText(x, y int, b []byte, fc, bc termbox.Attribute, chopLongLines bool) {
	prunes := NewScreenRunes(p, x, y, b, chopLongLines)
	for _, r := range prunes {
		termbox.SetCell(r.X, r.Y, r.Rune, fc, bc)
	}
}
