/*
Copyright © 2022 KAI CHU CHUNG <cage.chung@gmail.com>

*/
package cmd

import (
	"github.com/spf13/cobra"
)

// alfredCmd represents the alfred command
var alfredCmd = &cobra.Command{
	Use:   "alfred",
	Short: "used to manage Go-based Alfred workflows",
}

func init() {
	rootCmd.AddCommand(alfredCmd)
}
