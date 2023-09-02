package generator

import (
	"github.com/cage1016/ak/fs"
	"github.com/cage1016/ak/template"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

type MakefileGenerator struct{}

func (mf *MakefileGenerator) Generate() error {
	te := template.NewEngine()
	defaultFs := fs.Get()

	data, err := te.Execute("makefile", map[string]interface{}{
		"GoModPackage": viper.GetString("go_mod_package"),
	})
	if err != nil {
		return err
	}

	err = defaultFs.WriteFile("Makefile", data, viper.GetBool("ak_force"))
	if err != nil {
		return err
	}
	logrus.Debugf("wrote Makefile")

	return nil
}

func NewMakefileGenerator() *MakefileGenerator {
	return &MakefileGenerator{}
}
