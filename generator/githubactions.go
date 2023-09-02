package generator

import (
	"strings"

	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"

	"github.com/cage1016/ak/fs"
	"github.com/cage1016/ak/template"
)

type GithubActionGeneratorOptions struct {
	Enabled_Code_Sign_Notarize bool
}

type GithubActionGenerator struct {
	Enabled_Code_Sign_Notarize bool
}

func (gg *GithubActionGenerator) Generate() error {
	te := template.NewEngine()
	defaultFs := fs.Get()

	m, err := te.Execute("release.yml", map[string]interface{}{
		"EnabledCodeSign":     gg.Enabled_Code_Sign_Notarize,
		"WorkflowName":        strings.ReplaceAll(viper.GetString("workflow.name"), " ", ""),
		"BundleID":            viper.GetString("workflow.bundle_id"),
		"ApplicationIdentity": viper.GetString("gon.application_identity"),
	})
	if err != nil {
		return err
	}

	err = defaultFs.MkdirAll(".github/workflows")
	if err != nil {
		return err
	}
	logrus.Debug("Creating \".github/workflows\"folder")

	err = fs.NewDefaultFs(".github/workflows").WriteFile("release.yml", m, viper.GetBool("ak_force"))
	if err != nil {
		return err
	}
	logrus.Debugf("generating release.yml")

	return nil
}

func NewGithubActionGenerator(opts *GithubActionGeneratorOptions) *GithubActionGenerator {
	return &GithubActionGenerator{
		Enabled_Code_Sign_Notarize: opts.Enabled_Code_Sign_Notarize,
	}
}
