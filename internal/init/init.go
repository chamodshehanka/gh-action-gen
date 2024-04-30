package init

import (
	"fmt"
	"github.com/chamodshehanka/gh-action-gen/internal/constants"
	"github.com/chamodshehanka/gh-action-gen/internal/directory"
	"github.com/chamodshehanka/gh-action-gen/internal/workflow"
	"github.com/charmbracelet/huh"
	"log"
)

var (
	actionType      string
	applicationType string
	extensions      []string
	name            string
)

func LoadInitForm() {
	form := huh.NewForm(
		huh.NewGroup(
			// Ask the user for the type of GitHub Action they want to create.
			huh.NewSelect[string]().
				Title("Choose your GitHub Action Type").
				Options(
					huh.NewOption("App CI", "ci"),
					huh.NewOption("Deployment", "deployment"),
					huh.NewOption("Publish", "publish"),
				).
				Value(&actionType), // store the chosen option in the "actionType" variable
		),
	)

	err := form.Run()
	if err != nil {
		log.Fatal(err)
	}

	actionTypeOptions := constants.GetActionTypeOptions(actionType)
	if actionTypeOptions == nil {
		log.Fatal("Invalid action type")
	}

	form = huh.NewForm(
		huh.NewGroup(
			huh.NewSelect[string]().
				Title("Choose your application technology").
				Options(actionTypeOptions...).Value(&applicationType),
		),
	)

	err = form.Run()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Application Type: ", applicationType)
	fmt.Println("Action Type: ", actionType)

	err = createWorkflowFile()
	if err != nil {
		log.Fatal(err)
		return
	}

	err = workflow.CreateWorkflowFile(applicationType, actionType)
	if err != nil {
		log.Fatal(err)
		return
	}

}

func createWorkflowFile() error {
	if directory.IsWorkflowDirExists() {

	} else {
		err := directory.CreateWorkflowDir()
		if err != nil {
			return err
		}
	}
	return nil
}
