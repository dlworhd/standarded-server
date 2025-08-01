package handler

import (
	"database/sql"
	"net/http"

	"github.com/dlworhd/standarded/model"
	"github.com/gin-gonic/gin"
)

type ProfileHandler struct {
	Repository *sql.DB
}

func (h *ProfileHandler) ReadHandler(ctx *gin.Context) {
	profile_id := ctx.Param("id")
	rows := h.Repository.QueryRow(`SELECT ID,
									NICKNAME,
									EMAIL,
									EMAIL_PUBLIC,
									JOB,
									JOB_PUBLIC,
									EDUCATION,
									EDUCATION_PUBLIC,
									NATIONALITY,
									NATIONALITY_PUBLIC,
									LOCATION,
									LOCATION_PUBLIC,
									COMPANY,
									LINKS
								FROM PROFILES
								WHERE ID = $1`, profile_id)

	profile := model.Profile{}
	rows.Scan(
		&profile.ID,
		&profile.NickName,
		&profile.Email,
		&profile.EmailPublic,
		&profile.Job,
		&profile.JobPublic,
		&profile.Education,
		&profile.EducationPublic,
		&profile.Nationality,
		&profile.NationalityPublic,
		&profile.Location,
		&profile.LocationPublic,
		&profile.Company,
		&profile.Links,
	)

	ctx.JSON(http.StatusOK, gin.H{
		"profile_id": profile.ID,
	})

}

func (h *ProfileHandler) ReadAllHandler(ctx *gin.Context) {
	rows, err := h.Repository.Query(`SELECT ID,
											NICKNAME,
											EMAIL,
											EMAIL_PUBLIC,
											JOB,
											JOB_PUBLIC,
											EDUCATION,
											EDUCATION_PUBLIC,
											NATIONALITY,
											NATIONALITY_PUBLIC,
											LOCATION,
											LOCATION_PUBLIC,
											COMPANY,
											LINKS
										FROM PROFILES`)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	var profiles []model.Profile

	for rows.Next() {
		profile := model.Profile{}
		rows.Scan(
			&profile.ID,
			&profile.NickName,
			&profile.Email,
			&profile.EmailPublic,
			&profile.Job,
			&profile.JobPublic,
			&profile.Education,
			&profile.EducationPublic,
			&profile.Nationality,
			&profile.NationalityPublic,
			&profile.Location,
			&profile.LocationPublic,
			&profile.Company,
			&profile.Links,
		)

		profiles = append(profiles, profile)
	}

	ctx.JSON(http.StatusOK, profiles)
}
