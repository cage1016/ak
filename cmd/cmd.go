/*
Copyright © 2022 KAI CHU CHUNG <cage.chung@gmail.com>
*/
package cmd

import (
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"

	"github.com/cage1016/ak/generator"
)

// cmdCmd represents the cobra command
var cmdCmd = &cobra.Command{
	Use:     "cmd",
	Short:   "create cobra command",
	Aliases: []string{"c"},
	Run: func(cmd *cobra.Command, args []string) {
		if err := generator.NewCmdGenerator().Generate(); err != nil {
			logrus.Fatal(err)
		}
	},
}

func init() {
	newCmd.AddCommand(cmdCmd)
}
