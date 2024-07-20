package schedule_repository_contract

import (
	"context"
	"time"

	"github.com/jfelipearaujo-healthmed/scheduler-service/internal/core/domain/entities"
)

type ListFilter struct {
	DoctorID          *uint
	DateTImeAvailable *time.Time
	Active            *bool
}

type Repository interface {
	GetByID(ctx context.Context, id uint) (*entities.Schedule, error)
	GetByDoctorIDAndDateTimeAvailable(ctx context.Context, doctorID uint, dateTimeAvailable time.Time) (*entities.Schedule, error)
	List(ctx context.Context, filter *ListFilter) ([]entities.Schedule, error)
	Create(ctx context.Context, schedule *entities.Schedule) (*entities.Schedule, error)
	Update(ctx context.Context, schedule *entities.Schedule) (*entities.Schedule, error)
	Delete(ctx context.Context, id uint) error
}
