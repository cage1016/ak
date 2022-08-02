/*
Copyright © 2022 KAI CHU CHUNG <cage.chung@gmail.com>

*/
package cmd

import (
	"github.com/spf13/cobra"
)

// newCmd represents the add command
var newCmd = &cobra.Command{
	Use:     "new",
	Short:   "Use to create workflow package",
	Aliases: []string{"n"},
}

func init() {
	rootCmd.AddCommand(newCmd)
}
