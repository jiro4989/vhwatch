package main

import (
	"fmt"

	"github.com/jiro4989/vhwatch/log"
	"github.com/spf13/cobra"
)

func init() {
	cobra.OnInitialize()
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

		// どちらも指定がないときだけ引数を処理
		if len(verticalCmds) < 1 && len(horizontalCmds) < 1 {
			return
		}

		log.Debug("end 'vhwatch'")
	},
}
