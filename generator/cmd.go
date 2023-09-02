package generator

import (
	"fmt"
	"os"
	"path"
	"path/filepath"
	"strings"

	"github.com/Songmu/prompter"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"

	"github.com/cage1016/ak/alfred"
	"github.com/cage1016/ak/fs"
	"github.com/cage1016/ak/template"
)

const (
	goMod = "go.mod"
	goSum = "go.sum"
)

type CmdGenerator struct {
	EnabledAutoUpdate bool
}

func (ig *CmdGenerator) Generate() error {
	te := template.NewEngine()

	// generate main.go
	{
		m, err := te.Execute("cmd.main", map[string]interface{}{
			"GoModPackage": viper.GetString("go_mod_package"),
			"Year":         viper.GetString("license.year"),
			"Author":       viper.GetString("license.name"),
		})
		if err != nil {
			return err
		}

		err = fs.Get().WriteFile("main.go", m, viper.GetBool("ak_force"))
		if err != nil {
			logrus.Debugf("generating main.go, err: %v", err)
			return err
		}
		logrus.Debugf("generating main.go")
	}

	// generate cmd/root.go
	{
		m, err := te.Execute("cmd.root", map[string]interface{}{
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

		err = fs.Get().MkdirAll("cmd")
		if err != nil {
			return err
		}
		logrus.Debug("creating cmd folder")

		err = fs.Get().WriteFile("cmd/root.go", m, viper.GetBool("ak_force"))
		if err != nil {
			return err
		}
		logrus.Debugf("generating cmd/root.go")

		if ig.EnabledAutoUpdate {
			m, err = te.Execute("cmd.update", map[string]interface{}{
				"Name":   viper.GetString("workflow.name"),
				"Year":   viper.GetString("license.year"),
				"Author": viper.GetString("license.name"),
			})
			if err != nil {
				return err
			}

			err = fs.Get().WriteFile("update.go", m, viper.GetBool("ak_force"))
			if err != nil {
				return err
			}
			logrus.Debugf("generating cmd/update.go")
		}
	}

	// generate update-available.png
	{
		if ig.EnabledAutoUpdate {
			// update-available.png
			err := fs.Get().WriteFile(".workflow/update-available.png", te.MustAssetString("icons/update-available.png"), viper.GetBool("ak_force"))
			if err != nil {
				return err
			}
			logrus.Debugf("generating update-available.png")
		}
	}

	// workflow folder
	VerifyWorkflowFolder()

	// go mod
	GoModGenerator()

	return nil
}

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

func NewCmdGenerator(e bool) *CmdGenerator {
	return &CmdGenerator{
		EnabledAutoUpdate: e,
	}
}
