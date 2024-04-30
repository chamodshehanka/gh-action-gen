package options

import (
	"github.com/chamodshehanka/gh-action-gen/internal/constants"
	"github.com/charmbracelet/huh"
)

func GetActionTypeOptions(actionType string) []huh.Option[string] {
	switch actionType {
	case constants.ActionTypeAppCI:
		return GetAppCIOptions()
	case constants.ActionTypeDeployment:
		return nil
	case constants.ActionTypePublish:
		return nil
	default:
		return nil
	}
}
