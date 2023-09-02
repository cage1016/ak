package generator

import (
	"fmt"
	"strings"

	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"

	"github.com/cage1016/ak/fs"
	"github.com/cage1016/ak/template"
)

// type GithubActionGeneratorOptions struct {
// 	Enabled_Code_Sign_Notarize bool
// }

type GithubActionGenerator struct {
	Enabled_Code_Sign_Notarize bool
}

func WithEnabled_Code_Sign_Notarize(enabled bool) func(*GithubActionGenerator) {
	return func(g *GithubActionGenerator) {
		g.Enabled_Code_Sign_Notarize = enabled
	}
}

func (gg *GithubActionGenerator) Generate() error {
	te := template.NewEngine()
	defaultFs := fs.Get()

	err := defaultFs.MkdirAll(".github/workflows")
	if err != nil {
		return err
	}
	logrus.Debug("Creating \".github/workflows\"folder")

	// generate release.yml
	{
		m, err := te.Execute("release.yml", map[string]interface{}{
			"EnabledCodeSign":     gg.Enabled_Code_Sign_Notarize,
			"WorkflowName":        strings.ReplaceAll(viper.GetString("workflow.name"), " ", ""),
			"BundleID":            viper.GetString("workflow.bundle_id"),
			"ApplicationIdentity": viper.GetString("gon.application_identity"),
		})
		if err != nil {
			return err
		}

		err = fs.NewDefaultFs(".github/workflows").WriteFile("release.yml", m, viper.GetBool("ak_force"))
		if err != nil {
			return err
		}
		logrus.Debugf("generating release.yml")
	}

	// generate with auto update release.yml
	{
		m, err := te.Execute("release.yml", map[string]interface{}{
			"EnabledCodeSign":     gg.Enabled_Code_Sign_Notarize,
			"WorkflowName":        fmt.Sprintf("%s_auto_update", strings.ReplaceAll(viper.GetString("workflow.name"), " ", "")),
			"Ldflags":             fmt.Sprintf("-X %s/cmd.EnabledAutoUpdate=true", viper.GetString("go_mod_package")),
			"BundleID":            viper.GetString("workflow.bundle_id"),
			"ApplicationIdentity": viper.GetString("gon.application_identity"),
		})
		if err != nil {
			return err
		}

		err = fs.NewDefaultFs(".github/workflows").WriteFile("release_auto_update.yml", m, viper.GetBool("ak_force"))
		if err != nil {
			return err
		}
		logrus.Debugf("generating release_auto_update.yml")
	}

	return nil
}

func NewGithubActionGenerator(options ...func(*GithubActionGenerator)) *GithubActionGenerator {
	generator := &GithubActionGenerator{}

	for _, o := range options {
		o(generator)
	}

	return generator
}
