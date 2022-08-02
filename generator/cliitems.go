package generator

import (
	"fmt"
	"strings"

	"github.com/Songmu/prompter"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"

	"github.com/cage1016/ak/alfred"
	"github.com/cage1016/ak/fs"
	template "github.com/cage1016/ak/templates"
)

type CliItemsGenerator struct {
	Plist alfred.Plist
}

func (ig *CliItemsGenerator) Generate() error {
	te := template.NewEngine()
	defaultFs := fs.Get()

	if b, _ := defaultFs.Exists(goMod); b && !viper.GetBool("ak_force") {
		b := prompter.YN(fmt.Sprintf("`%s` already exists do you want to override it ?", goMod), false)
		if b {
			alfred.Run("rm", "-f", goMod, goSum)
			alfred.Run("go", "mod", "init", viper.GetString("go_mod_package"))
		}
	} else {
		alfred.Run("rm", "-f", goMod, goSum)
		alfred.Run("go", "mod", "init", viper.GetString("go_mod_package"))
	}

	{
		// generate main.go
		m, err := te.Execute("cliitems.main", map[string]interface{}{
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
	}

	{
		// generate cmd/root.go
		m, err := te.Execute("cliitems.root", map[string]interface{}{
			"GithubRepo":  strings.Replace(viper.GetString("go_mod_package"), "github.com/", "", 1),
			"Name":        ig.Plist["name"].(string),
			"Description": ig.Plist["description"].(string),
			"Year":        viper.GetString("license.year"),
			"Author":      viper.GetString("license.name"),
		})
		if err != nil {
			return err
		}

		err = defaultFs.MkdirAll("cmd")
		logrus.Debug("Creating \"cmd\"folder")
		if err != nil {
			return err
		}

		err = fs.NewDefaultFs("cmd").WriteFile("root.go", m, viper.GetBool("ak_force"))
		if err != nil {
			return err
		}

		m, err = te.Execute("cliitems.update", map[string]interface{}{
			"Name":   ig.Plist["name"].(string),
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
	}

	// license
	o, err := alfred.RunWithOutput("license", "-year", viper.GetString("license.year"), "-name", viper.GetString("license.name"), viper.GetString("license.type"))
	if err != nil {
		return err
	}
	err = defaultFs.WriteFile("LICENSE", o, viper.GetBool("ak_force"))
	if err != nil {
		return err
	}

	return nil
}

func NewCliItemsGenerator(plist alfred.Plist) *CliItemsGenerator {
	return &CliItemsGenerator{plist}
}
