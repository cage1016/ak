/*
Copyright © 2022 KAI CHU CHUNG <cage.chung@gmail.com>
*/
package cmd

import (
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"

	"github.com/cage1016/ak/generator"
)

// scriptFilterCmd represents the cobra command
var scriptFilterCmd = &cobra.Command{
	Use:     "scriptFilter",
	Short:   "create scriptFilter items feedback",
	Aliases: []string{"c"},
	Run: func(cmd *cobra.Command, args []string) {
		if err := generator.NewScriptFilterGenerator().Generate(); err != nil {
			logrus.Fatal(err)
		}
	},
}

func init() {
	newCmd.AddCommand(scriptFilterCmd)
}
