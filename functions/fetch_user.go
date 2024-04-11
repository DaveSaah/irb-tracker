package functions

import (
	"log"

	"github.com/boring-school-work/irb-tracker/db"
	"github.com/boring-school-work/irb-tracker/model"
)

// FetchUser returns user email, student_id if any, and department
func FetchUser(userID int) (model.User, error) {
	conn, err := db.Init()
	if err != nil {
		log.Printf("Error initialising database: %v", err)
		return model.User{}, err
	}

	defer conn.Close()

	var user model.User

	err = conn.QueryRow(`
    SELECT email, user_type, department.name FROM users
    INNER JOIN department
    ON users.dept = department.id
    WHERE users.id =?`, userID,
	).Scan(&user.Email, &user.Type, &user.Department)
	if err != nil {
		log.Printf("Cannot fetch user information: %s\n", err)
		return model.User{}, err
	}

	if user.Type == "student" {
		s := model.Student{}

		err := conn.QueryRow(`
    SELECT majors.name, year_group, student_id
    FROM student
    INNER JOIN majors
    ON majors.id = student.major
    WHERE student.id = ?`, userID,
		).Scan(&s.Major, &s.YearGroup, &s.StudentID)
		if err != nil {
			log.Printf("Cannot fetch student information: %s\n", err)
			return model.User{}, err
		}

		user.AddStudent(s)
	}

	return user, nil
}
