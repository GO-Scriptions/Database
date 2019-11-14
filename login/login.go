package login

import (
	"database/sql"
	"fmt"
	"os"
	"os/exec"
)

const (
	host     = "localhost"
	port     = "5432"
	user     = "postgres"
	password = "postgres"
	dbname   = "postgres"
)

// CheckLogin quiries info from the database to check against args passed
func CheckLogin(log string) (bool, bool) {
	var logSuccess, authority, au bool
	logSuccess, authority = false, false
	fmt.Println("checking login")

	if len(os.Args) == 5 {
		fmt.Println("correct number of arguments")
		var job, name, pass, un, pw string
		job = os.Args[3]
		name = os.Args[4]
		pass = os.Args[5]

		dockerstart := exec.Command("docker", "start", "pgcontainer")
		dockerstart.Run()

		db := connect()
		if job == "d" {
			fmt.Println("checking doctor table")
			rows, _ := db.Query("SELECT Doctor_ID FROM Doctors")
			for rows.Next() {
				rows.Scan(&un)
				if name == un {
					fmt.Println("found username")
					row := db.QueryRow("SELECT Doc_Password FROM Doctors WHERE Doctor_ID = $1", name)
					row.Scan(&pw)
					if pass == pw {
						fmt.Println("found password")
						logSuccess = true
					}
				}
			}
		} else if job == "p" {
			rows, _ := db.Query("SELECT Employee_ID FROM Pharmacists")
			for rows.Next() {
				rows.Scan(&un)
				if name == un {
					row := db.QueryRow("SELECT Pharm_Password FROM Pharmacists WHERE Employee_ID = $1", name)
					row.Scan(&pw)
					if pass == pw {
						logSuccess = true
						aut := db.QueryRow("SELECT Is_Manager FROM Pharmacists WHERE Employee_ID = $1", name)
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
