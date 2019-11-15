package login

import (
	"fmt"
	"math/rand"
	"os"
	"time"
)

// ToolNav function for navigating the doctor tools
func ToolNav() string {
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
	var status string
	status = "fail"

	if len(os.Args) == 8 {

		doct := os.Args[3]
		frst := os.Args[4]
		last := os.Args[5]
		quan := os.Args[6]
		drug := os.Args[7]

		id := NewID()
		currentTime := time.Now()
		date := currentTime.Format("2006-01-02")
		prce := NewPrice()

		db := Connect()
		db.Exec("INSERT INTO Prescriptions VALUES ($1, $2, $3, $4, $5, $6, $7, 'unfilled', $8);",
			id, doct, drug, quan, frst, last, prce, date)

	} else {
		fmt.Println("Not enough arguments")
	}

	return status
}

// NewPrice generates a random price of a drug
func NewPrice() float64 {
	var inCents, inDollars float64
	inCents = float64(UseSeed(100, 300000))
	inDollars = inCents / 100
	return inDollars
}

// NewID generates a new six-diget ID number
func NewID() string {
	var newNum, sDig string
	var dig int
	newNum = ""

	for i := 0; i >= 5; i++ {
		dig = UseSeed(0, 9)
		sDig = string(dig)
		newNum = newNum + sDig
	}

	return newNum
}

// UseSeed generates a random integer using the current time as a seed
func UseSeed(min int, max int) int {
	rand.Seed(time.Now().UnixNano())
	random := min + rand.Intn(max-min+1)
	return random
}

// ViewP func for Viewing prescriptions
func ViewP() string {
	var r string
	r = "View Prescriptons"
	return r
}
