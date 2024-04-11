package functions

import (
	"log"

	"github.com/boring-school-work/irb-tracker/db"
)

const (
	// StatusSubmitted is the status id for a submitted project
	StatusSubmitted = 0
	// StatusPending is the status id for a project pending approval
	StatusPending = 1
	// StatusRejrcted is the status id for a rejected project
	StatusRejected = 4
	// StatusApproved is the status id for an approved project
	StatusApproved = 5
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

	if status == StatusSubmitted {
		err = conn.QueryRow(`
    SELECT COUNT(*) FROM projects
    INNER JOIN users
    ON users.id = ?`, userID).Scan(&count)
	} else {
		err = conn.QueryRow(`
    SELECT COUNT(*) FROM projects
    INNER JOIN users
    ON users.id = ?
    WHERE status_id = ?`,
			userID, status,
		).Scan(&count)
	}
	if err != nil {
		log.Printf("Error getting project count: %v", err)
		return 0, err
	}

	return count, nil
}
