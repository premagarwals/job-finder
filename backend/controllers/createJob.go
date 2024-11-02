package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/premagarwals/job-finder/initializers"
	"github.com/premagarwals/job-finder/models"
)

func JobCreate(c *gin.Context) {
	var jobData struct {
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
		ExperienceRequired int8   `json:"experience_required"`
		JobDescription     string `json:"job_description"`
	}

	if err := c.ShouldBindJSON(&jobData); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	job := models.Job{
		JobID:              jobData.JobID,
		JobRole:            jobData.JobRole,
		WorkSite:           jobData.WorkSite,
		City:               jobData.City,
		Country:            jobData.Country,
		EmploymentType:     jobData.EmploymentType,
		MinSalary:          jobData.MinSalary,
		MaxSalary:          jobData.MaxSalary,
		Profession:         jobData.Profession,
		Discipline:         jobData.Discipline,
		ExperienceRequired: jobData.ExperienceRequired,
		JobDescription:     jobData.JobDescription,
	}

	result := initializers.DB.Create(&job)

	if result.Error != nil {
		c.JSON(400, gin.H{"error": result.Error.Error()})
		return
	}

	c.JSON(200, gin.H{
		"job": job,
	})
}
