package generator

import (
	"strings"

	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"

	"github.com/cage1016/ak/fs"
	template "github.com/cage1016/ak/templates"
)

type ScriptFilterGenerator struct {
	EnabledAutoUpdate bool
}

func (ig *ScriptFilterGenerator) Generate() error {
	te := template.NewEngine()
	defaultFs := fs.Get()

	// workflow folder
	VerifyWorkflowFolder()

	// go mod
	GoModGenerator()

	{
		// generate main.go
		m, err := te.Execute("scriptFilter.main", map[string]interface{}{
			"GoModPackage": viper.GetString("go_mod_package"),
			"Year":         viper.GetString("license.year"),
			"Author":       viper.GetString("license.name"),
		})
		if err != nil {
			return err
		}

		err = defaultFs.WriteFile("main.go", m, viper.GetBool("ak_force"))
		if err != nil {
			return err
		}
		logrus.Debugf("generating main.go")
	}

	{
		// generate cmd/root.go
		m, err := te.Execute("scriptFilter.root", map[string]interface{}{
			"EnabledAutoUpdate": ig.EnabledAutoUpdate,
			"GithubRepo":        strings.Replace(viper.GetString("go_mod_package"), "github.com/", "", 1),
			"Name":              viper.GetString("workflow.name"),
			"Description":       viper.GetString("workflow.description"),
			"Year":              viper.GetString("license.year"),
			"Author":            viper.GetString("license.name"),
		})
		if err != nil {
			return err
		}

		err = defaultFs.MkdirAll("cmd")
		if err != nil {
			return err
		}
		logrus.Debug("creating cmd folder")

		err = fs.NewDefaultFs("cmd").WriteFile("root.go", m, viper.GetBool("ak_force"))
		if err != nil {
			return err
		}
		logrus.Debugf("generating cmd/root.go")

		if ig.EnabledAutoUpdate {
			m, err = te.Execute("scriptFilter.update", map[string]interface{}{
				"Name":   viper.GetString("workflow.name"),
				"Year":   viper.GetString("license.year"),
				"Author": viper.GetString("license.name"),
			})
			if err != nil {
				return err
			}

			err = fs.NewDefaultFs("cmd").WriteFile("update.go", m, viper.GetBool("ak_force"))
			if err != nil {
				return err
			}
			logrus.Debugf("generating cmd/update.go")
		}
	}

	{
		if ig.EnabledAutoUpdate {
			// update-available.png
			err := fs.NewDefaultFs(".workflow").WriteFile("update-available.png", te.MustAssetString("icons/update-available.png"), viper.GetBool("ak_force"))
			if err != nil {
				return err
			}
			logrus.Debugf("generating update-available.png")
		}
	}

	return nil
}

func NewScriptFilterGenerator(e bool) *ScriptFilterGenerator {
	return &ScriptFilterGenerator{
		EnabledAutoUpdate: e,
	}
}
