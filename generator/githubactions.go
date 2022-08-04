package generator

import (
	"strings"

	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"

	"github.com/cage1016/ak/fs"
	template "github.com/cage1016/ak/templates"
)

type GithubActionGenerator struct{}

func (gg *GithubActionGenerator) Generate() error {
	te := template.NewEngine()
	defaultFs := fs.Get()

	m, err := te.Execute("release.yml", map[string]interface{}{
		"WorkflowName": strings.ReplaceAll(viper.GetString("workflow.name"), " ", ""),
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

func NewGithubActionGenerator() *GithubActionGenerator {
	return &GithubActionGenerator{}
}
