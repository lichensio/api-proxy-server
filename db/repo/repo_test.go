package db

import (
	"github.com/lichensio/api_proxy_server/pkg/db/model"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"testing"
	"time"
)

func setupTestDB() (*gorm.DB, error) {
	// Using SQLite for testing purposes
	db, err := gorm.Open(sqlite.Open("file::memory:?cache=shared"), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	// Migrate the schema
	err = db.AutoMigrate(&model.Employee{}, &model.Schedule{})
	if err != nil {
		return nil, err
	}

	return db, nil
}

func TestLoadEmployees(t *testing.T) {
	db, err := setupTestDB()
	assert.NoError(t, err)

	repo := &repository{db: db}

	employees := []model.Employee{
		{
			Name:      "John Doe",
			StartDate: time.Now(),
		},
		{
			Name:      "Jane Doe",
			StartDate: time.Now(),
		},
	}

	err = repo.LoadEmployees(employees)
	assert.NoError(t, err)

	var dbEmployees []model.Employee
	result := db.Find(&dbEmployees)
	assert.NoError(t, result.Error)
	assert.Equal(t, len(employees), len(dbEmployees))
}

func TestGetEmployees(t *testing.T) {
	db, err := setupTestDB()
	assert.NoError(t, err)

	repo := &repository{db: db}

	expectedEmployees := []model.Employee{
		{
			Name:      "John Doe",
			StartDate: time.Now(),
		},
		{
			Name:      "Jane Doe",
			StartDate: time.Now(),
		},
	}

	// Preload the database
	for _, emp := range expectedEmployees {
		err := db.Create(&emp).Error
		assert.NoError(t, err)
	}

	employees, err := repo.GetEmployees()
	assert.NoError(t, err)
	assert.Equal(t, len(expectedEmployees), len(employees))
}

// Add more tests for other methods following the pattern established above
