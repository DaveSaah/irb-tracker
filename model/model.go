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
	DeptID      int
}

