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

type registerData struct {
	MailError   string
	PasswdError string
	Depts       []model.Department
}

type studentData struct {
	Majors       []model.Major
	YearGroupMax int
	YearGroupMin int
}

var depts, _ = functions.FetchDepts()

// RegisterView renders the registration page
func RegisterView(c echo.Context) error {
	_, isLoggedIn := helpers.CheckSession(c)
	if isLoggedIn {
		return c.Redirect(http.StatusSeeOther, "/dashboard")
	}

	return c.Render(http.StatusOK, "register", registerData{Depts: depts})
}

// RegisterEmailValidate validates the email used to register
func RegisterEmailValidate(c echo.Context) error {
	email := c.FormValue("email")
	if ok := helpers.IsMailValid(email); !ok {
		return c.Render(http.StatusBadRequest, "mail_error", registerData{MailError: "Invalid email"})
	}

	exists, err := functions.DoesMailExist(email)
	if err != nil {
		return err
	}
	if exists {
		return c.Render(http.StatusBadRequest, "mail_error",
			registerData{MailError: "Email already exists"},
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
			studentData{
				Majors:       majors,
				YearGroupMin: time.Now().Year(),
				YearGroupMax: time.Now().Year() + 4,
			})
	}

	return nil
}

// RegisterPasswdView validates the password used to register
func RegisterPasswdValidate(c echo.Context) error {
	passwd1 := c.FormValue("passwd1")
	passwd2 := c.FormValue("passwd2")

	if passwd2 != "" && passwd1 != passwd2 {
		return c.Render(http.StatusBadRequest, "passwd_error",
			registerData{PasswdError: "Passwords do not match"},
		)
	}

	return nil
}

// RegisterUser registers a user
func RegisterUser(c echo.Context) error {
	email := c.FormValue("email")

	// check if email exists
	exists, err := functions.DoesMailExist(email)
	if err != nil {
		return err
	}
	if exists {
		return c.Render(http.StatusBadRequest, "register",
			registerData{
				MailError: "Email already exists, try again",
				Depts:     depts,
			},
		)
	}

	fname := c.FormValue("fname")
	lname := c.FormValue("lname")
	passwd := c.FormValue("passwd2")
	user_type := c.FormValue("type")
	dept_id, _ := strconv.Atoi(c.FormValue("dept"))

	user := model.User{
		FName:  fname,
		LName:  lname,
		Email:  email,
		Passwd: passwd,
		DeptID: dept_id,
		Type:   user_type,
	}

	if user_type == "student" {
		major, _ := strconv.Atoi(c.FormValue("major"))
		year, _ := strconv.Atoi(c.FormValue("year_group"))
		student_id := c.FormValue("student_id")

		student := model.Student{
			StudentID: student_id,
			MajorID:   major,
			YearGroup: year,
		}
		user.AddStudent(student)
	}

	err = actions.RegisterUser(&user)
	if err != nil {
		return err
	}

	return c.Render(http.StatusOK, "login", nil)
}
