package generator

import (
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

	err = defaultFs.WriteFile("LICENSE", data, viper.GetBool("ak_force"))
	if err != nil {
		return err
	}
	logrus.Debugf("wrote LICENSE")

	return nil
}

func NewLicenseGenerator() *LicenseGenerator {
	return &LicenseGenerator{}
}
