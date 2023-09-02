/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"

	"github.com/cage1016/ak/generator"
)

// makefileCmd represents the makefile command
var makefileCmd = &cobra.Command{
	Use:     "makefile",
	Short:   "Add Makefile to project",
	Aliases: []string{"mf"},
	Run: func(cmd *cobra.Command, args []string) {
		if err := generator.NewMakefileGenerator().Generate(); err != nil {
			logrus.Fatal(err)
		}
	},
}

func init() {
	addCmd.AddCommand(makefileCmd)
}
