package schedule_dto

type CreateScheduleRequest struct {
	DateTimeAvailable string `json:"date_time_available" validate:"required"`
	Active            bool   `json:"active" validate:"required"`
}
