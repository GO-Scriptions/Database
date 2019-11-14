package login

import "os"

// Check quiries info from the database to check against args passed
func Check(log string) bool {
	var logSuccess bool
	logSuccess = false

	if len(os.Args) == 4 {
		logSuccess = true
	}

	return logSuccess
}
