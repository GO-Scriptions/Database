package login

import (
	"database/sql"
	"fmt"
	"os"
	"os/exec"

	_ "github.com/lib/pq" //Needed but VS Code says I don't.
	//Then it wonders why everything implodes. Hmmm...
)

// These constants are the default names for docker containers
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

	// If for some reason more or less arguments are passed this will simply return a failed login
	if len(os.Args) == 5 {
		var job, id, pass, un, pw string
		job = os.Args[2]
		id = os.Args[3]
		pass = os.Args[4]

		// Connect to the database
		dockerstart := exec.Command("docker", "start", "pgcontainer")
		dockerstart.Run()
		db := Connect()

		if job == "d" {
			// For doctor login "d" should be the first argument passed after the flag
			rows, _ := db.Query("SELECT Doc_Username FROM Doctors")
			for rows.Next() {
				rows.Scan(&un)
				if id == un {
					// If a username, the second argument afther the flag, is found in the database
					row := db.QueryRow("SELECT Doc_Password FROM Doctors WHERE Doc_Username = $1", id)
					row.Scan(&pw)
					if pass == pw {
						// If the password in the table matches the password passed in the third argument after the flag
						logSuccess = true
					}
				}
			}
		} else if job == "p" {
			// For pharmacists login "p" should be passed instead of "d"
			// This is all now repeat code, probably easier to put it in a function and save file space. Oh well!
			rows, _ := db.Query("SELECT Username FROM Pharmacists")
			for rows.Next() {
				rows.Scan(&un)
				if id == un {
					row := db.QueryRow("SELECT Pharm_Password FROM Pharmacists WHERE Username = $1", id)
					row.Scan(&pw)
					if pass == pw {
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
	// This is what the client server will be expecting when running the login function
	return logSuccess, authority
}

// Connect sets up the connection with the database in the docker container
func Connect() *sql.DB {
	var conn string
	conn = fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	db, err := sql.Open("postgres", conn)
	if err != nil {
		panic(err)
	}
	return db
}
