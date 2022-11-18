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

		if err := generator.NewGithubActionGenerator(&generator.GithubActionGeneratorOptions{
			Enabled_Code_Sign_Notarize: s,
		}).Generate(); err != nil {
			logrus.Fatal(err)
		}
	},
}

func init() {
	addCmd.AddCommand(githubActionCmd)
	githubActionCmd.PersistentFlags().BoolP("sign", "s", false, "enable code sign and notarize")
}
