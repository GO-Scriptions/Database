package main

import (
	"flag"
	"fmt"

	"github.com/GO-Scriptions/Database/login"
)

func main() {
	log := flag.String("log", "", "Login type. Use the flags -d, -p, or -e to login.")
	flag.Parse()

	if log != nil {
		check := login.Check(*log)
		if check == true {
			fmt.Println("Pass")
		} else {
			fmt.Println("Fail")
		}
	} else {
		fmt.Println("No Flags Passed")
	}
}
