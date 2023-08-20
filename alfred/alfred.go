package alfred

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"os/user"
	"path"
	"path/filepath"
	"strings"

	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func Run(cmd string, args ...string) {
	if output, err := exec.Command(cmd, args...).CombinedOutput(); err != nil {
		println(string(output))
		logrus.Fatal(err)
	}
}

func RunWithOutput(cmd string, args ...string) (string, error) {
	if output, err := exec.Command(cmd, args...).CombinedOutput(); err != nil {
		return string(output), err
	} else {
		return string(output), nil
	}
}

func RunIfFile(file, cmd string, args ...string) {
	if _, err := os.Stat(file); err == nil {
		Run(cmd, args...)
	}
}

func dirExists(dir string) bool {
	stat, err := os.Stat(dir)
	if err != nil {
		return os.IsExist(err)
	}
	return stat.IsDir()
}

func fileExists(file string) bool {
	stat, err := os.Stat(file)
	if err != nil {
		return os.IsExist(err)
	}
	return !stat.IsDir()
}

// getAlfredVersion returns the highest installed version of Alfred. It uses a very naive algorithm.
func getAlfredVersion() string {
	files, _ := ioutil.ReadDir("/Applications")
	name := ""
	for _, file := range files {
		fname := file.Name()
		if strings.HasPrefix(fname, "Alfred ") && fname > name {
			name = fname
			break
		}
	}
	if name != "" {
		name = strings.TrimSuffix(name, ".app")
		parts := strings.Split(name, " ")
		if len(parts) == 2 {
			return parts[1]
		}
	}
	return ""
}

func getPrefsDirectory() string {
	currentUser, _ := user.Current()
	version := getAlfredVersion()

	prefFile := path.Join(currentUser.HomeDir, "Library", "Preferences", "com.runningwithcrayons.Alfred-Preferences.plist")
	preferences := LoadPlist(prefFile)

	var folder string
	if preferences["syncfolder"] != nil && preferences["syncfolder"] != "" {
		folder = preferences["syncfolder"].(string)
		if strings.HasPrefix(folder, "~/") {
			folder = path.Join(currentUser.HomeDir, folder[2:])
		}
	} else {
		if version != "4" {
			folder = path.Join(currentUser.HomeDir, "Library", "Application Support", "Alfred "+version)
		} else {
			folder = path.Join(currentUser.HomeDir, "Library", "Application Support", "Alfred")
		}
	}

	var info os.FileInfo
	var err error
	if info, err = os.Stat(folder); err != nil {
		logrus.Infof("folder %s does not exist, creating it", folder)
		if err := os.MkdirAll(folder, 0755); err != nil {
			logrus.Fatalf("creating folder: %s", folder)
		}
		info, _ = os.Stat(folder)
	}

	if !info.IsDir() {
		logrus.Fatalf("%s is not a directory", folder)
	}

	return folder
}

type Alfred struct {
	WorkflowsPath string
	WorkflowPath  string
	BuildDir      string
	PrefsDir      string
	VersionTag    string
	ZipName       string
}

func NewAlfred() *Alfred {
	a := &Alfred{
		BuildDir: func(a, b string) string {
			pwd, _ := filepath.Abs(".")
			return path.Join(pwd, a, b)
		}(viper.GetString("ak_folder"), viper.GetString("workflow.folder")),
	}
	logrus.Debugf("build dir: %s", a.BuildDir)

	a.PrefsDir = getPrefsDirectory()
	logrus.Debugf("prefs dir: %s", a.PrefsDir)

	a.WorkflowsPath = path.Join(a.PrefsDir, "Alfred.alfredpreferences/workflows")
	a.WorkflowPath, _ = filepath.Abs(".")
	logrus.Debugf("workflows path: %s", a.WorkflowsPath)

	plistFile := path.Join(a.BuildDir, "info.plist")
	logrus.Debugf("plist file: %s", plistFile)

	if fileExists(plistFile) {
		plist := LoadPlist(plistFile)
		workflowVersion := plist["version"]
		if workflowVersion != nil {
			a.VersionTag = fmt.Sprintf("-%s", workflowVersion)
		}

		workflowName := plist["name"]
		if workflowName != nil {
			a.ZipName = fmt.Sprintf("%s%s.alfredworkflow", strings.ReplaceAll(workflowName.(string), " ", ""), a.VersionTag)
			logrus.Debugf("zipName: %s", a.ZipName)
		}
	}

	return a
}

func (a *Alfred) GetExistingLink() (string, error) {
	dir, err := os.Open(a.WorkflowsPath)
	if err != nil {
		return "", err
	}
	defer dir.Close()

	dirs, err := dir.Readdir(-1)
	if err != nil {
		return "", err
	}

	wd, _ := os.Getwd()
	buildPath := path.Join(wd, a.BuildDir)

	for _, dir := range dirs {
		if dir.Mode()&os.ModeSymlink == os.ModeSymlink {
			fullDir := path.Join(a.WorkflowsPath, dir.Name())
			link, err := filepath.EvalSymlinks(fullDir)
			if err == nil && link == buildPath {
				return fullDir, nil
			}
		}
	}

	return "", nil
}

