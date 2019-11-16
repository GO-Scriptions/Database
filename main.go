package main

import (
	"flag"
	"fmt"

	"github.com/GO-Scriptions/Database/login"
)

func main() {
	log := flag.String("log", "", "Login type. Use the flags d or p login.")
	doc := flag.String("doc", "", "Doctor Tools. Use this flag with wp or vp for doctor tools")
	flag.Parse()

	// Login flag was passed, see the login folder for more details.
	if *log != "" {
		check, aut := login.CheckLogin(*log)
		if check {
			fmt.Println(check, aut)
		} else {
			fmt.Println(check, aut)
		}
	} else if *doc != "" {
		// Doc flag was passed, see the doc folder for more details.
		login.ToolNav()
	} else {
		fmt.Println("No Flags Passed")
	}
}
