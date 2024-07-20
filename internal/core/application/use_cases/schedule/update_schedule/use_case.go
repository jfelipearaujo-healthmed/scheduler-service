package update_schedule_uc

import (
	"context"
	"net/http"
	"time"

	"github.com/jfelipearaujo-healthmed/scheduler-service/internal/core/domain/dtos/schedule_dto"
	"github.com/jfelipearaujo-healthmed/scheduler-service/internal/core/domain/entities"
	schedule_repository_contract "github.com/jfelipearaujo-healthmed/scheduler-service/internal/core/domain/repositories/schedule"
	update_schedule_contract "github.com/jfelipearaujo-healthmed/scheduler-service/internal/core/domain/use_cases/schedule/update_schedule"
	"github.com/jfelipearaujo-healthmed/scheduler-service/internal/core/infrastructure/shared/app_error"
)

const (
	dateTimeLayout = "2006-01-02 15:04"
)

type useCase struct {
	repository schedule_repository_contract.Repository
	location   *time.Location
}

func NewUseCase(repository schedule_repository_contract.Repository, location *time.Location) update_schedule_contract.UseCase {
	return &useCase{
		repository: repository,
		location:   location,
	}
}

func (uc *useCase) Execute(ctx context.Context, doctorID uint, scheduleID uint, request *schedule_dto.UpdateScheduleRequest) (*entities.Schedule, error) {
	schedule, err := uc.repository.GetByID(ctx, doctorID, scheduleID)
	if err != nil {
		return nil, err
	}

	if request.DateTimeAvailable != nil {
		parsedTime, err := time.ParseInLocation(dateTimeLayout, *request.DateTimeAvailable, uc.location)
		if err != nil {
			return nil, app_error.New(http.StatusBadRequest, "unable to parse the date and time provided")
		}

		year, month, day := parsedTime.Date()
		hour, minute, _ := parsedTime.Clock()

		finalTime := time.Date(year, month, day, hour, minute, 0, 0, uc.location)

		if finalTime.Before(time.Now()) {
			return nil, app_error.New(http.StatusBadRequest, "date and time must be in the future")
		}

		schedule.DateTimeAvailable = finalTime
	}

	if request.Active != nil {
		schedule.Active = *request.Active
	}

	existingSchedule, err := uc.repository.GetByDoctorIDAndDateTimeAvailable(ctx, doctorID, schedule.DateTimeAvailable)
	if err != nil && !app_error.IsAppError(err) {
		return nil, err
	}

	if existingSchedule != nil && existingSchedule.ID != schedule.ID {
		return nil, app_error.New(http.StatusConflict, "schedule already exists for this doctor and date time")
	}

	return uc.repository.Update(ctx, doctorID, schedule)
}
