package functions

import (
	"database/sql"
	"log"

	"github.com/boring-school-work/irb-tracker/db"
)

// DoesMailExist checks if the mail exists in the database
func DoesMailExist(mail string) (bool, error) {
	conn, err := db.Init()
	if err != nil {
		log.Printf("Error initialising db: %v", err)
		return false, err
	}

	defer conn.Close()

	err = conn.QueryRow("SELECT * FROM users WHERE email = ?", mail).Scan()

	switch {
	case err == sql.ErrNoRows:
		return false, nil
	case err != nil:
		log.Printf("Error checking mail: %v", err)
	}

	return true, nil
}
