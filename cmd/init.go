/*
Copyright © 2022 KAI CHU CHUNG <cage.chung@gmail.com>

*/
package cmd

import (
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"

	"github.com/cage1016/ak/generator"
)

// initCmd represents the init command
var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initiates a workflow",
	Run:   runInit,
}

func runInit(cmd *cobra.Command, args []string) {
	if err := generator.NewPlistGenerator().Generate(); err != nil {
		logrus.Fatal(err)
	}
}

func init() {
	rootCmd.AddCommand(initCmd)
}
