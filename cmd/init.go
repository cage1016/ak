/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	survey "github.com/AlecAivazis/survey/v2"
	"github.com/cage1016/ak/generator"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// initCmd represents the init command
var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initiates a workflow",
	Run:   runInit,
}

// the questions to ask
var qs = []*survey.Question{
	{
		Name:      "name",
		Prompt:    &survey.Input{Message: "What is alfred workflow name?"},
		Validate:  survey.Required,
		Transform: survey.Title,
	},
	{
		Name:      "description",
		Prompt:    &survey.Input{Message: "What is alfred workflow description?"},
		Transform: survey.Title,
	},
	{
		Name: "category",
		Prompt: &survey.Select{
			Message: "Choose a Category:",
			Options: []string{"Tools", "Internet", "Productivity", "Uncategorised"},
			Default: "Uncategorised",
		},
	},
	{
		Name:      "bundle_id",
		Prompt:    &survey.Input{Message: "What is alfred workflow bundle ID?"},
		Validate:  survey.Required,
		Transform: survey.Title,
	},
	{
		Name:      "created_by",
		Prompt:    &survey.Input{Message: "What alfred workflow create by?"},
		Transform: survey.Title,
	},
	{
		Name:      "web_address",
		Prompt:    &survey.Input{Message: "What is alfred workflow web address?"},
		Transform: survey.Title,
	},
	{
		Name:      "version",
		Prompt:    &survey.Input{Message: "What is alfred workflow version?"},
		Transform: survey.Title,
	},
}

func runInit(cmd *cobra.Command, args []string) {
	answers := generator.Workflow{}
	err := survey.Ask(qs, &answers)
	if err != nil {
		logrus.Fatal(err)
	}

	if err := generator.NewPlistGenerator().Generate(answers); err != nil {
		logrus.Fatal(err)
	}
}

func init() {
	rootCmd.AddCommand(initCmd)
}
