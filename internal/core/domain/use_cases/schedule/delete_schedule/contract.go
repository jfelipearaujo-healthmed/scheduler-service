package delete_schedule_contract

import (
	"context"
)

type UseCase interface {
	Execute(ctx context.Context, doctorID uint, scheduleID uint) error
}
