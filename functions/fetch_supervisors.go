package functions

import (
	"log"

	"github.com/boring-school-work/irb-tracker/db"
	"github.com/boring-school-work/irb-tracker/model"
)

// FetchSupervisors fetches faculty for students to choose
// as a supervisor
func FetchSupervisors() ([]model.Supervisor, error) {
	conn, err := db.Init()
	if err != nil {
		log.Println("Error initialising db: ", err)
		return nil, err
	}

	defer conn.Close()

	rows, err := conn.Query(
		`SELECT users.id, CONCAT(fname, ' ', lname) AS name 
    FROM users, faculty
    WHERE users.id = faculty.id`,
	)
	if err != nil {
		log.Println("Error fetching supervisors: ", err)
		return nil, err
	}

	var sups []model.Supervisor

	for rows.Next() {
		var s model.Supervisor

		err = rows.Scan(&s.ID, &s.Name)
		if err != nil {
			log.Println("Error scanning depts: ", err)
			return nil, err
		}

		sups = append(sups, s)
	}

	return sups, nil
}
