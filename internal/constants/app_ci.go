package constants

import "github.com/charmbracelet/huh"

type Option[T comparable] struct {
	Key      string
	Value    T
	selected bool
}

func GetAppCIOptions() []huh.Option[string] {
	return []huh.Option[string]{
		huh.NewOption("Golang", "golang"),
		huh.NewOption("Node.JS", "nodejs"),
		huh.NewOption("React.JS", "react"),
		huh.NewOption(".NET", "dotnet"),
		huh.NewOption("Angular", "angular"),
	}
}
