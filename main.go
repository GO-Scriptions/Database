package main

import (
	"flag"
	"fmt"

	"github.com/GO-Scriptions/Database/login"
)

func main() {
	log := flag.String("log", "", "Login type. Use the flags d or p login.")
	flag.Parse()

	// Login flag was passed, see the login folder for more details.
	if *log != "" {
		check, aut := login.CheckLogin(*log)
		if check {
			fmt.Println(check, aut)
		} else {
			fmt.Println(check, aut)
		}
	} else {
		fmt.Println("No Flags Passed")
	}
}
