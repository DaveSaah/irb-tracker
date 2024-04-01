package actions

import (
	"log"

	"github.com/boring-school-work/irb-tracker/db"
	"github.com/boring-school-work/irb-tracker/helpers"
	"github.com/boring-school-work/irb-tracker/model"
)

// RegisterUser registers a new user
func RegisterUser(u *model.User, user_type string) error {
	conn, err := db.Init()
	if err != nil {
		log.Printf("Error initializing db: %s", err)
		return err
	}

	defer conn.Close()

	hash, err := helpers.GetHashPassword(u.Passwd)
	if err != nil {
		log.Printf("Error hashing password: %s", err)
		return err
	}

	// replace the user's password with the provided hash
	u.AddPasswdHash(hash)

	// begin transaction
	tx, err := conn.Begin()
	if err != nil {
		log.Printf("Cannot start db transaction: %s\n", err)
		return err
	}

	defer func() {
		_ = tx.Rollback() // aborts transaction
	}()

	// execute first transaction
	res, err := tx.Exec(
		`INSERT INTO users(fname, lname, email, passwd, dept) 
    VALUES(?, ?, ?, ?, ?)`,
		u.FName, u.LName, u.Email, u.Passwd, u.DeptID,
	)
	if err != nil {
		log.Printf("Cannot insert user information: %s\n", err)
		return err
	}

	// get last inserted id
	user_id, _ := res.LastInsertId()

	if user_type == "student" {
		// add student info
		_, err = tx.Exec(
			`INSERT INTO student(id, student_id, major, year_group)
      VALUES(?, ?, ?, ?)`,
			user_id,
			u.StudentUser.StudentID,
			u.StudentUser.MajorID,
			u.StudentUser.YearGroup,
		)
		if err != nil {
			log.Printf("Cannot insert student information: %s\n", err)
			return err
		}
	} else {
		// add faculty info
		_, err = tx.Exec("INSERT INTO faculty VALUES(?)", user_id)
		if err != nil {
			log.Printf("Cannot insert faculty information: %s\n", err)
			return err
		}
	}

	// commit transaction
	if err = tx.Commit(); err != nil {
		log.Printf("Cannot register user: %s\n", err)
		return err
	}

	log.Printf("User registered successfully. ID: %d\n", user_id)
	return nil
}
