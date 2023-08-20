/*
Copyright © 2022 KAI CHU CHUNG <cage.chung@gmail.com>
*/
package cmd

import (
	"github.com/spf13/cobra"

	"github.com/cage1016/ak/alfred"
)

// infoCmd represents the info command
var infoCmd = &cobra.Command{
	Use:   "info",
	Short: "Display information about the workflow",
	Run: func(cmd *cobra.Command, args []string) {
		alfred.NewAlfred().Info()
	},
}

func init() {
	alfredCmd.AddCommand(infoCmd)
}
