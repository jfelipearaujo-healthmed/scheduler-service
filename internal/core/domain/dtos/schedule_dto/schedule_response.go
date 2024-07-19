package schedule_dto

import (
	"time"

	"github.com/jfelipearaujo-healthmed/scheduler-service/internal/core/domain/entities"
)

type ScheduleResponse struct {
	ID uint `json:"id"`

	DoctorID          uint      `json:"doctor_id"`
	DateTimeAvailable time.Time `json:"date_time_available"`
	Active            bool      `json:"active"`

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func MapFromDomain(schedule *entities.Schedule) *ScheduleResponse {
	return &ScheduleResponse{
		ID:                schedule.ID,
		DoctorID:          schedule.DoctorID,
		DateTimeAvailable: schedule.DateTimeAvailable,
		Active:            schedule.Active,
		CreatedAt:         schedule.CreatedAt,
		UpdatedAt:         schedule.UpdatedAt,
	}
}
