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
	UseVertical   bool
	UseHorizontal bool
	Col           int
	Interval      int
	ChopLongLines bool
}

func init() {
	cobra.OnInitialize()
	RootCommand.Flags().SortFlags = false
	RootCommand.Flags().IntP("col", "c", 2, "column count")
	RootCommand.Flags().BoolP("vertical", "V", false, "vertical split panes")
	RootCommand.Flags().BoolP("horizontal", "H", false, "horizontal split panes")
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

		// コマンドラインオプションの取得
		opt, err := getOption(cmd, args)
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

func getOption(cmd *cobra.Command, args []string) (opt RootOption, err error) {
	f := cmd.Flags()

	opt.Col, err = f.GetInt("col")
	if err != nil {
		return
	}

	opt.UseVertical, err = f.GetBool("vertical")
	if err != nil {
		return
	}

	opt.UseHorizontal, err = f.GetBool("horizontal")
	if err != nil {
		return
	}

	opt.Interval, err = f.GetInt("interval")
	if err != nil {
		return
	}

	opt.ChopLongLines, err = f.GetBool("chop-long-lines")
	if err != nil {
		return
	}

	switch {
	case opt.UseVertical:
		// 垂直分割のみにする
		opt.Col = len(args)
	case opt.UseHorizontal:
		// 水平分割のみにする
		opt.Col = 1
	}

	return opt, nil
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
			out, err := pipeline.CombinedOutput(cmds...)
			// エラーが発生してもwatchを継続してほしいためerrチェックはしない
			// if string(out) == "" {
			// 	// 存在しないコマンドを実行しようとしたときはエラーが返るため
			// 	out = []byte(err.Error())
			// }
			p.DrawHeader()
			p.DrawText(0, 1, out, fc, bc, chopLongLines)
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
