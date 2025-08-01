package model

type Profile struct {
	ID                string `json:"user_id"`
	Avatar            string `json:"avatar"`
	NickName          string `json:"nickname"`
	Email             string `json:"email"`
	EmailPublic       bool   `json:"email_public"`
	Job               string `json:"job"`
	JobPublic         bool   `json:"job_public"`
	Education         string `json:"education"`
	EducationPublic   string `json:"education_public"`
	Nationality       string `json:"nationality"`
	NationalityPublic bool   `json:"nationality_public"`
	Location          string `json:"location"`
	LocationPublic    bool   `json:"location_public"`
	Company           string `json:"company"`
	Links             []Link `json:"links"`
}

type Link struct {
	ID   string `json:"link_id"`
	Name string `json:"link_name"`
	URL  string `json:"link_url"`
}
