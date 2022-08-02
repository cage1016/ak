/*
Copyright © 2022 KAI CHU CHUNG <cage.chung@gmail.com>

*/
package cmd

import (
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"

	"github.com/cage1016/ak/alfred"
	"github.com/cage1016/ak/generator"
)

// cliItems represents the cobra command
var cliItems = &cobra.Command{
	Use:     "cliItems",
	Short:   "create a workflow with cobra items feedback",
	Aliases: []string{"c"},
	Run: func(cmd *cobra.Command, args []string) {
		p := alfred.NewAlfred().GetPlist()
		if err := generator.NewCliItemsGenerator(p).Generate(); err != nil {
			logrus.Fatal(err)
		}
	},
}

func init() {
	newCmd.AddCommand(cliItems)
}
