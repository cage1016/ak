/*
Copyright © 2022 KAI CHU CHUNG <cage.chung@gmail.com>
*/
package cmd

import (
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"

	"github.com/cage1016/ak/alfred"
)

// packCmd represents the pack command
var packCmd = &cobra.Command{
	Use:   "pack",
	Short: "Package the workflow for distribution",
	Run: func(cmd *cobra.Command, args []string) {
		if err := alfred.NewAlfred().Pack(); err != nil {
			logrus.Fatal(err)
		}
	},
}

func init() {
	alfredCmd.AddCommand(packCmd)
}
