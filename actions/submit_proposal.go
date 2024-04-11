package actions

import (
	"log"
	"strings"

	"github.com/boring-school-work/irb-tracker/db"
	"github.com/boring-school-work/irb-tracker/model"
)

// SubmitProposal submits a proposal form
func SubmitProposal(p model.Project) error {
	conn, err := db.Init()
	if err != nil {
		log.Printf("Error initializing database: %s\n", err)
		return err
	}

	defer conn.Close()

	_, err = conn.Exec(
		`INSERT INTO projects (
    brief, title, associated_documents, principal_investigator,
    dept, supervisor, start_date, end_date, other_investigators,
    purpose, results_dissemination, participants_count,
    participants_type, recruitment_method
    ) 
    VALUES 
    (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`,
		p.Brief, p.Title, p.DocumentPath, p.PrincipalID,
		p.DeptID, p.SupID, p.StartDate, p.EndDate, strings.Join(p.Investigators, ","),
		strings.Join(p.Purpose, ","), strings.Join(p.ResultsDissemination, ","), p.ParticipantCount,
		strings.Join(p.ParticipantType, ","), strings.Join(p.RecruitmentMethod, ","),
	)
	if err != nil {
		log.Printf("Error inserting project: %s\n", err)
		return err
	}

	return nil
}
