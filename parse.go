package main

import (
	"github.com/mattn/go-shellwords"
)

func isSpace(r byte) bool {
	switch r {
	case ' ', '\t', '\r', '\n':
		return true
	}
	return false
}

func ParseCommand(line string) ([][]string, error) {
	parser := shellwords.NewParser()
	var ret [][]string
	for {
		args, err := parser.Parse(line)
		if err != nil {
			return nil, err
		}
		ret = append(ret, args)
		if parser.Position < 0 {
			break
		}
		i := parser.Position
		for ; i < len(line); i++ {
			if isSpace(line[i]) {
				break
			}
		}
		line = line[i+1:]
	}
	return ret, nil
}
