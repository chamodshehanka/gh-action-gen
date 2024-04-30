package workflow

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"io/ioutil"
	"strings"
)

func CreateWorkflowFile(applicationType string, actionType string) error {
	filePath, err := generateFileName(applicationType, actionType)
	if err != nil {
		return err
	}

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
      uses: actions/checkout@v4

    - name: Set up Go 1.21
      uses: actions/setup-go@v5
      with:
        go-version: 1.21

    - name: Run tests
      run: go test ./...

    - name: Run linters
      run: go vet ./...

    - name: Install dependencies
      run: go mod tidy
`
	err = ioutil.WriteFile(filePath, []byte(yamlContent), 0644)
	if err != nil {
		return err
	}

	fmt.Println("Config file created successfully at: ", filePath)
	return nil
}

func generateFileName(applicationType string, actionType string) (string, error) {
	filePath := ".github/workflows/" + applicationType + "-" + actionType

	filePath = strings.ReplaceAll(filePath, " ", "-")
	filePath = strings.ReplaceAll(filePath, "(", "")
	filePath = strings.ReplaceAll(filePath, ")", "")
	filePath = strings.ReplaceAll(filePath, "", "")
	filePath = strings.ToLower(filePath)

	randomString, err := generateRandomString(8)
	if err != nil {
		return "", err
	}

	filePath = filePath + "-" + randomString + ".yml"

	return filePath, nil
}

func generateRandomString(length int8) (string, error) {
	bytes := make([]byte, length)
	if _, err := rand.Read(bytes); err != nil {
		return "", err
	}
	return hex.EncodeToString(bytes)[:length], nil
}
