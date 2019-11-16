package login

import (
	"fmt"
	"math/rand"
	"os"
	"os/exec"
	"strconv"
	"time"
)

// ToolNav function for navigating the doctor tools
func ToolNav() {
	var doctool string
	doctool = os.Args[2]

	if doctool == "wp" {
		WriteP()
	} else if doctool == "vp" {
		ViewP()
	} else {
		fmt.Println("Invalid argument tag")
	}
}

// WriteP for writing prescriptions
func WriteP() {
	var status string
	status = "fail"

	if len(os.Args) == 8 {

		doct := os.Args[3]
		frst := os.Args[4]
		last := os.Args[5]
		quan := os.Args[6]
		drug := os.Args[7]

		currentTime := time.Now()
		id := NewID()
		date := currentTime.Format("2006-01-02")
		prce := NewPrice()

		dockerstart := exec.Command("docker", "start", "pgcontainer")
		dockerstart.Run()
		db := Connect()
		db.Exec("INSERT INTO Prescriptions VALUES ($1, $2, $3, $4, $5, $6, $7, 'unfilled', $8);",
			id, doct, drug, quan, frst, last, prce, date)
		status = "success"
	}

	fmt.Println(status)
}

// NewPrice generates a random price of a drug
func NewPrice() float64 {
	var inCents, inDollars float64
	inCents = float64(UseSeed(100, 300000))
	inDollars = inCents / 100
	return inDollars
}

// NewID generates a new six-digit ID number
func NewID() string {
	var newNum, sDig string
	var dig int
	newNum = ""

	for i := 0; i <= 5; i++ {
		dig = UseSeed(0, 9)
		sDig = strconv.Itoa(dig)
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
func ViewP() {
	var user, prid, docn, drug, patf, patl, stat, date string
	var cost float64
	var quan int

	if len(os.Args) == 4 {
		user = os.Args[3]

		//we start up docker
		dockerstart := exec.Command("docker", "start", "pgcontainer")
		dockerstart.Run()
		db := Connect()

		rows, _ := db.Query("SELECT * FROM Prescriptions WHERE Doc_Name = $1", user)
		for rows.Next() {
			rows.Scan(&prid, &docn, &drug, &quan, &patf, &patl, &cost, &stat, &date)
			fmt.Println(prid, docn, drug, quan, patf, patl, cost, stat, date)
		}
	} else {
		fmt.Println("Invalid arguments")
	}
}
