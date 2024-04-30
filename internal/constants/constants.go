package constants

import "github.com/charmbracelet/huh"

const (
	AppCI      = "ci"
	Deployment = "deployment"
	Publish    = "publish"
)

func GetActionTypeOptions(actionType string) []huh.Option[string] {
	switch actionType {
	case AppCI:
		return GetAppCIOptions()
	case Deployment:
		return nil
	case Publish:
		return nil
	default:
		return nil
	}
}
