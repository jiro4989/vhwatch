package main

import (
	"time"

	pipeline "github.com/mattn/go-pipeline"
	termbox "github.com/nsf/termbox-go"
	"github.com/spf13/cobra"
)

func init() {
	cobra.OnInitialize()
	RootCommand.Flags().IntP("col", "c", 2, "column count")
	RootCommand.Flags().IntP("interval", "n", 2, "seconds to wait between updates")
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

		col, err := f.GetInt("col")
		if err != nil {
			panic(err)
		}

		interval, err := f.GetInt("interval")
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
		go mainloop(col, args, time.Duration(interval))

		// Ctrl-Cで終了されるまで待機
		waitKeyInput()
	},
}

func mainloop(col int, args []string, interval time.Duration) {
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
				panic(err)
			}
			out, err := pipeline.Output(cmds...)
			if err != nil {
				panic(err)
			}
			p.DrawHeader()
			p.DrawText(out, Offset{Y: 1}, fc, bc)
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
