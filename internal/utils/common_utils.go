package utils

func StringArrayToString(array []string) string {
	triggerString := ""
	for _, trigger := range array {
		triggerString += trigger + ", "
	}

	// Remove the trailing comma and space
	if len(triggerString) > 0 {
		triggerString = triggerString[:len(triggerString)-2]
	}

	return triggerString
}
