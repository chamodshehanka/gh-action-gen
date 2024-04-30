package constants

import "github.com/charmbracelet/huh"

func GetActionTypeOptions(actionType string) []huh.Option[string] {
	switch actionType {
	case "ci":
		return GetAppCIOptions()
	default:
		return nil
	}
}
