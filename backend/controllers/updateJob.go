package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/premagarwals/job-finder/initializers"
	"github.com/premagarwals/job-finder/models"
)

func JobUpdate(c *gin.Context) {
	var job models.Job
	id := c.Param("id")
	initializers.DB.First(&job, id)

	var jobData struct {
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

	result := initializers.DB.Model(&job).Updates(models.Job{
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
	})

	if result.Error != nil {
		c.JSON(400, gin.H{"error": result.Error.Error()})
		return
	}

	c.JSON(200, gin.H{
		"job": job,
	})
}
