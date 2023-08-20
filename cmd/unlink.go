/*
Copyright © 2022 KAI CHU CHUNG <cage.chung@gmail.com>
*/
package cmd

import (
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"

	"github.com/cage1016/ak/alfred"
)

// unlinkCmd represents the unlink command
var unlinkCmd = &cobra.Command{
	Use:   "unlink",
	Short: "Unlink the \".workflow\" subdirectory from Alfred's preferences directory, uninstalling it.",
	Run: func(cmd *cobra.Command, args []string) {
		if err := alfred.NewAlfred().Unlink(); err != nil {
			logrus.Fatal(err)
		}
	},
}

func init() {
	alfredCmd.AddCommand(unlinkCmd)
}
