package functions

import (
	"log"

	"github.com/boring-school-work/irb-tracker/db"
	"github.com/boring-school-work/irb-tracker/model"
)

// ProjectSearch returns a slice of projects that match
// a specific pattern
func FindProjects(pattern string) ([]model.Project, error) {
	conn, err := db.Init()
	if err != nil {
		log.Println("Error initialising db: ", err)
		return nil, err
	}

	defer conn.Close()

	rows, err := conn.Query(
		`SELECT projects.id, sname, department.name AS dept_name, date_submitted, 
    CONCAT(fname, " ", lname) as supervisor_name
    FROM projects
    INNER JOIN status
    ON status.id = projects.status_id
    INNER JOIN department
    ON department.id = projects.dept
    INNER JOIN users
    ON users.id = projects.supervisor
    WHERE projects.title LIKE '%?%'
    OR department.name LIKE '%?%'
    OR projects.brief LIKE '%?%'
    `,
		pattern, pattern, pattern,
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
			&project.Status,
			&project.Department,
			&project.DateSubmitted,
			&project.Supervisor,
		)
		if err != nil {
			log.Println("Error scanning projects: ", err)
			return nil, err
		}

		projects = append(projects, project)
	}

	return projects, nil
}
