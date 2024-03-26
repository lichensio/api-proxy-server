package db

import (
	"github.com/lichensio/api_proxy_server/pkg/db/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Repository interface {
	LoadEmployees(employees []model.Employee) error
	UpdateEmployee(employee model.Employee) error
	UpdateSchedule(schedule model.Schedule) error
	GetSchedule(employeeID int, weekType string) ([]model.Schedule, error)
	GetEmployees() ([]model.Employee, error)
	// Define more methods for analytics or other operations as needed
}

type repository struct {
	db *gorm.DB
}

func NewRepository(dsn string) (Repository, error) {
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	// Migrate the schema
	err = db.AutoMigrate(&model.Employee{}, &model.Schedule{})
	if err != nil {
		return nil, err
	}

	return &repository{db: db}, nil
}

func (r *repository) LoadEmployees(employees []model.Employee) error {
	return r.db.Create(&employees).Error
}

func (r *repository) UpdateEmployee(employee model.Employee) error {
	return r.db.Save(&employee).Error
}

func (r *repository) UpdateSchedule(schedule model.Schedule) error {
	return r.db.Save(&schedule).Error
}

func (r *repository) GetSchedule(employeeID int, weekType string) ([]model.Schedule, error) {
	var schedules []model.Schedule
	err := r.db.Where("employee_id = ? AND week_type = ?", employeeID, weekType).Find(&schedules).Error
	return schedules, err
}

func (r *repository) GetEmployees() ([]model.Employee, error) {
	var employees []model.Employee
	err := r.db.Find(&employees).Error
	return employees, err
}

// Additional methods for analytics or other operations can be defined here
