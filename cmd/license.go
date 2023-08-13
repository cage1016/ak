/*
Copyright © 2022 KAI CHU CHUNG <cage.chung@gmail.com>
*/
package cmd

import (
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"

	"github.com/cage1016/ak/generator"
)

// licenseCmd represents the license command
var licenseCmd = &cobra.Command{
	Use:   "license",
	Short: "Add license to project",
	Run: func(cmd *cobra.Command, args []string) {
		if err := generator.NewLicenseGenerator().Generate(); err != nil {
			logrus.Fatal(err)
		}
	},
}

func init() {
	addCmd.AddCommand(licenseCmd)
}
