/*
Copyright © 2022 KAI CHU CHUNG <cage.chung@gmail.com> (https://kaichu.io)

*/
package main

import (
	"errors"
	"log"
	"os"
	"os/exec"

	aw "github.com/deanishe/awgo"
	"github.com/deanishe/awgo/update"
	_ "github.com/joho/godotenv/autoload"
)

const updateJobName = "checkForUpdate"

var (
	repo = "cage1016/ak-test"
	wf   *aw.Workflow
)

func init() {
	wf = aw.New(update.GitHub(repo), aw.HelpURL(repo+"/issues"))
}

func run() {
	args := wf.Args()
	if len(args) == 0 {
		wf.FatalError(errors.New("please provide some input 👀"))
	}

	handlers := map[string]func(*aw.Workflow, []string){
		"update": func(wf *aw.Workflow, _ []string) {
			wf.Configure(aw.TextErrors(true))
			log.Println("Checking for updates...")
			if err := wf.CheckForUpdate(); err != nil {
				wf.FatalError(err)
			}

			wf.SendFeedback()
		},
	}

	if wf.UpdateCheckDue() && !wf.IsRunning(updateJobName) {
		log.Println("Running update check in background...")
		cmd := exec.Command(os.Args[0], "update")
		if err := wf.RunInBackground(updateJobName, cmd); err != nil {
			log.Printf("Error starting update check: %s", err)
		}
	}

	if wf.UpdateAvailable() {
		wf.Configure(aw.SuppressUIDs(true))
		log.Println("Update available!")
		wf.NewItem("An update is available!").
			Subtitle("⇥ or ↩ to install update").
			Valid(false).
			Autocomplete("workflow:update").
			Icon(&aw.Icon{Value: "update-available.png"})
	}

	h, found := handlers[args[0]]
	if !found {
		wf.FatalError(errors.New("command not recognized 👀"))
	}

	h(wf, args[1:])
}

func main() {
	wf.Run(run)
}
