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
	Aliases: []string{"sf"},
	Run: func(cmd *cobra.Command, args []string) {
		e, err := cmd.Flags().GetBool("enabled-auto-update")
		if err != nil {
			logrus.Fatal(err)
		}

		if err := generator.NewScriptFilterGenerator(e).Generate(); err != nil {
			logrus.Fatal(err)
		}
	},
}

func init() {
	newCmd.AddCommand(scriptFilterCmd)
	scriptFilterCmd.Flags().BoolP("enabled-auto-update", "e", false, "enable auto update")
}
