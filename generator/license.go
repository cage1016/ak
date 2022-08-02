package generator

import (
	"fmt"

	"github.com/Songmu/prompter"
	"github.com/cage1016/ak/alfred"
	"github.com/cage1016/ak/fs"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

const (
	license = "LICENSE"
)

type LicenseGenerator struct{}

func (lg *LicenseGenerator) Generate() error {
	defaultFs := fs.Get()

	data, err := alfred.RunWithOutput("license", "-year", viper.GetString("license.year"), "-name", viper.GetString("license.name"), viper.GetString("license.type"))
	if err != nil {
		return err
	}

	if b, _ := defaultFs.Exists(license); b && !viper.GetBool("ak_force") {
		s, _ := defaultFs.ReadFile(license)
		if s == data {
			logrus.Warnf("`%s` exists and is identical it will be ignored", license)
			return nil
		}
		b := prompter.YN(fmt.Sprintf("`%s` already exists do you want to override it ?", license), false)
		if !b {
			return nil
		}
	}

	err = defaultFs.WriteFile("LICENSE", data, viper.GetBool("ak_force"))
	if err != nil {
		return err
	}
	return nil
}

func NewLicenseGenerator() *LicenseGenerator {
	return &LicenseGenerator{}
}
