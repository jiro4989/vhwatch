package main

import (
	"strings"

	"github.com/mattn/go-runewidth"
)

// ScreenRune は座標とRune文字を合わせもつ
// この構造体の中の座標は、端末の画面の絶対座標である。
type ScreenRune struct {
	X    int
	Y    int
	Rune rune
}

type ScreenRunes []ScreenRune

// NewScreenRunes は端末スクリーンの絶対座標と紐づくRune文字列を生成する。
// 内部ではNewLineScreenRunesを読んでいるため詳細はそちらを参照。
// この関数では複数行のバイト文字列を処理する。
func NewScreenRunes(p *Pane, x, y int, b []byte, chopLongLines bool) (ret ScreenRunes) {
	lines := strings.Split(string(b), "\n")
	for i, line := range lines {
		i += y
		if p.Height < i {
			return
		}
		fy := p.Y + i
		lprs := NewLineScreenRunes(p, x, fy, line, chopLongLines)
		ret = append(ret, lprs...)
		// 折り返しが発生していたときに、折り返された行数分ずらす
		if !chopLongLines {
			w := runewidth.StringWidth(line)
			y += w / p.Width
		}
	}
	return
}

// NewLineScreenRunes は端末スクリーンの絶対座標と紐づくRune文字列を生成する。
// ただし渡す文字列は１行単位の文字列である必要がある。
// 文字列はchopLongLinesによって折り返す、あるいはペイン端で切り落とす。
// 返す文字列は、端末の見た目上で連続して描画されるように座標をセットするため
// マルチバイト文字が混在した場合は、1byte分X座標に空白が空く点に留意する。
func NewLineScreenRunes(p *Pane, x, y int, line string, chopLongLines bool) (ret ScreenRunes) {
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
		ret = append(ret, ScreenRune{X: fx, Y: y, Rune: c})
		// マルチバイト文字を処理したときは1文字ずらす
		if 1 < l {
			// ret = append(ret, ScreenRune{X: fx + 1, Y: y, Rune: ' '})
			colPos++
		}
		colPos++
	}
	return
}
