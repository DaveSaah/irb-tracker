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
	Major     string
	MajorID   int
	YearGroup int
}

type Supervisor struct {
	Name string
	ID   int
}

type User struct {
	FName      string
	LName      string
	Email      string
	Passwd     string
	Type       string
	Department string
	Student    Student
	DeptID     int
	ID         int
}

// AddStudent adds a student to the user
func (u *User) AddStudent(s Student) {
	u.Student = s
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
	EndDate               string
	Department            string
	DateSubmitted         string
	Supervisor            string
	Brief                 string
	Status                string
	Title                 string
	Proposal              string
	StartDate             string
	DocumentPath          string
	PrincipalInvestigator string
	Investigators         []string
	ResultsDissemination  []string
	ParticipantType       []string
	RecruitmentMethod     []string
	Purpose               []string
	PrincipalID           int
	ParticipantCount      int
	DeptID                int
	SupID                 int
	ID                    int
}
