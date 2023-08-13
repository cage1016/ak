package generator

import (
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"

	"github.com/cage1016/ak/fs"
	template "github.com/cage1016/ak/templates"
)

const (
	plist = "info.plist"
)

type InitGenerator struct{}

func (pg *InitGenerator) Generate() error {
	te := template.NewEngine()
	defaultFs := fs.Get()

	logrus.Info("info.plist Generating...")
	pc, err := te.Execute(plist, map[string]string{
		"Name":        viper.GetString("workflow.name"),
		"Category":    viper.GetString("workflow.category"),
		"Description": viper.GetString("workflow.description"),
		"BundleID":    viper.GetString("workflow.bundle_id"),
		"CreatedBy":   viper.GetString("workflow.created_by"),
		"WebAddress":  viper.GetString("workflow.web_address"),
		"Version":     viper.GetString("workflow.version"),
	})
	if err != nil {
		return err
	}

	err = defaultFs.MkdirAll(viper.GetString("workflow.folder"))
	logrus.Debug("Creating .workflow folder")
	if err != nil {
		return err
	}

	err = fs.NewDefaultFs(viper.GetString("workflow.folder")).WriteFile(plist, pc, viper.GetBool("ak_force"))
	if err != nil {
		return err
	}

	de, err := te.Execute("dotenv", map[string]interface{}{
		"BundleID": viper.GetString("workflow.bundle_id"),
		"Version":  viper.GetString("workflow.version"),
	})
	if err != nil {
		return err
	}

	err = defaultFs.WriteFile(".env", de, viper.GetBool("ak_force"))
	if err != nil {
		return err
	}

	dt, _ := template.Asset("tmpl/.gitignore.tmpl")
	err = defaultFs.WriteFile(".gitignore", string(dt), viper.GetBool("ak_force"))
	if err != nil {
		return err
	}

	return nil
}

func NewPlistGenerator() *InitGenerator {
	return &InitGenerator{}
}
