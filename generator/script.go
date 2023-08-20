package generator

import (
	"fmt"
	"os"
	"path"
	"path/filepath"

	"github.com/Songmu/prompter"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"

	"github.com/cage1016/ak/alfred"
	"github.com/cage1016/ak/fs"
	template "github.com/cage1016/ak/templates"
)

const (
	goMod = "go.mod"
	goSum = "go.sum"
)

type ScriptGenerator struct{}

func GoModGenerator() {
	fn := func() {
		if viper.GetString("ak_folder") != "" {
			pwd, _ := filepath.Abs(".")
			if err := os.Chdir(path.Join(pwd, viper.GetString("ak_folder"))); err != nil {
				logrus.Fatalf("failed to change directory: %s", err)
			}

			alfred.Run("rm", "-f", goMod, goSum)
			logrus.Debugf("removed go mod and go sum")

			alfred.Run("go", "mod", "init", viper.GetString("go_mod_package"))
			logrus.Debugf("go mod init: %s", viper.GetString("go_mod_package"))

			if err := os.Chdir(".."); err != nil {
				logrus.Fatalf("failed to change directory: %s", err)
			}
		} else {
			alfred.Run("rm", "-f", goMod, goSum)
			logrus.Debugf("removed go mod and go sum")

			alfred.Run("go", "mod", "init", viper.GetString("go_mod_package"))
			logrus.Debugf("go mod init: %s", viper.GetString("go_mod_package"))
		}
	}

	if viper.GetString("go_mod_package") == "" {
		logrus.Fatalf("ak.json 'go_mod_package' is required")
	}

	defaultFs := fs.Get()
	if b, _ := defaultFs.Exists(goMod); b && !viper.GetBool("ak_force") {
		b := prompter.YN(fmt.Sprintf("`%s` already exists do you want to override it ?", goMod), false)
		if b {
			fn()
		}
	} else {
		fn()
	}
}

func VerifyWorkflowFolder() {
	if b, _ := fs.Get().Exists(viper.GetString("workflow.folder")); !b {
		logrus.Fatalf("workflow folder does not exist: %s", viper.GetString("workflow.folder"))
	}
}

func (vg *ScriptGenerator) Generate() error {
	te := template.NewEngine()

	m, err := te.Execute("script.main", map[string]interface{}{
		"Year":   viper.GetString("license.year"),
		"Author": viper.GetString("license.name"),
	})
	if err != nil {
		return err
	}

	err = fs.Get().WriteFile("main.go", m, viper.GetBool("ak_force"))
	if err != nil {
		return err
	}
	logrus.Debugf("generating main.go")

	// workflow folder
	VerifyWorkflowFolder()

	// go mod
	GoModGenerator()

	return nil
}

func NewScriptGenerator() *ScriptGenerator {
	return &ScriptGenerator{}
}
