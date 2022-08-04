/*
Copyright © 2022 KAI CHU CHUNG <cage.chung@gmail.com>

*/
package cmd

import (
	"os"

	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/cage1016/ak/fs"
	template "github.com/cage1016/ak/templates"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "ak",
	Short: "A generator for awgo that helps you create boilerplate code",
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	rootCmd.PersistentFlags().Bool("testing", false, "If testing the generator.")
	rootCmd.PersistentFlags().BoolP("debug", "d", false, "If you want to se the debug logs.")
	rootCmd.PersistentFlags().BoolP("force", "f", false, "Force overwrite existing files without asking.")
	rootCmd.PersistentFlags().String("folder", "", "If you want to specify the base folder of the workflow.")
	viper.BindPFlag("ak_testing", rootCmd.PersistentFlags().Lookup("testing"))
	viper.BindPFlag("ak_folder", rootCmd.PersistentFlags().Lookup("folder"))
	viper.BindPFlag("ak_force", rootCmd.PersistentFlags().Lookup("force"))
	viper.BindPFlag("ak_debug", rootCmd.PersistentFlags().Lookup("debug"))
}

func initConfig() {
	initViperDefaults()
	viper.SetFs(fs.NewDefaultFs("").Fs)
	viper.SetConfigFile("ak.json")
	if viper.GetBool("ak_debug") {
		logrus.SetLevel(logrus.DebugLevel)
	} else {
		logrus.SetLevel(logrus.InfoLevel)
	}
	if err := viper.ReadInConfig(); err == nil {
		logrus.Debug("Using config file:", viper.ConfigFileUsed())
	} else {
		logrus.Info("No config file found initializing the project with the default config file.")
		te := template.NewEngine()
		st, err := te.Execute("ak.json", nil)
		if err != nil {
			logrus.Panic(err)
		}
		err = fs.Get().WriteFile("ak.json", st, false)
		if err != nil {
			logrus.Panic(err)
		}
		initConfig()
	}
}

func initViperDefaults() {
	viper.SetDefault("go_mod_package", "github.com/cage1016/aa")

	viper.SetDefault("workflow.folder", ".workflow")
	viper.SetDefault("workflow.name", "aa")
	viper.SetDefault("workflow.category", "")
	viper.SetDefault("workflow.description", "")
	viper.SetDefault("workflow.bundle_id", "com.kaichu.aa")
	viper.SetDefault("workflow.created_by", "")
	viper.SetDefault("workflow.web_address", "")
	viper.SetDefault("workflow.version", "0.1.0")

	viper.SetDefault("update.github_repo", "https://github.com/cage1016/aa")

	viper.SetDefault("license.type", "mit")
	viper.SetDefault("license.year", "2022")
	viper.SetDefault("license.name", "")
}
