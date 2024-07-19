package schedule_dto

import "time"

type UpdateScheduleRequest struct {
	DoctorID          *uint      `json:"doctor_id"`
	DateTimeAvailable *time.Time `json:"date_time_available"`
	Active            *bool      `json:"active"`
}
