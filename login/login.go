package login

import (
	"database/sql"
	"fmt"
	"os"
	"os/exec"

	_ "github.com/lib/pq" //Needed but VS Code says I don't.
	//Then it wonders why everything implodes. Hmmm...
)

const (
	host     = "localhost"
	port     = "5432"
	user     = "postgres"
	password = "postgres"
	dbname   = "postgres"
)

// CheckLogin queries info from the database to check against args passed
func CheckLogin(log string) (bool, bool) {
	var logSuccess, authority, au bool
	logSuccess, authority = false, false
	fmt.Println("checking login")

	if len(os.Args) == 5 {
		fmt.Println("correct number of arguments")
		var job, id, pass, un, pw string
		job = os.Args[2]
		id = os.Args[3]
		pass = os.Args[4]

		dockerstart := exec.Command("docker", "start", "pgcontainer")
		dockerstart.Run()

		db := connect()
		if job == "d" {
			fmt.Println("checking doctor table")
			rows, _ := db.Query("SELECT Doc_Username FROM Doctors")
			for rows.Next() {
				rows.Scan(&un)
				if id == un {
					fmt.Println("found doctor username")
					row := db.QueryRow("SELECT Doc_Password FROM Doctors WHERE Doc_Username = $1", id)
					row.Scan(&pw)
					if pass == pw {
						fmt.Println("found password")
						logSuccess = true
					}
				}
			}
		} else if job == "p" {
			fmt.Println("checking pharmacist table")
			rows, _ := db.Query("SELECT Username FROM Pharmacists")
			for rows.Next() {
				rows.Scan(&un)
				if id == un {
					fmt.Println("found employee id")
					row := db.QueryRow("SELECT Pharm_Password FROM Pharmacists WHERE Username = $1", id)
					row.Scan(&pw)
					if pass == pw {
						fmt.Println("found password")
						logSuccess = true
						aut := db.QueryRow("SELECT Is_Manager FROM Pharmacists WHERE Username = $1", id)
						aut.Scan(&au)
						if au == true {
							authority = true
						}
					}
				}
			}
		}
	}

	return logSuccess, authority
}

func connect() *sql.DB {
	var conn string
	conn = fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	db, err := sql.Open("postgres", conn)
	if err != nil {
		panic(err)
	}
	fmt.Println("successfully connected to database")
	return db
}
