package generator

import (
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"

	"github.com/cage1016/ak/fs"
	template "github.com/cage1016/ak/templates"
)

const (
	workflowDir = ".workflow"
	plist       = "info.plist"
)

type PlistGenerator struct{}

func (pg *PlistGenerator) Generate(answers Workflow) error {
	te := template.NewEngine()
	defaultFs := fs.Get()

	logrus.Info("info.plist Generating...")
	pc, err := te.Execute(plist, map[string]string{
		"Name":        answers.Name,
		"Category":    answers.Category,
		"Description": answers.Description,
		"BundleID":    answers.BundleID,
		"CreatedBy":   answers.CreatedBy,
		"Version":     answers.Version,
		"WebAddress":  answers.WebAddress,
	})
	if err != nil {
		return err
	}

	err = defaultFs.MkdirAll(workflowDir)
	logrus.Debug("Creating .workflow folder")
	if err != nil {
		return err
	}

	err = fs.NewDefaultFs(workflowDir).WriteFile(plist, pc, viper.GetBool("ak_force"))
	if err != nil {
		return err
	}

	return nil
}

func NewPlistGenerator() *PlistGenerator {
	return &PlistGenerator{}
}
