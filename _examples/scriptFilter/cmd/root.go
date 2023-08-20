/*
Copyright © 2023 KAI CHU CHUNG
*/
package cmd

import (
	"log"
	"os"

	aw "github.com/deanishe/awgo"
	"github.com/deanishe/awgo/update"
	"github.com/spf13/cobra"
)

var (
	repo = "cage1016/ak-scriptFilter-example"
	wf   *aw.Workflow
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "AK script example",
	Short: "",
	Run: func(cmd *cobra.Command, args []string) {

		wf.SendFeedback()
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	wf.Run(func() {
		if err := rootCmd.Execute(); err != nil {
			log.Println(err)
			os.Exit(1)
		}
	})
}

func init() {
	wf = aw.New(update.GitHub(repo), aw.HelpURL(repo+"/issues"))
	wf.Args() // magic for "workflow:update"
}
