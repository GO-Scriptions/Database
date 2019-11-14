package main

import (
	"flag"
	"fmt"

	"github.com/GO-Scriptions/Database/login"
)

func main() {
	log := flag.String("log", "", "Login type. Use the flags d or p login.")
	flag.Parse()

	if *log != "" {
		check, aut := login.CheckLogin(*log)
		if check == true {
			fmt.Println("Pass")
			fmt.Println(check, aut)
		} else {
			fmt.Println("Fail")
			fmt.Println(check, aut)
		}
	} else {
		fmt.Println("No Flags Passed")
	}
}
