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
	FName       string
	LName       string
	Email       string
	Passwd      string
	Type        string
	StudentUser Student
	DeptID      int
	ID          int
}

// AddStudent adds a student to the user
func (u *User) AddStudent(s Student) {
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

type Project struct {
	Title         string
	Status        string
	Department    string
	DateSubmitted string
	Supervisor    string
	ID            int
}
