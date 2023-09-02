/*
Copyright © 2022 KAI CHU CHUNG <cage.chung@gmail.com>
*/
package cmd

import (
	"strings"

	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"

	"github.com/cage1016/ak/alfred"
)

// buildCmd represents the build command
var buildCmd = &cobra.Command{
	Use:   "build",
	Short: "Build the workflow executable and output it into the \".workflow\" subdirectory",
	Run: func(cmd *cobra.Command, args []string) {
		l, _ := cmd.Flags().GetString("ldflags")

		a := alfred.NewAlfred()
		if l != "" {
			a = alfred.NewAlfred(alfred.WithLdflags(strings.Split(l, " ")...))
		}

		if err := a.Build(); err != nil {
			logrus.Fatal(err)
		}
	},
}

func init() {
	alfredCmd.AddCommand(buildCmd)
	buildCmd.PersistentFlags().StringP("ldflags", "l", "", "ldflags")
}
