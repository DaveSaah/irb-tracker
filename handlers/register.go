package handlers

import (
	"net/http"
	"strconv"
	"time"

	"github.com/boring-school-work/irb-tracker/actions"
	"github.com/boring-school-work/irb-tracker/functions"
	"github.com/boring-school-work/irb-tracker/helpers"
	"github.com/boring-school-work/irb-tracker/model"
	"github.com/labstack/echo/v4"
)

type Department struct {
	Depts []model.Department
}

type StudentData struct {
	Majors       []model.Major
	YearGroupMax int
	YearGroupMin int
}

type Error struct {
	MailError   string
	PasswdError string
}

// RegisterView renders the registration page
func RegisterView(c echo.Context) error {
	depts, err := functions.FetchDepts()
	if err != nil {
		return err
	}

	return c.Render(http.StatusOK, "register", &Department{Depts: depts})
}

// RegisterEmailValidate validates the email used to register
func RegisterEmailValidate(c echo.Context) error {
	email := c.FormValue("email")
	if ok := helpers.IsMailValid(email); !ok {
		return c.Render(http.StatusOK, "mail_error", &Error{MailError: "Invalid email"})
	}

	exists, err := functions.DoesMailExist(email)
	if err != nil {
		return err
	}
	if exists {
		return c.Render(http.StatusOK, "mail_error",
			&Error{MailError: "Email already exists"},
		)
	}

	return nil
}

// RegisterStudentView renders the student block of the registration page
func RegisterStudentView(c echo.Context) error {
	user_type := c.FormValue("type")
	if user_type == "student" {
		majors, err := functions.FetchMajors()
		if err != nil {
			return err
		}

		return c.Render(http.StatusOK, "student_info",
			&StudentData{
				Majors:       majors,
				YearGroupMin: time.Now().Year(),
				YearGroupMax: time.Now().Year() + 4,
			})
	}

	return nil
}

// RegisterPasswdView validates the password used to register
func RegisterPasswdView(c echo.Context) error {
	passwd1 := c.FormValue("passwd1")
	passwd2 := c.FormValue("passwd2")

	if passwd2 != "" && passwd1 != passwd2 {
		return c.Render(http.StatusOK, "passwd_error",
			&Error{PasswdError: "Passwords do not match"},
		)
	}

	return nil
}

// RegisterUser registers a user
func RegisterUser(c echo.Context) error {
	fname := c.FormValue("fname")
	lname := c.FormValue("lname")
	email := c.FormValue("email")
	passwd := c.FormValue("passwd2")
	dept_id, _ := strconv.Atoi(c.FormValue("dept"))
	user_type := c.FormValue("type")

	user := &model.User{
		FName:  fname,
		LName:  lname,
		Email:  email,
		Passwd: passwd,
		DeptID: dept_id,
	}

	if user_type == "student" {
		major, _ := strconv.Atoi(c.FormValue("major"))
		year, _ := strconv.Atoi(c.FormValue("year_group"))
		student_id := c.FormValue("student_id")

		student := &model.Student{
			StudentID: student_id,
			MajorID:   major,
			YearGroup: year,
		}
		user.AddStudent(student)
	}

	err := actions.RegisterUser(user, user_type)
	if err != nil {
		return err
	}

	return nil // TODO: Redirect to login page
}
