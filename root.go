package main

import (
	"fmt"
	"log"
	"os"
	"time"

	pipeline "github.com/mattn/go-pipeline"
	termbox "github.com/nsf/termbox-go"
	"github.com/spf13/cobra"
)

type RootOption struct {
	Col           int
	Interval      int
	ChopLongLines bool
}

func init() {
	cobra.OnInitialize()
	RootCommand.Flags().IntP("col", "c", 2, "column count")
	RootCommand.Flags().IntP("interval", "n", 2, "seconds to wait between updates")
	RootCommand.Flags().BoolP("chop-long-lines", "S", false, "cause lines longer than the screen width to be chopped (truncated)")
}

var RootCommand = &cobra.Command{
	Use:     "vhwatch",
	Short:   "vhwatch is Vertical/Horizontal Watch",
	Example: "vhwatch -c 3 'echo test' 'date' 'ls -1' 'ls -lah'",
	Version: Version,
	Long: `
vhwatch provides watching multiple commands execution.

Repository: https://github.com/jiro4989/vhwatch
    Author: jiro4989
	`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 1 {
			cmd.Help()
			return
		}

		f := cmd.Flags()
		var opt RootOption
		var err error

		opt.Col, err = f.GetInt("col")
		if err != nil {
			panic(err)
		}

		opt.Interval, err = f.GetInt("interval")
		if err != nil {
			panic(err)
		}

		opt.ChopLongLines, err = f.GetBool("chop-long-lines")
		if err != nil {
			panic(err)
		}

		// termboxの初期化
		if err := termbox.Init(); err != nil {
			panic(err)
		}
		defer termbox.Close()
		termbox.SetInputMode(termbox.InputEsc)
		termbox.Flush()

		// 各ペイン毎にコマンドを定期実行
		go mainloop(args, opt)

		// Ctrl-Cで終了されるまで待機
		waitKeyInput()
	},
}

func mainloop(args []string, opt RootOption) {
	col := opt.Col
	interval := time.Duration(opt.Interval)
	chopLongLines := opt.ChopLongLines
	const fc = termbox.ColorDefault
	const bc = termbox.ColorDefault
	for {
		// 端末の幅を取得
		w, h := termbox.Size()
		// コマンドを描画するペインを取得
		panes := NewPanes(col, w, h, args)
		for _, p := range panes {
			cmds, err := ParseCommand(p.Command)
			if err != nil {
				termbox.Close()
				log.Println(fmt.Sprintf("parse command error. command=%v, err=%v", p.Command, err))
				os.Exit(1)
			}
			out, err := pipeline.Output(cmds...)
			if err != nil {
				termbox.Close()
				log.Println(fmt.Sprintf("execute commands error. commands=%v, err=%v", cmds, err))
				os.Exit(2)
			}
			p.DrawHeader()
			p.DrawText(out, fc, bc, Offset{Y: 1}, chopLongLines)
		}
		termbox.Flush()
		time.Sleep(interval * time.Second)
		termbox.Clear(fc, bc)
	}
}

func waitKeyInput() {
	for {
		switch ev := termbox.PollEvent(); ev.Type {
		case termbox.EventKey:
			switch ev.Key {
			case termbox.KeyCtrlC, termbox.KeyCtrlD:
				return
			}
		}
	}
}
