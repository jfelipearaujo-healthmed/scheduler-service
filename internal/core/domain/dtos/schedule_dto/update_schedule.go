package schedule_dto

type UpdateScheduleRequest struct {
	DateTimeAvailable *string `json:"date_time_available"`
	Active            *bool   `json:"active"`
}

func (r *UpdateScheduleRequest) IsValid() bool {
	return r.DateTimeAvailable != nil || r.Active != nil
}
