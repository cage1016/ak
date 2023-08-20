/*
Copyright © 2022 KAI CHU CHUNG <cage.chung@gmail.com>
*/
package cmd

import (
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"

	"github.com/cage1016/ak/generator"
)

// scriptCmd represents the simple command
var scriptCmd = &cobra.Command{
	Use:     "script",
	Short:   "create script feedback",
	Aliases: []string{"s"},
	Run: func(cmd *cobra.Command, args []string) {
		if err := generator.NewScriptGenerator().Generate(); err != nil {
			logrus.Fatal(err)
		}
	},
}

func init() {
	newCmd.AddCommand(scriptCmd)
}
