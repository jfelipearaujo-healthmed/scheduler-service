package entities

import (
	"time"

	"gorm.io/gorm"
)

type Schedule struct {
	gorm.Model

	DoctorID          uint      `json:"doctor_id" gorm:"uniqueIndex:idx_doctor_date_time"`
	DateTimeAvailable time.Time `json:"date_time_available" gorm:"uniqueIndex:idx_doctor_date_time"`
	Active            bool      `json:"active"`
}
