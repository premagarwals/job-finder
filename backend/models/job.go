package models

import (
	"gorm.io/gorm"
)

type Job struct {
	gorm.Model
	JobID              int32  `json:"job_id"`
	JobRole            string `json:"job_role"`
	WorkSite           string `json:"work_site"`
	City               string `json:"city"`
	Country            string `json:"country"`
	EmploymentType     string `json:"employment_type"`
	MinSalary          int32  `json:"min_salary"`
	MaxSalary          int32  `json:"max_salary"`
	Profession         string `json:"profession"`
	Discipline         string `json:"discipline"`
	ExperienceRequired string `json:"experience_required"`
	JobDescription     string `json:"job_description"`
}
