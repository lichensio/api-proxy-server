package model

import "time"

// Employee represents an employee record in the database and the JSON structure.
type Employee struct {
	ID        int       `db:"id" json:"id,omitempty"`
	Name      string    `db:"name" json:"name"`
	StartDate time.Time `db:"start_date" json:"startDate"`
	Weeks     []Week    `json:"weeks,omitempty"` // This is not directly stored in DB, managed via Schedule entries
}

// Week represents a flexible struct to accommodate the dynamic "A" and "B" week types with schedules.
// In practice, this struct needs to be adapted or handled dynamically in the logic rather than directly stored.
type Week struct {
	WeekType string        `json:"weekType"` // "A" or "B"
	Days     []DaySchedule `json:"days"`
}

// DaySchedule represents the schedule for a single day, matching the schedules table structure.
type DaySchedule struct {
	DayName   string    `db:"day_name" json:"dayName"`
	StartTime time.Time `db:"start_time" json:"start"`
	EndTime   time.Time `db:"end_time" json:"end"`
}

// Schedule represents the schedule of an employee, aligning with the schedules table.
type Schedule struct {
	ID         int       `db:"id" json:"id,omitempty"`
	EmployeeID int       `db:"employee_id" json:"employeeId"`
	WeekType   string    `db:"week_type" json:"weekType"` // "A" or "B"
	DayName    string    `db:"day_name" json:"dayName"`
	StartTime  time.Time `db:"start_time" json:"startTime"`
	EndTime    time.Time `db:"end_time" json:"endTime"`
}
