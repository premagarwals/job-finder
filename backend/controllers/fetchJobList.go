package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/premagarwals/job-finder/initializers"
	"github.com/premagarwals/job-finder/models"
)

func JobList(c *gin.Context) {
	var jobs []models.Job

	var jobData struct {
		JobRole            string `json:"job_role"`
		WorkSite           string `json:"work_site"`
		City               string `json:"city"`
		Country            string `json:"country"`
		EmploymentType     string `json:"employment_type"`
		Salary             int32  `json:"salary"`
		Profession         string `json:"profession"`
		Discipline         string `json:"discipline"`
		ExperienceRequired string `json:"experience_required"`
		JobDescription     string `json:"job_description"`
		Limit              int32  `json:"limit"`
		Offset             int32  `json:"offset"`
	}

	if err := c.ShouldBindJSON(&jobData); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	var limitInt, offsetInt int

	if jobData.Limit > 0 {
		limitInt = int(jobData.Limit)
	} else {
		limitInt = 10 // Default limit
	}

	if jobData.Offset > 0 {
		offsetInt = int(jobData.Offset)
	} else {
		offsetInt = 0 // Default offset
	}

	query := initializers.DB.Model(&models.Job{})

	if jobData.JobRole != "" {
		query = query.Where("job_role LIKE ?", jobData.JobRole)
	}

	if jobData.Discipline != "" {
		query = query.Where("discipline = ?", jobData.Discipline)
	}

	if jobData.ExperienceRequired != "" {
		query = query.Where("experience_required = ?", jobData.ExperienceRequired)
	}

	if jobData.Salary != 0 {
		query = query.Where("min_salary <= ? AND max_salary >= ?", jobData.Salary, jobData.Salary)
	}

	if jobData.JobRole != "" {
		query = query.Where("job_role = ?", jobData.JobRole)
	}

	if jobData.WorkSite != "" {
		query = query.Where("work_site = ?", jobData.WorkSite)
	}

	if jobData.City != "" {
		query = query.Where("city = ?", jobData.City)
	}

	if jobData.Country != "" {
		query = query.Where("country = ?", jobData.Country)
	}

	if jobData.EmploymentType != "" {
		query = query.Where("employment_type = ?", jobData.EmploymentType)
	}

	if jobData.Profession != "" {
		query = query.Where("profession = ?", jobData.Profession)
	}

	query = query.Offset(offsetInt).Limit(limitInt)

	query.Find(&jobs)

	c.JSON(200, gin.H{
		jobData.JobRole: jobs,
	})
}
