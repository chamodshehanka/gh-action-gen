package directory

import (
	"fmt"
	"os"
)

func IsWorkflowDirExists() bool {
	_, err := os.Stat(".github/workflows")
	if os.IsNotExist(err) {
		return false
	}
	return true
}

func CreateWorkflowDir() error {
	dirPath := ".github"

	// Create the directory with 0777 permissions
	err := os.Mkdir(dirPath, 0777)
	if err != nil {
		return err
	}

	dirPath = ".github/workflows"

	// Create the directory with 0777 permissions
	err = os.Mkdir(dirPath, 0777)
	if err != nil {
		return err
	}

	fmt.Println("Directory created successfully:", dirPath)

	// return error or success
	return nil
}
