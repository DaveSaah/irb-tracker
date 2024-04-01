// Package model defines the data model for the application.
package model

type Department struct {
	Name string
	ID   int
}

type Major struct {
	Name string
	ID   int
}

type Student struct {
	StudentID string
	MajorID   int
	YearGroup int
}

type User struct {
	StudentUser *Student
	FName       string
	LName       string
	Email       string
	Passwd      string
	Type        string
	DeptID      int
}

// AddStudent adds a student to the user
func (u *User) AddStudent(s *Student) {
	u.StudentUser = s
}

// AddPasswdHash replaces the user's password with the provided hash
func (u *User) AddPasswdHash(hash string) {
	u.Passwd = hash
}

// ClearPasswd removes the password from the user
func (u *User) ClearPasswd() {
	u.Passwd = ""
}
