/*
Copyright © 2022 KAI CHU CHUNG <cage.chung@gmail.com>

*/
package cmd

import (
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"

	"github.com/cage1016/ak/alfred"
)

// linkCmd represents the link command
var linkCmd = &cobra.Command{
	Use:   "link",
	Short: "Link the \".workflow\" subdirectory into Alfred's preferences directory, installing it.",
	Run: func(cmd *cobra.Command, args []string) {
		if err := alfred.NewAlfred().Link(); err != nil {
			logrus.Fatal(err)
		}
	},
}

func init() {
	alfredCmd.AddCommand(linkCmd)
}
