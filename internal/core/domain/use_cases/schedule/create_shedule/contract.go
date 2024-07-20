package create_schedule_contract

import (
	"context"

	"github.com/jfelipearaujo-healthmed/scheduler-service/internal/core/domain/dtos/schedule_dto"
	"github.com/jfelipearaujo-healthmed/scheduler-service/internal/core/domain/entities"
)

type UseCase interface {
	Execute(ctx context.Context, doctorID uint, request *schedule_dto.CreateScheduleRequest) (*entities.Schedule, error)
}
