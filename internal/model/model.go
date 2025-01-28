package model

import (
	"log"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Task struct {
	gorm.Model
	Name string
	Status string
}

func Connect() (*gorm.DB, error) {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		log.Fatal("failed to connect database")
		return nil, err
	}

	// Migrate the schema
	db.AutoMigrate(&Task{})

	return db, nil
}

func Insert(db *gorm.DB, task *Task) {
	db.Create(task)
}

func GetOneByName(db *gorm.DB, name string) *Task {
	var result Task

	db.Where("name = ?", name).First(&result)

	return &result
}

func GetAll(db *gorm.DB) *[]Task {
	var tasks []Task

	db.Find(&tasks)

	return &tasks
}

// UpdateOneByName updates the status of a task to 'Done' by its name
func UpdateOneByName(db *gorm.DB, name string) {
	db.Model(&Task{}).Where("name = ?", name).Update("status", "Done")
}

func CloseDB(db *gorm.DB) {
	sqlDB, err := db.DB()
	if err != nil {
		log.Fatal("failed to get db")
	}
	sqlDB.Close()
}