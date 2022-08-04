package generator

import (
	"strings"

	"github.com/spf13/viper"

	"github.com/cage1016/ak/fs"
	template "github.com/cage1016/ak/templates"
)

type ItemsGenerator struct{}

func (ig *ItemsGenerator) Generate() error {
	te := template.NewEngine()
	defaultFs := fs.Get()

	// workflow folder
	VerifyWorkflowFolder()

	// go mod
	GoModGenerator()

	{
		// main.go
		m, err := te.Execute("items.main", map[string]interface{}{
			"GithubRepo": strings.Replace(viper.GetString("go_mod_package"), "github.com/", "", 1),
			"Year":       viper.GetString("license.year"),
			"Author":     viper.GetString("license.name"),
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
		// update-available.png
		err := fs.NewDefaultFs(".workflow").WriteFile("update-available.png", te.MustAssetString("icons/update-available.png"), viper.GetBool("ak_force"))
		if err != nil {
			return err
		}
	}

	return nil
}

func NewItemsGenerator() *ItemsGenerator {
	return &ItemsGenerator{}
}
