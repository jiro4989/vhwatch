package main

import (
	"strings"

	"github.com/mattn/go-runewidth"
)

// PaneRune は座標とRune文字を合わせもつ
// この構造体の中の座標は、ペイン内の座標である点に注意する。
// つまり、端末の絶対座標ではなく、Paneの原点座標からの相対座標である。
type PaneRune struct {
	X    int
	Y    int
	Rune rune
}

type PaneRunes []PaneRune

// NewPaneRunes はPane用の文字を生成する。
func NewPaneRunes(p *Pane, x, y int, b []byte, chopLongLines bool) (ret PaneRunes) {
	lines := strings.Split(string(b), "\n")
	for i, line := range lines {
		i += y
		if p.Height < i {
			return
		}
		fy := p.Y + i
		lprs := NewLinePaneRunes(p, x, fy, line, chopLongLines)
		ret = append(ret, lprs...)
		// 折り返しが発生していたときに、折り返された行数分ずらす
		if !chopLongLines {
			w := runewidth.StringWidth(line)
			y += w / p.Width
		}
	}
	return
}

func NewLinePaneRunes(p *Pane, x, y int, line string, chopLongLines bool) (ret PaneRunes) {
	var colPos int
	for _, c := range []rune(line) {
		colPos += x
		l := runewidth.StringWidth(string(c))
		if p.Width < colPos+l {
			if chopLongLines {
				// はみ出してしまっていたときは
				// テキストが切り落とされていることを明示する
				if 1 <= len(ret) {
					ret[len(ret)-1].Rune = '>'
				}
				return
			}
			colPos = 0
			y++
		}
		fx := p.X + colPos
		ret = append(ret, PaneRune{X: fx, Y: y, Rune: c})
		// マルチバイト文字を処理したときは1文字ずらす
		if 1 < l {
			// ret = append(ret, PaneRune{X: fx + 1, Y: y, Rune: ' '})
			colPos++
		}
		colPos++
	}
	return
}
