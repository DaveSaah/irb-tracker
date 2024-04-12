package functions

import (
	"database/sql"
	"log"

	"github.com/boring-school-work/irb-tracker/db"
	"github.com/boring-school-work/irb-tracker/types"
)

// GetProjectCount returns the number of projects in the database for a user
//
// Arguments:
//
//	userID: the user id
//	status: the status of the project
//
// Returns:
//
//	int: the number of projects
//	error: the error if any
func GetProjectCount(userID, status int) (int, error) {
	conn, err := db.Init()
	if err != nil {
		log.Printf("Error initialising db: %v", err)
		return 0, err
	}

	defer conn.Close()

	var count int

	if status == types.StatusSubmitted {
		err = conn.QueryRow(`
    SELECT COUNT(*) FROM projects
    WHERE principal_investigator = ?`, userID).Scan(&count)
	} else {
		err = conn.QueryRow(`
    SELECT COUNT(*) FROM projects
    WHERE principal_investigator = ?
    AND status_id = ?`,
			userID, status,
		).Scan(&count)
	}
	if err != nil && err != sql.ErrNoRows {
		log.Printf("Error getting project count: %v", err)
		return 0, err
	}

	return count, nil
}

func GetProjectReviewCount(userID, status int) (int, error) {
	conn, err := db.Init()
	if err != nil {
		log.Printf("Error initialising db: %v", err)
		return 0, err
	}

	defer conn.Close()

	var count int

	if status == types.StatusSubmitted {
		err = conn.QueryRow(`
    SELECT COUNT(*) FROM projects
    WHERE supervisor = ?`, userID).Scan(&count)
	} else {
		err = conn.QueryRow(`
    SELECT COUNT(*) FROM projects
    WHERE supervisor = ?
    AND status_id = ?`,
			userID, status,
		).Scan(&count)
	}
	if err != nil && err != sql.ErrNoRows {
		log.Printf("Error getting project count: %v", err)
		return 0, err
	}

	return count, nil
}
