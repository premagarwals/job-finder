package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"

	"github.com/premagarwals/job-finder/initializers"
	"github.com/premagarwals/job-finder/models"
	"gorm.io/gorm"
)

func init() {
	initializers.LoadEnv()
	initializers.ConnectToDB()
}

func clearDatabase(db *gorm.DB) {
	if err := db.Exec("DELETE FROM jobs").Error; err != nil {
		log.Fatalf("Could not clear jobs table: %v", err)
	}
}

func populateDatabase(db *gorm.DB, filename string) {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatalf("Could not open JSON file: %v", err)
	}
	defer file.Close()

	bytes, err := io.ReadAll(file)
	if err != nil {
		log.Fatalf("Could not read JSON file: %v", err)
	}

	var jobs []models.Job
	if err := json.Unmarshal(bytes, &jobs); err != nil {
		log.Fatalf("Could not unmarshal JSON data: %v", err)
	}

	for _, job := range jobs {
		if err := db.Create(&job).Error; err != nil {
			log.Fatalf("Could not insert job data: %v", err)
		}
	}
}

func main() {
	db := initializers.DB

	clearDatabase(db)

	if err := db.AutoMigrate(&models.Job{}); err != nil {
		log.Fatalf("Could not migrate database: %v", err)
	}

	populateDatabase(db, "job.json")

	fmt.Println("Database migration and seeding completed successfully!")
}
