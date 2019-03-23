package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseCommand(t *testing.T) {
	type TestData struct {
		expect [][]string
		err    error
		cmd    string
		msg    string
	}
	testdatas := []TestData{
		{expect: [][]string{{"ls"}}, err: nil, cmd: "ls", msg: "single command"},
		{expect: [][]string{{"grep", "--color", "-v", "test"}}, err: nil, cmd: "grep --color -v test", msg: "command and short/long options"},
		{expect: [][]string{{"echo", "1"}, {"grep", "1"}}, err: nil, cmd: "echo 1 | grep 1", msg: "command pipeline"},
		{expect: [][]string{{"echo", "あいうえお"}}, err: nil, cmd: "echo あいうえお", msg: "multi byte"},
		{expect: [][]string{{"echo", "かきくけこ"}, {"grep", "か"}}, err: nil, cmd: "echo かきくけこ | grep か", msg: "multi byte"},
		{expect: [][]string{{"grep", "-r", "s/[0-9]+/0/g"}}, err: nil, cmd: "grep -r 's/[0-9]+/0/g'", msg: "single quote"},
	}
	for _, v := range testdatas {
		got, err := ParseCommand(v.cmd)
		assert.Equal(t, v.expect, got, v.msg)
		assert.Equal(t, v.err, err, v.msg)
	}
}
