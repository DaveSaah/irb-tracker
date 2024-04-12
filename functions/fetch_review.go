package functions

import (
	"log"

	"github.com/boring-school-work/irb-tracker/db"
	"github.com/boring-school-work/irb-tracker/model"
)

func FindProjectForReview(userID int, pattern string) ([]model.Project, error) {
	conn, err := db.Init()
	if err != nil {
		log.Println("Error initialising db: ", err)
		return nil, err
	}

	defer conn.Close()

	pattern = "%" + pattern + "%"

	rows, err := conn.Query(
		`SELECT projects.id, title, brief, sname, date_submitted
    FROM projects
    INNER JOIN status
    ON status.id = projects.status_id
    WHERE supervisor = ?
    AND (
    projects.title LIKE ?
    OR projects.brief LIKE ?)
    `,
		userID,
		pattern, pattern,
	)
	if err != nil {
		log.Println("Error fetching projects: ", err)
		return nil, err
	}

	var projects []model.Project

	for rows.Next() {
		var project model.Project

		err = rows.Scan(
			&project.ID,
			&project.Title,
			&project.Brief,
			&project.Status,
			&project.DateSubmitted,
		)
		if err != nil {
			log.Println("Error scanning projects: ", err)
			return nil, err
		}

		projects = append(projects, project)
	}

	return projects, nil
}

func FetchAllProjectsForReview(userID int) ([]model.Project, error) {
	conn, err := db.Init()
	if err != nil {
		log.Println("Error initialising db: ", err)
		return nil, err
	}

	defer conn.Close()

	rows, err := conn.Query(
		`SELECT projects.id, title, brief, sname, date_submitted
    FROM projects
    INNER JOIN status
    ON status.id = projects.status_id
    WHERE supervisor = ?
    `, userID,
	)
	if err != nil {
		log.Println("Error fetching projects: ", err)
		return nil, err
	}

	var projects []model.Project

	for rows.Next() {
		var project model.Project

		err = rows.Scan(
			&project.ID,
			&project.Title,
			&project.Brief,
			&project.Status,
			&project.DateSubmitted,
		)
		if err != nil {
			log.Println("Error scanning projects: ", err)
			return nil, err
		}

		projects = append(projects, project)
	}

	return projects, nil
}
