package doc

import (
	"os"
)

// ToolNav function for navigating the doctor tools
func ToolNav(doc string) string {
	var doctool, r string
	doctool = os.Args[2]

	if doctool == "wp" {
		r = WriteP()
	} else {
		r = ViewP()
	}
	return r
}

// WriteP for writing prescriptions
func WriteP() string {
	var r string
	r = "Write prescriptions"
	return r
}

// ViewP func for Viewing prescriptions
func ViewP() string {
	var r string
	r = "View Prescriptons"
	return r
}
