package doc

import (
	"os"
	"os/exec"
	"database/sql"
	"fmt"
)

// These constants are the default names for docker containers
const (
	host     = "localhost"
	port     = "5432"
	user     = "postgres"
	password = "postgres"
	dbname   = "postgres"
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

func actuallyVP() {

	if len(os.Args) == 4 {
		var docun string
		docun = os.Args[3]
		
		//we start up docker
		dockerstart := exec.Command("docker", "start", "pgcontainer")
		dockerstart.Run()
		db := Connect()

		rows, _ := db.Query("SELECT * FROM Prescriptions WHERE Doc_Name = $1", docun)
		rows.Next() {
			prescID := make([]string,)
			var docName string
			var drugName string
			var amount int
			var patientFirst string
			var patientLast string
			var cost float64
			var prescStatus string
			var datePrescribed string
		}
		rows.Scan(&prescID, &docName, &drugName, &amount, &patientFirst, &patientLast)
	}
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
