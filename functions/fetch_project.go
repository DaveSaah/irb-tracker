package functions

import (
	"log"
	"strings"

	"github.com/boring-school-work/irb-tracker/db"
	"github.com/boring-school-work/irb-tracker/model"
)

func FetchProject(userID, projID int) (model.Project, error) {
	conn, err := db.Init()
	if err != nil {
		log.Printf("Error initializing db:%v", err)
		return model.Project{}, err
	}

	defer conn.Close()

	var project model.Project
	var participantType, recruitmentMtd, purpose, investigators, results_dissemination string

	err = conn.QueryRow(`
    SELECT projects.id, title, department.name, CONCAT(fname, " ", lname),
    brief, start_date, end_date, participants_count, participants_type,
    recuitment_method, proposal, purpose, other_investigators, 
    results_dissemination, status.sname
    FROM projects
    INNER JOIN department
    ON department.id = projects.dept
    INNER JOIN status
    ON status.id = projects.status_id
    INNER JOIN users
    ON projects.supervisor = users.id
    WHERE projects.principal_investigator = ?
    AND projects.id = ?;
    `, userID, projID).Scan(
		&project.ID,
		&project.Title,
		&project.Department,
		&project.Supervisor,
		&project.Brief,
		&project.StartDate,
		&project.EndDate,
		&project.ParticipantCount,
		&participantType,
		&recruitmentMtd,
		&project.Proposal,
		&purpose,
		&investigators,
		&results_dissemination,
		&project.Status,
	)
	if err != nil {
		log.Printf("Error fetching project info: %s", err)
		return model.Project{}, err
	}

	project.ParticipantType = strings.Split(participantType, ",")
	project.RecruitmentMethod = strings.Split(recruitmentMtd, ",")
	project.Purpose = strings.Split(purpose, ",")
	project.Investigators = strings.Split(investigators, ",")
	project.ResultsDissemination = strings.Split(results_dissemination, ",")

	return project, nil
}
