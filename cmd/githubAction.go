/*
Copyright © 2022 KAI CHU CHUNG <cage.chung@gmail.com>
*/
package cmd

import (
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"

	"github.com/cage1016/ak/generator"
)

// githubActionCmd represents the githubAction command
var githubActionCmd = &cobra.Command{
	Use:     "githubAction",
	Short:   "Add Github Action release to project",
	Aliases: []string{"ga"},
	Run: func(cmd *cobra.Command, args []string) {
		s, _ := cmd.Flags().GetBool("sign")
		c, _ := cmd.Flags().GetBool("codecov")
		g, _ := cmd.Flags().GetBool("golang")

		if err := generator.NewGithubActionGenerator(
			generator.WithEnabled_Code_Sign_Notarize(s),
			generator.WithEnabled_Codecov(c),
			generator.WithEnabled_Golang(g),
		).Generate(); err != nil {
			logrus.Fatal(err)
		}
	},
}

func init() {
	addCmd.AddCommand(githubActionCmd)
	githubActionCmd.PersistentFlags().BoolP("sign", "s", false, "enable code sign and notarize")
	githubActionCmd.PersistentFlags().BoolP("codecov", "c", false, "enable codecov")
	githubActionCmd.PersistentFlags().BoolP("golang", "g", false, "enable goland")
}
