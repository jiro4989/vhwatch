package main

import (
	"os/exec"
	"time"

	"github.com/jiro4989/vhwatch/log"
	"github.com/mattn/go-shellwords"
	termbox "github.com/nsf/termbox-go"
	"github.com/spf13/cobra"
)

func init() {
	cobra.OnInitialize()
	RootCommand.Flags().IntP("col", "c", 2, "column count")
}

var RootCommand = &cobra.Command{
	Use:   "vhwatch",
	Short: "Vertical/Horizontal Watch",
	Long: `
	TODO
	`,
	Run: func(cmd *cobra.Command, args []string) {
		log.Debug("start 'vhwatch'")

		f := cmd.Flags()
		col, err := f.GetInt("col")
		if err != nil {
			panic(err)
		}

		// termboxの初期化
		if err := termbox.Init(); err != nil {
			panic(err)
		}
		defer termbox.Close()
		termbox.Flush()

		// 端末の幅を取得
		w, h := termbox.Size()

		// コマンドを描画するペインを取得
		panes := NewPanes(col, w, h, args)

		// Ctrl-Cで終了されるまで無限ループ
		for {
			for _, p := range panes {
				var out []byte
				c, err := shellwords.Parse(p.Command)
				if err != nil {
					panic(err)
				}
				switch len(c) {
				case 0:
					// 空の文字列が渡された場合
					return
				case 1:
					// コマンドのみを渡された場合
					out, err = exec.Command(c[0]).Output()
				default:
					// コマンド+オプションを渡された場合
					// オプションは可変長でexec.Commandに渡す
					out, err = exec.Command(c[0], c[1:]...).Output()
				}
				if err != nil {
					panic(err)
				}
				const fc = termbox.ColorDefault
				const bc = termbox.ColorDefault
				p.SetHeader()
				p.SetText(out, Offset{Y: 1}, fc, bc)
			}
			termbox.Flush()
			time.Sleep(1 * time.Second)
		}

		log.Debug("end 'vhwatch'")
	},
}
