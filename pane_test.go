package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewPanes(t *testing.T) {
	assert.Equal(t, Panes{
		{Name: "a", X: 0, Y: 0, Width: 30, Height: 30, Command: "a"},
		{Name: "b", X: 30, Y: 0, Width: 30, Height: 60, Command: "b"},
		{Name: "c", X: 0, Y: 30, Width: 30, Height: 30, Command: "c"},
	}, NewPanes(2, 60, 60, []string{"a", "b", "c"}))

	assert.Equal(t, Panes{
		{Name: "a", X: 0, Y: 0, Width: 30, Height: 30, Command: "a"},
		{Name: "b", X: 30, Y: 0, Width: 30, Height: 30, Command: "b"},
		{Name: "c", X: 0, Y: 30, Width: 30, Height: 30, Command: "c"},
		{Name: "d", X: 30, Y: 30, Width: 30, Height: 30, Command: "d"},
	}, NewPanes(2, 60, 60, []string{"a", "b", "c", "d"}))

	assert.Equal(t, Panes{
		{Name: "a", X: 0, Y: 0, Width: 20, Height: 20, Command: "a"},
		{Name: "b", X: 20, Y: 0, Width: 20, Height: 20, Command: "b"},
		{Name: "c", X: 40, Y: 0, Width: 20, Height: 20, Command: "c"},
		{Name: "d", X: 0, Y: 20, Width: 20, Height: 20, Command: "d"},
		{Name: "e", X: 20, Y: 20, Width: 20, Height: 40, Command: "e"},
		{Name: "f", X: 40, Y: 20, Width: 20, Height: 40, Command: "f"},
		{Name: "g", X: 0, Y: 40, Width: 20, Height: 20, Command: "g"},
	}, NewPanes(3, 60, 60, []string{"a", "b", "c", "d", "e", "f", "g"}))

}

func TestPaneWidth(t *testing.T) {
	assert.Equal(t, 30, paneWidth(2, 60))
}

func TestPaneHeight(t *testing.T) {
	assert.Equal(t, 60, paneHeight(2, 60, 2))
	assert.Equal(t, 30, paneHeight(2, 60, 3))
	assert.Equal(t, 20, paneHeight(2, 60, 5))
	assert.Equal(t, 60, paneHeight(3, 60, 3))
	assert.Equal(t, 30, paneHeight(3, 60, 4))
	assert.Equal(t, 30, paneHeight(3, 60, 6))
	assert.Equal(t, 20, paneHeight(3, 60, 7))
}

func TestPaneX(t *testing.T) {
	assert.Equal(t, 0, paneX(2, 30, 0))
	assert.Equal(t, 30, paneX(2, 30, 1))
	assert.Equal(t, 0, paneX(2, 30, 2))
	assert.Equal(t, 40, paneX(3, 20, 2))
	assert.Equal(t, 0, paneX(3, 20, 3))
	assert.Equal(t, 40, paneX(3, 20, 5))
}

func TestPaneY(t *testing.T) {
	assert.Equal(t, 0, paneY(2, 60, 0))
	assert.Equal(t, 0, paneY(2, 60, 1))
	assert.Equal(t, 30, paneY(2, 30, 2))
	assert.Equal(t, 30, paneY(2, 30, 3))
	assert.Equal(t, 0, paneY(3, 60, 2))
	assert.Equal(t, 30, paneY(3, 30, 3))
	assert.Equal(t, 30, paneY(3, 30, 4))
	assert.Equal(t, 30, paneY(3, 30, 5))
}
