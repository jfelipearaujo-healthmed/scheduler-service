package schedule_dto

type CreateScheduleRequest struct {
	DoctorID          uint   `json:"doctor_id" validate:"required"`
	DateTimeAvailable string `json:"date_time_available" validate:"required"`
	Active            bool   `json:"active" validate:"required"`
}
