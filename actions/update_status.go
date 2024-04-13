package actions

import (
	"log"

	"github.com/boring-school-work/irb-tracker/db"
)

func UpdateStatus(projID, status int) error {
	conn, err := db.Init()
	if err != nil {
		log.Printf("Error initializing database: %v", err)
		return err
	}

	defer conn.Close()

	_, err = conn.Exec(
		`UPDATE projects SET status_id=? WHERE id=?`,
		status, projID,
	)
	if err != nil {
		log.Printf("Error updating project status: %v", err)
		return err
	}

	return nil
}
