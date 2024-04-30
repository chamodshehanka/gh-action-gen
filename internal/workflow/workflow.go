package workflow

import (
	"fmt"
	"io/ioutil"
)

func CreateWorkflowFile(applicationType string, actionType string) error {
	filePath := ".github/workflows/" + applicationType + ".yml"

	yamlContent := `name: Go CI
on:
  push:
    branches:
      - main
      - master
  pull_request:
    branches:
      - main
      - master

jobs:
  build:
    name: Build and Test
    runs-on: ubuntu-latest

    steps:
    - name: Checkout code
      uses: actions/checkout@v2

    - name: Set up Go 1.21
      uses: actions/setup-go@v2
      with:
        go-version: 1.21

    - name: Run tests
      run: go test ./...

    - name: Run linters
      run: go vet ./...

    - name: Install dependencies
      run: go mod tidy
`
	err := ioutil.WriteFile(filePath, []byte(yamlContent), 0644)
	if err != nil {
		return err
	}

	fmt.Println("Config file created successfully at: ", filePath)
	return nil
}
