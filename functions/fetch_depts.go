package functions

import (
	"log"

	"github.com/boring-school-work/irb-tracker/db"
	"github.com/boring-school-work/irb-tracker/model"
)

// FetchDepts fetches all departments from the database
func FetchDepts() ([]model.Department, error) {
	conn, err := db.Init()
	if err != nil {
		log.Println("Error initialising db: ", err)
		return nil, err
	}

	rows, err := conn.Query("SELECT * FROM department")
	if err != nil {
		log.Println("Error fetching depts: ", err)
		return nil, err
	}

	var depts []model.Department

	for rows.Next() {
		var dept model.Department

		err = rows.Scan(&dept.ID, &dept.Name)
		if err != nil {
			log.Println("Error scanning depts: ", err)
			return nil, err
		}

		depts = append(depts, dept)
	}

	return depts, nil
}
