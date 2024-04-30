package options

import "github.com/charmbracelet/huh"

func GetAppCIOptions() []huh.Option[string] {
	return []huh.Option[string]{
		huh.NewOption("Golang", "golang"),
		huh.NewOption("Node.JS", "nodejs"),
		huh.NewOption("React.JS", "react"),
		huh.NewOption(".NET", "dotnet"),
		huh.NewOption("Angular", "angular"),
	}
}
