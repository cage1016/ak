/*
Copyright © 2022 KAI CHU CHUNG <cage.chung@gmail.com>

*/
package cmd

import (
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"

	"github.com/cage1016/ak/generator"
)

// itemsCmd represents the items command
var itemsCmd = &cobra.Command{
	Use:     "items",
	Short:   "create a workflow with items feedback",
	Aliases: []string{"i"},
	Run: func(cmd *cobra.Command, args []string) {
		if err := generator.NewItemsGenerator().Generate(); err != nil {
			logrus.Fatal(err)
		}
	},
}

func init() {
	newCmd.AddCommand(itemsCmd)
}
