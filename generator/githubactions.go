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
	Enabled_Codecov            bool
	Enabled_Golang             bool
}

func WithEnabled_Code_Sign_Notarize(enabled bool) func(*GithubActionGenerator) {
	return func(g *GithubActionGenerator) {
		g.Enabled_Code_Sign_Notarize = enabled
	}
}

func WithEnabled_Codecov(enabled bool) func(*GithubActionGenerator) {
	return func(g *GithubActionGenerator) {
		g.Enabled_Codecov = enabled
	}
}

func WithEnabled_Golang(enabled bool) func(*GithubActionGenerator) {
	return func(g *GithubActionGenerator) {
		g.Enabled_Golang = enabled
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
			"ReleaseName":         "Gallery Release",
			"EnabledCodeSign":     gg.Enabled_Code_Sign_Notarize,
			"EnabledCodecov":      gg.Enabled_Codecov,
			"EnabledGolang":       gg.Enabled_Golang,
			"WorkflowName":        fmt.Sprintf("%s_GALLERY", strings.ReplaceAll(viper.GetString("workflow.name"), " ", "")),
			"Ldflags":             fmt.Sprintf("-X %s/cmd.EnabledAutoUpdate=false", viper.GetString("go_mod_package")),
			"BundleID":            viper.GetString("workflow.bundle_id"),
			"ApplicationIdentity": viper.GetString("gon.application_identity"),
		})
		if err != nil {
			return err
		}

		err = fs.NewDefaultFs(".github/workflows").WriteFile("release-gallery.yml", m, viper.GetBool("ak_force"))
		if err != nil {
			return err
		}
		logrus.Debugf("generating release-gallery.yml")
	}

	// generate with auto update release.yml
	{
		m, err := te.Execute("release.yml", map[string]interface{}{
			"ReleaseName":         "Github Release",
			"EnabledCodeSign":     gg.Enabled_Code_Sign_Notarize,
			"EnabledCodecov":      gg.Enabled_Codecov,
			"EnabledGolang":       gg.Enabled_Golang,
			"WorkflowName":        fmt.Sprintf("%s_GITHUB", strings.ReplaceAll(viper.GetString("workflow.name"), " ", "")),
			"Ldflags":             fmt.Sprintf("-X %s/cmd.EnabledAutoUpdate=true", viper.GetString("go_mod_package")),
			"BundleID":            viper.GetString("workflow.bundle_id"),
			"ApplicationIdentity": viper.GetString("gon.application_identity"),
		})
		if err != nil {
			return err
		}

		err = fs.NewDefaultFs(".github/workflows").WriteFile("release-github.yml", m, viper.GetBool("ak_force"))
		if err != nil {
			return err
		}
		logrus.Debugf("generating release-github.yml")
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
