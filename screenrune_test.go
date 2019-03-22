package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewScreenRunes(t *testing.T) {
	p := &Pane{X: 0, Y: 0, Width: 6, Height: 10}

	assert.Equal(t, ScreenRunes{
		{X: 0, Y: 0, Rune: 't'},
		{X: 1, Y: 0, Rune: 'e'},
		{X: 2, Y: 0, Rune: 's'},
		{X: 3, Y: 0, Rune: 't'},
	}, NewScreenRunes(p, 0, 0, []byte("test"), false), "半角英数字のみ")

	assert.Equal(t, ScreenRunes{
		{X: 0, Y: 0, Rune: 't'},
		{X: 1, Y: 0, Rune: 'e'},
		{X: 2, Y: 0, Rune: 's'},
		{X: 3, Y: 0, Rune: 't'},
		{X: 0, Y: 1, Rune: 't'},
		{X: 1, Y: 1, Rune: 'e'},
	}, NewScreenRunes(p, 0, 0, []byte("test\nte"), false), "半角英数字のみ+改行")

	assert.Equal(t, ScreenRunes{
		{X: 0, Y: 0, Rune: 't'},
		{X: 1, Y: 0, Rune: 'e'},
		{X: 2, Y: 0, Rune: 'あ'},
		{X: 0, Y: 1, Rune: 'う'},
		{X: 2, Y: 1, Rune: 'え'},
	}, NewScreenRunes(p, 0, 0, []byte("teあ\nうえ"), false), "半角全角+改行")

	assert.Equal(t, ScreenRunes{
		{X: 0, Y: 0, Rune: 't'},
		{X: 1, Y: 0, Rune: 'e'},
		{X: 2, Y: 0, Rune: 's'},
		{X: 3, Y: 0, Rune: 't'},
		{X: 4, Y: 0, Rune: 'A'},
		{X: 5, Y: 0, Rune: 'B'},
		{X: 0, Y: 1, Rune: 'C'},
		{X: 0, Y: 2, Rune: 't'},
		{X: 1, Y: 2, Rune: 'e'},
	}, NewScreenRunes(p, 0, 0, []byte("testABC\nte"), false), "半角英数字のみ+改行+折り返し")

	assert.Equal(t, ScreenRunes{
		{X: 0, Y: 0, Rune: 't'},
		{X: 1, Y: 0, Rune: 'e'},
		{X: 2, Y: 0, Rune: 's'},
		{X: 3, Y: 0, Rune: 't'},
		{X: 4, Y: 0, Rune: 'A'},
		{X: 5, Y: 0, Rune: '>'},
		{X: 0, Y: 1, Rune: 't'},
		{X: 1, Y: 1, Rune: 'e'},
	}, NewScreenRunes(p, 0, 0, []byte("testABC\nte"), true), "半角英数字のみ+改行+切り詰め")

}
func TestNewLineScreenRunes(t *testing.T) {
	p := &Pane{X: 0, Y: 0, Width: 6, Height: 10}

	assert.Equal(t, ScreenRunes{
		{X: 0, Y: 0, Rune: 't'},
		{X: 1, Y: 0, Rune: 'e'},
		{X: 2, Y: 0, Rune: 's'},
		{X: 3, Y: 0, Rune: 't'},
	}, NewLineScreenRunes(p, 0, 0, "test", false), "半角英数字のみ")

	assert.Equal(t, ScreenRunes{
		{X: 0, Y: 0, Rune: 'a'},
		{X: 1, Y: 0, Rune: 'あ'},
		{X: 3, Y: 0, Rune: 'b'},
	}, NewLineScreenRunes(p, 0, 0, "aあb", false), "半角全角の混在")

	assert.Equal(t, ScreenRunes{
		{X: 0, Y: 0, Rune: '漢'},
		{X: 2, Y: 0, Rune: '字'},
	}, NewLineScreenRunes(p, 0, 0, "漢字", false), "全角のみ")

	assert.Equal(t, ScreenRunes{
		{X: 0, Y: 0, Rune: '1'},
		{X: 1, Y: 0, Rune: '2'},
		{X: 2, Y: 0, Rune: '3'},
		{X: 3, Y: 0, Rune: '4'},
		{X: 4, Y: 0, Rune: '5'},
		{X: 5, Y: 0, Rune: '6'},
		{X: 0, Y: 1, Rune: 'あ'},
		{X: 2, Y: 1, Rune: 'い'},
	}, NewLineScreenRunes(p, 0, 0, "123456あい", false), "折り返し")

	assert.Equal(t, ScreenRunes{
		{X: 0, Y: 0, Rune: '1'},
		{X: 1, Y: 0, Rune: '2'},
		{X: 2, Y: 0, Rune: '3'},
		{X: 3, Y: 0, Rune: '4'},
		{X: 4, Y: 0, Rune: '5'},
		{X: 0, Y: 1, Rune: 'あ'},
		{X: 2, Y: 1, Rune: 'い'},
	}, NewLineScreenRunes(p, 0, 0, "12345あい", false), "折り返し（マルチバイト）")

	p = &Pane{X: 0, Y: 0, Width: 4, Height: 10}

	assert.Equal(t, ScreenRunes{
		{X: 0, Y: 0, Rune: 't'},
		{X: 1, Y: 0, Rune: 'e'},
		{X: 2, Y: 0, Rune: 's'},
		{X: 3, Y: 0, Rune: 't'},
	}, NewLineScreenRunes(p, 0, 0, "test", true), "切り詰めなし")

	assert.Equal(t, ScreenRunes{
		{X: 0, Y: 0, Rune: 't'},
		{X: 1, Y: 0, Rune: 'e'},
		{X: 2, Y: 0, Rune: 's'},
		{X: 3, Y: 0, Rune: '>'},
	}, NewLineScreenRunes(p, 0, 0, "testA", true), "切り詰めあり")

	assert.Equal(t, ScreenRunes{
		{X: 0, Y: 0, Rune: 'あ'},
		{X: 2, Y: 0, Rune: 'い'},
	}, NewLineScreenRunes(p, 0, 0, "あい", true), "切り詰めなし(全角)")

	assert.Equal(t, ScreenRunes{
		{X: 0, Y: 0, Rune: 'a'},
		{X: 1, Y: 0, Rune: '>'},
	}, NewLineScreenRunes(p, 0, 0, "aあい", true), "切り詰めあり(全角)")

	assert.Equal(t, ScreenRunes{
		{X: 0, Y: 0, Rune: 'あ'},
		{X: 2, Y: 0, Rune: '>'},
	}, NewLineScreenRunes(p, 0, 0, "あいう", true), "切り詰めあり(全角)")

}
