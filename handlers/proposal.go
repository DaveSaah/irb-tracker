package handlers

import (
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/boring-school-work/irb-tracker/actions"
	"github.com/boring-school-work/irb-tracker/functions"
	"github.com/boring-school-work/irb-tracker/helpers"
	"github.com/boring-school-work/irb-tracker/model"
	"github.com/gorilla/sessions"
	"github.com/labstack/echo/v4"
)

type proposalData struct {
	Today       string
	Session     *sessions.Session
	Supervisors []model.Supervisor
	Departments []model.Department
}

// ProposalView renders the proposal form for the user
func ProposalView(c echo.Context) error {
	sess, isLoggedIn := helpers.CheckSession(c)
	if !isLoggedIn {
		return c.Render(http.StatusBadRequest, "index", nil)
	}

	sups, err := functions.FetchSupervisors()
	if err != nil {
		return err
	}
	depts, err := functions.FetchDepts()
	if err != nil {
		return err
	}

	return c.Render(http.StatusOK, "proposal",
		proposalData{
			Supervisors: sups,
			Session:     sess,
			Departments: depts,
			Today:       time.Now().Format("2006-01-02"),
		},
	)
}

// SubmitProposal validates and submites a proposal form
func SubmitProposal(c echo.Context) error {
	form, err := c.MultipartForm()
	if err != nil {
		log.Printf("Error parsing form: %s\n", err)
		return err
	}

	title := c.FormValue("title")
	brief := c.FormValue("brief")
	dept := c.FormValue("dept")
	sup := c.FormValue("supervisor")
	other_investigators := c.FormValue("investigators")
	investigators := strings.Split(other_investigators, ",")
	start_date := c.FormValue("start_date")
	end_date := c.FormValue("end_date")
	purpose := form.Value["purpose"]
	results_dissemination := form.Value["results_dissemination"]
	participant_count := c.FormValue("participant_count")
	participant_type := form.Value["participant_type"]
	recruitment_method := form.Value["recruitment_method"]
	associated_documents, _ := c.FormFile("associated_documents")

	if c.FormValue("p_has_other") == "true" {
		purpose = append(purpose, c.FormValue("other_purpose"))
	}
	if c.FormValue("d_has_other") == "true" {
		results_dissemination = append(results_dissemination, c.FormValue("other_dissemination"))
	}
	if c.FormValue("r_has_other") == "true" {
		recruitment_method = append(recruitment_method, c.FormValue("other_recruitment"))
	}

	// convert the numerical values to int
	dept_id, _ := strconv.Atoi(dept)
	sup_id, _ := strconv.Atoi(sup)
	part_count, _ := strconv.Atoi(participant_count)

	// get the user's ID
	sess, _ := helpers.CheckSession(c)
	principal_id := sess.Values["id"].(int)

	project := model.Project{
		Title:                title,
		Brief:                brief,
		DeptID:               dept_id,
		SupID:                sup_id,
		Investigators:        investigators,
		StartDate:            start_date,
		EndDate:              end_date,
		Purpose:              purpose,
		ResultsDissemination: results_dissemination,
		ParticipantCount:     part_count,
		ParticipantType:      participant_type,
		RecruitmentMethod:    recruitment_method,
		DocumentPath:         associated_documents.Filename,
		PrincipalID:          principal_id,
	}

	err = actions.SubmitProposal(project)
	if err != nil {
		return c.String(http.StatusInternalServerError, "Project not added!")
	}

	return c.String(http.StatusOK, "Project added successfully!")
}
