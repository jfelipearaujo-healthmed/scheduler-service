package get_schedule_by_id_contract

import (
	"context"

	"github.com/jfelipearaujo-healthmed/scheduler-service/internal/core/domain/entities"
)

type UseCase interface {
	Execute(ctx context.Context, doctorID uint, scheduleID uint) (*entities.Schedule, error)
}
