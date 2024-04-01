package functions

import (
	"log"

	"github.com/boring-school-work/irb-tracker/db"
	"github.com/boring-school-work/irb-tracker/model"
)

// FetchMajors fetches all majors from the database
func FetchMajors() ([]model.Major, error) {
	conn, err := db.Init()
	if err != nil {
		log.Println("Error initialising db: ", err)
		return nil, err
	}

	defer conn.Close()

	rows, err := conn.Query("SELECT * FROM majors")
	if err != nil {
		log.Println("Error fetching depts: ", err)
		return nil, err
	}

	var majors []model.Major

	for rows.Next() {
		var major model.Major

		err = rows.Scan(&major.ID, &major.Name)
		if err != nil {
			log.Println("Error scanning depts: ", err)
			return nil, err
		}

		majors = append(majors, major)
	}

	return majors, nil
}
