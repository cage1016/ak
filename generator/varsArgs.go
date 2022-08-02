package generator

import (
	"fmt"

	"github.com/Songmu/prompter"
	"github.com/spf13/viper"

	"github.com/cage1016/ak/alfred"
	"github.com/cage1016/ak/fs"
	template "github.com/cage1016/ak/templates"
)

const (
	goMod = "go.mod"
	goSum = "go.sum"
)

type VarsArgsGenerator struct{}

func (vg *VarsArgsGenerator) Generate() error {
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

	m, err := te.Execute("varsArgs.main", map[string]interface{}{
		"Year":   viper.GetString("license.year"),
		"Author": viper.GetString("license.name"),
	})
	if err != nil {
		return err
	}

	err = defaultFs.WriteFile("main.go", m, viper.GetBool("ak_force"))
	if err != nil {
		return err
	}

	return nil
}

func NewVarsArgsGenerator() *VarsArgsGenerator {
	return &VarsArgsGenerator{}
}