func (a *Alfred) GetExistingInstall() (string, error) {
	dir, err := os.Open(a.WorkflowsPath)
	if err != nil {
		return "", err
	}
	defer dir.Close()

	plistFile := path.Join(a.BuildDir, "info.plist")
	info := LoadPlist(plistFile)
	id := info["bundleid"]

	dirs, err := dir.Readdir(-1)
	if err != nil {
		return "", err
	}

	for _, d := range dirs {
		infoFile := path.Join(dir.Name(), d.Name(), "info.plist")
		if !fileExists(infoFile) {
			continue
		}

		infoPlist := LoadPlist(infoFile)
		workflowID := infoPlist["bundleid"]
		if workflowID == id {
			return d.Name(), nil
		}
	}

	return "", nil
}

func (a *Alfred) Link() error {
	logrus.Printf("Linking workflow...")
	existing, err := a.GetExistingLink()
	if err != nil {
		return err
	}

	if existing != "" {
		logrus.Println("Existing link", filepath.Base(existing))
		return nil
	}

	existing, err = a.GetExistingInstall()
	if err != nil {
		return err
	}

	if existing != "" {
		plistFile := path.Join(a.WorkflowsPath, existing, "info.plist")
		logrus.Printf("Reading from plist file %s", plistFile)
		info := LoadPlist(plistFile)
		info["disabled"] = true
		SavePlist(plistFile, info)
		logrus.Println("disabled existing install at", existing)
	}

	uuidgen, _ := exec.Command("uuidgen").Output()
	uuid := strings.TrimSpace(string(uuidgen))
	target := path.Join(a.WorkflowsPath, "user.workflow."+string(uuid))
	logrus.Printf("Creating new link to target %s", target)
	buildPath := path.Join(a.WorkflowPath, a.BuildDir)
	logrus.Printf("Build path is %s", buildPath)
	Run("ln", "-s", buildPath, target)
	logrus.Println("created link", filepath.Base(target))

	return nil
}

func (a *Alfred) Unlink() error {
	logrus.Printf("Unlinking workflow...")
	existing, err := a.GetExistingLink()
	if err != nil {
		return err
	}

	if existing == "" {
		return nil
	}

	Run("rm", existing)
	logrus.Println("removed link", filepath.Base(existing))

	if existing, err = a.GetExistingInstall(); err != nil {
		return err
	}

	if existing != "" {
		plistFile := path.Join(a.WorkflowsPath, existing, "info.plist")
		info := LoadPlist(plistFile)
		info["disabled"] = false
		SavePlist(plistFile, info)
		logrus.Println("enabled existing install at", existing)
	}
	return nil
}

func (a *Alfred) Info() {
	logrus.Printf("Getting workflow info...")
	width := -15

	printField := func(name, value string) {
		logrus.Printf("%*s %s\n", width, name+":", value)
	}

	printField("Workflows", a.WorkflowsPath)

	if link, _ := a.GetExistingLink(); link != "" {
		printField("This workflow", path.Base(link))
	}

	plistFile := path.Join(a.BuildDir, "info.plist")
	info := LoadPlist(plistFile)
	printField("Version", info["version"].(string))
	printField("WebAddress", info["webaddress"].(string))
}

func (a *Alfred) Build() error {
	command := flag.NewFlagSet("build", flag.ExitOnError)
	help := command.Bool("h", false, "show this message")
	command.Parse(os.Args[2:])

	if *help {
		logrus.Printf("Showing help")
		command.PrintDefaults()
		os.Exit(0)
	}

	logrus.Printf("Building the workflow...")

	// use go generate, along with custom build tools, to handle any auxiliary
	// build steps
	Run("go", "generate")

	cmdAmd64 := exec.Command("go", "build", "-ldflags", "-s -w", "-o", "exe_amd64")
	cmdAmd64.Env = append(os.Environ(), "GOOS=darwin", "GOARCH=amd64")
	if output, err := cmdAmd64.CombinedOutput(); err != nil {
		logrus.Println(string(output))
		logrus.Fatalf("Failed to build amd64 binary: %s", err)
	}
	cmdArm64 := exec.Command("go", "build", "-ldflags", "-s -w", "-o", "exe_arm64")
	cmdArm64.Env = append(os.Environ(), "GOOS=darwin", "GOARCH=arm64")
	if output, err := cmdArm64.CombinedOutput(); err != nil {
		logrus.Println(string(output))
		logrus.Fatalf("Failed to build arm64 binary: %s", err)
	}

	Run(
		"lipo",
		"-create",
		"-output",
		".workflow/exe",
		"exe_amd64",
		"exe_arm64",
	)

	Run("rm", "exe_amd64")
	Run("rm", "exe_arm64")
	return nil
}

func (a *Alfred) Pack() error {
	logrus.Debugf("Changing directory to %s", a.BuildDir)
	if err := os.Chdir(a.BuildDir); err != nil {
		return err
	}

	zipfile := path.Join(a.BuildDir, "..", a.ZipName)
	logrus.Infof("Creating archive %s", zipfile)
	Run("zip", "-r", zipfile, ".")

	if err := os.Chdir(path.Join(a.BuildDir, "..")); err != nil {
		return err
	}

	return nil
}
