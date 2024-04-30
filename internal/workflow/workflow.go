package workflow

import (
	"fmt"
	"github.com/spf13/viper"
)

func CreateWorkflowFile(applicationType string, actionType string) error {
	filePath := ".github/workflows/" + applicationType + ".yml"

	viper.Set("actionType", actionType)
	viper.Set("applicationType", applicationType)
	viper.Set("filePath", filePath)
	err := viper.WriteConfigAs(filePath)
	if err != nil {
		return err
	}

	fmt.Println("Config file created successfully")
	return nil
}
