package main

import (
	"fmt"

	"github.com/jiro4989/vhwatch/log"
	"github.com/spf13/cobra"
)

func init() {
	cobra.OnInitialize()
	RootCommand.Flags().IntP("col", "c", 2, "column count")
	RootCommand.Flags().StringArrayP("vertical", "V", nil, "vertical")
	RootCommand.Flags().StringArrayP("horizontal", "H", nil, "horizontal")
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

		verticalCmds, err := f.GetStringArray("vertical")
		if err != nil {
			panic(err)
		}

		horizontalCmds, err := f.GetStringArray("horizontal")
		if err != nil {
			panic(err)
		}

		log.Debug(fmt.Sprintf("command line option parameteres. "+
			"vertical=%v, horizontal=%v",
			verticalCmds, horizontalCmds))

		// // termboxの初期化
		// if err := termbox.Init(); err != nil {
		// 	panic(err)
		// }
		// defer termbox.Close()
		// termbox.SetInputMode(termbox.InputEsc)
		// termbox.Flush()
		//
		// // 端末の幅を取得
		// w, h := termbox.Size()
		//
		// vcnt := len(verticalCmds)
		// hcnt := len(horizontalCmds)
		//
		// for i, v := range verticalCmds {
		// 	p := Pane{
		// 		Name: v,
		// 		X:    w / vcnt * i,
		// 		Y:    h / hcnt * 0,
		// 	}
		// }
		//
		// // どちらも指定がないときだけ引数を処理
		// if len(verticalCmds) < 1 && len(horizontalCmds) < 1 {
		// 	return
		// }
		//
		// c, err := shellwords.Parse(cmdstr)
		// if err != nil {
		// 	return err
		// }
		// switch len(c) {
		// case 0:
		// 	// 空の文字列が渡された場合
		// 	return nil
		// case 1:
		// 	// コマンドのみを渡された場合
		// 	err = exec.Command(c[0]).Run()
		// default:
		// 	// コマンド+オプションを渡された場合
		// 	// オプションは可変長でexec.Commandに渡す
		// 	err = exec.Command(c[0], c[1:]...).Run()
		// }
		// if err != nil {
		// 	return err
		// }
		//
		// p1 := Pane{
		// 	Name:   "Pane1",
		// 	X:      0,
		// 	Y:      0,
		// 	Width:  w,
		// 	Height: h / 2,
		// }
		// p2 := Pane{
		// 	Name:   "Pane2",
		// 	X:      0,
		// 	Y:      h / 2,
		// 	Width:  w,
		// 	Height: h / 2,
		// }
		//
		// const fc = termbox.ColorDefault
		// const bc = termbox.ColorDefault
		//
		// offset := Offset{X: 0, Y: 1}
		// out, err := exec.Command("ls", "-l").Output()
		// if err != nil {
		// 	panic(err)
		// }
		// p1.SetHeader()
		// p1.SetText(out, offset, fc, bc)
		//
		// out, err = exec.Command("ls", "-la").Output()
		// if err != nil {
		// 	panic(err)
		// }
		// p2.SetHeader()
		// p2.SetText(out, offset, fc, bc)
		//
		// termbox.Flush()
		//
		// time.Sleep(5 * time.Second)

		log.Debug("end 'vhwatch'")
	},
}
