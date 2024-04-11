package actions

import (
	"database/sql"
	"errors"
	"log"

	"github.com/boring-school-work/irb-tracker/db"
	"github.com/boring-school-work/irb-tracker/model"
	"golang.org/x/crypto/bcrypt"
)

// LoginUser logs in a user
func LoginUser(email, passwd string) (model.User, error) {
	conn, err := db.Init()
	if err != nil {
		log.Printf("Cannot create the db connection: %s\n", err)
		return model.User{}, err
	}

	defer conn.Close()

	u := model.User{}

	err = conn.QueryRow(
		`SELECT 
    id, fname, lname, passwd, dept, user_type 
    FROM users 
    WHERE email=?`,
		email,
	).Scan(
		&u.ID,
		&u.FName,
		&u.LName,
		&u.Passwd,
		&u.DeptID,
		&u.Type,
	)

	// check if the email exists
	switch {
	case err == sql.ErrNoRows:
		log.Printf("No user with email: %s", email)
		return model.User{}, errors.New("no email found")
	case err != nil:
		log.Printf("Error querying the database: %s", err)
		return model.User{}, err
	}

	// check if the password matches
	err = bcrypt.CompareHashAndPassword([]byte(u.Passwd), []byte(passwd))
	if err != nil {
		log.Printf("Password does not match: %s", err)
		return model.User{}, errors.New("wrong password")
	}

	u.ClearPasswd() // remove the password from the user struct

	if u.Type == "student" {
		// create a memory location to store the student data
		s := model.Student{}
		u.AddStudent(s)

		_ = conn.QueryRow(
			`SELECT major, year_group, student_id FROM student WHERE id = ?`,
			u.ID,
		).Scan(
			&u.Student.MajorID,
			&u.Student.YearGroup,
			&u.Student.StudentID,
		)

	}

	log.Printf("User logged in: %s", email)
	return u, nil
}
