package functions

import (
	"log"

	"github.com/boring-school-work/irb-tracker/db"
	"github.com/boring-school-work/irb-tracker/model"
)

func FetchStatus() ([]model.Status, error) {
	conn, err := db.Init()
	if err != nil {
		log.Println("Error initialising db: ", err)
		return nil, err
	}

	defer conn.Close()

	rows, err := conn.Query("SELECT * FROM status")
	if err != nil {
		log.Println("Error fetching depts: ", err)
		return nil, err
	}

	var status []model.Status

	for rows.Next() {
		var s model.Status

		err = rows.Scan(&s.ID, &s.Name)
		if err != nil {
			log.Println("Error scanning depts: ", err)
			return nil, err
		}

		status = append(status, s)
	}

	return status, nil
}
