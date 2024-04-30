package init

import (
	"fmt"
	"github.com/chamodshehanka/gh-action-gen/internal/constants"
	"github.com/chamodshehanka/gh-action-gen/internal/directory"
	"github.com/chamodshehanka/gh-action-gen/internal/options"
	"github.com/chamodshehanka/gh-action-gen/internal/utils"
	"github.com/chamodshehanka/gh-action-gen/internal/workflow"
	"github.com/charmbracelet/huh"
	"github.com/charmbracelet/lipgloss"
	"github.com/charmbracelet/log"
)

var (
	actionType      string
	applicationType string
	actionTriggers  []string
	name            string
	branch          string
)

func LoadInitForm() {
	form := huh.NewForm(
		huh.NewGroup(
			huh.NewSelect[string]().
				Title("Choose the GitHub Action type you want to create:").
				Options(
					huh.NewOption(constants.ActionTypeAppCI, constants.ActionTypeAppCI),
					huh.NewOption(constants.ActionTypeDeployment, constants.ActionTypeDeployment),
					huh.NewOption(constants.ActionTypePublish, constants.ActionTypePublish),
				).
				Value(&actionType),
		),
	)

	err := form.Run()
	if err != nil {
		log.Fatal(err)
	}

	style := lipgloss.NewStyle().
		Foreground(lipgloss.Color("205")).
		Bold(true).PaddingLeft(2)

	// Use the style
	fmt.Println(style.Render("GitHub Action type: "), actionType)

	actionTypeOptions := options.GetActionTypeOptions(actionType)
	if actionTypeOptions == nil {
		log.Fatal("Not implemented yet!")
		return
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

	fmt.Println(style.Render("Application technology: "), applicationType)

	form = huh.NewForm(
		huh.NewGroup(
			huh.NewMultiSelect[string]().Title("When should this action be triggered?").Options(
				huh.NewOption("On a new commit", constants.TriggerOnCommit),
				huh.NewOption("On a Pull Request", constants.TriggerOnPullRequest),
				huh.NewOption("On Workflow Dispatch", constants.TriggerOnWorkflowDispatch),
			).Value(&actionTriggers),

			huh.NewInput().Title("What's the branch name? (default 'main' will be used)").Value(&branch),
		),
	)
	err = form.Run()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(style.Render("GitHub Action will be triggered on: "), utils.StringArrayToString(actionTriggers))
	err = createWorkflowFile()
	if err != nil {
		log.Fatal(err)
		return
	}

}

func createWorkflowFile() error {
	if directory.IsWorkflowDirExists() {
		err := workflow.CreateWorkflowFile(applicationType, actionType)
		if err != nil {
			return err
		}
	} else {
		err := directory.CreateWorkflowDir()
		if err != nil {
			return err
		}

		err = createWorkflowFile()
		if err != nil {
			return err
		}
	}
	return nil
}
