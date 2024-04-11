package functions

import (
	"log"

	"github.com/boring-school-work/irb-tracker/db"
	"github.com/boring-school-work/irb-tracker/model"
)

// FetchProjects fetches all the projects for a user from the database
func FetchProjects(userID int) ([]model.Project, error) {
	conn, err := db.Init()
	if err != nil {
		log.Printf("Error initialising db: %v", err)
		return nil, err
	}

	defer conn.Close()

	rows, err := conn.Query(`
    SELECT projects.id, title, brief, sname, CONCAT(fname, " ", lname)
    FROM projects
    INNER JOIN status
    ON status.id = projects.status_id
    INNER JOIN users
    ON users.id = projects.supervisor
    WHERE projects.principal_investigator = ?
    ORDER BY date_submitted;
    `, userID)
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
			&project.Supervisor,
		)
		if err != nil {
			log.Println("Error scanning depts: ", err)
			return nil, err
		}

		projects = append(projects, project)
	}

	return projects, nil
}
