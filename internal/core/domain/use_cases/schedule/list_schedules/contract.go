package list_schedules_contract

import (
	"context"

	"github.com/jfelipearaujo-healthmed/scheduler-service/internal/core/domain/entities"
)

type UseCase interface {
	Execute(ctx context.Context, doctorID uint) ([]entities.Schedule, error)
}
