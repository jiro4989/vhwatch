package main

import (
	"strings"
	"time"

	"github.com/mattn/go-runewidth"
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
	p.DrawLineText(p.Y, bgline, termbox.ColorBlack, termbox.ColorWhite, Offset{}, true)

	// 上書きでテキストをセット
	now := time.Now().Format("2006/01/02 15:04:05")
	line := p.Name + " " + now
	p.DrawLineText(p.Y, line, termbox.ColorBlack, termbox.ColorWhite, Offset{}, true)
}

// DrawText はテキストをペインにセットする。
// セット対象のテキストがペインの表示領域を超過しそうな場合は
// 超過しないように切り落とす。
// termbox.Flushしないので、別途Flushが必要
func (p *Pane) DrawText(b []byte, fc, bc termbox.Attribute, offset Offset, chopLongLines bool) {
	var yGap int // 文字列の折り返しが発生したときのズレ行数
	s := string(b)
	lines := strings.Split(s, "\n")
	for i, line := range lines {
		i += offset.Y
		i += yGap
		if p.Height < i {
			break
		}
		y := p.Y + i
		p.DrawLineText(y, line, fc, bc, offset, chopLongLines)
		if !chopLongLines {
			w := runewidth.StringWidth(line)
			yGap += w / p.Width
		}
	}
}

func (p *Pane) DrawLineText(y int, line string, fc, bc termbox.Attribute, offset Offset, chopLongLines bool) {
	var xGap int
	for j, c := range []rune(line) {
		j += offset.X
		j += xGap // マルチバイト文字が出現した数分だけずらす
		if p.Width < j {
			// はみ出してしまっていたときは
			// テキストが切り落とされていることを明示する
			termbox.SetCell(p.X+j-1, y, '.', fc, bc)
			termbox.SetCell(p.X+j-2, y, '.', fc, bc)
			break
		}
		x := p.X + j
		termbox.SetCell(x, y, c, fc, bc)
		l := runewidth.StringWidth(string(c))
		if 1 < l {
			termbox.SetCell(x+1, y, ' ', fc, bc)
			xGap++
		}
	}
}
