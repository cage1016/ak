/*
Copyright © 2022 KAI CHU CHUNG <cage.chung@gmail.com>
*/
package cmd

import (
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"

	"github.com/cage1016/ak/generator"
)

// varsArgsCmd represents the simple command
var varsArgsCmd = &cobra.Command{
	Use:     "varsArgs",
	Short:   "create a workflow with variables and arguments",
	Aliases: []string{"v"},
	Run: func(cmd *cobra.Command, args []string) {
		if err := generator.NewVarsArgsGenerator().Generate(); err != nil {
			logrus.Fatal(err)
		}
	},
}

func init() {
	newCmd.AddCommand(varsArgsCmd)
}
