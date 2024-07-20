package schedule_repository

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/jfelipearaujo-healthmed/scheduler-service/internal/core/domain/entities"
	schedule_repository_contract "github.com/jfelipearaujo-healthmed/scheduler-service/internal/core/domain/repositories/schedule"
	"github.com/jfelipearaujo-healthmed/scheduler-service/internal/core/infrastructure/shared/app_error"
	"github.com/jfelipearaujo-healthmed/scheduler-service/internal/external/cache"
	"github.com/jfelipearaujo-healthmed/scheduler-service/internal/external/persistence"
	"gorm.io/gorm"
)

const (
	cacheKey string        = "schedule:%d:%d"
	ttl      time.Duration = time.Hour * 24
)

type repository struct {
	cache     cache.Cache
	dbService *persistence.DbService
}

func NewRepository(cache cache.Cache, dbService *persistence.DbService) schedule_repository_contract.Repository {
	return &repository{
		cache:     cache,
		dbService: dbService,
	}
}

func (rp *repository) GetByID(ctx context.Context, doctorID uint, scheduleID uint) (*entities.Schedule, error) {
	return cache.WithCache(ctx, rp.cache, fmt.Sprintf(cacheKey, doctorID, scheduleID), ttl, func() (*entities.Schedule, error) {
		tx := rp.dbService.Instance.WithContext(ctx)

		schedule := new(entities.Schedule)
		result := tx.Where("doctor_id = ? AND id = ?", doctorID, scheduleID).First(schedule)
		if result.Error != nil {
			if errors.Is(result.Error, gorm.ErrRecordNotFound) {
				return nil, app_error.New(http.StatusNotFound, fmt.Sprintf("schedule with id %d not found", scheduleID))
			}

			return nil, result.Error
		}

		return schedule, nil
	})
}

func (rp *repository) GetByDoctorIDAndDateTimeAvailable(ctx context.Context, doctorID uint, dateTimeAvailable time.Time) (*entities.Schedule, error) {
	tx := rp.dbService.Instance.WithContext(ctx)

	schedule := new(entities.Schedule)
	result := tx.Where("doctor_id = ? AND date_time_available = ?", doctorID, dateTimeAvailable).First(schedule)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, app_error.New(http.StatusNotFound, fmt.Sprintf("schedule with doctor id %d and date time available %s not found", doctorID, dateTimeAvailable))
		}

		return nil, result.Error
	}

	return schedule, nil
}

func (rp *repository) List(ctx context.Context, filter *schedule_repository_contract.ListFilter) ([]entities.Schedule, error) {
	tx := rp.dbService.Instance.WithContext(ctx)

	schedules := new([]entities.Schedule)

	query := tx

	if filter.DoctorID != nil {
		query = query.Where("doctor_id = ?", *filter.DoctorID)
	}

	if filter.DateTimeAvailable != nil {
		query = query.Where("date_time_available = ?", *filter.DateTimeAvailable)
	}

	if filter.Active != nil {
		query = query.Where("active = ?", *filter.Active)
	}

	if err := query.Find(&schedules).Error; err != nil {
		return nil, err
	}

	return *schedules, nil
}

func (rp *repository) Create(ctx context.Context, schedule *entities.Schedule) (*entities.Schedule, error) {
	tx := rp.dbService.Instance.WithContext(ctx)

	if err := tx.Create(schedule).Error; err != nil {
		return nil, err
	}

	return schedule, nil
}

func (rp *repository) Update(ctx context.Context, schedule *entities.Schedule) (*entities.Schedule, error) {
	tx := rp.dbService.Instance.WithContext(ctx)

	result := tx.Model(schedule).Save(schedule)

	if err := result.Error; err != nil {
		return nil, err
	}

	return schedule, nil
}

func (rp *repository) Delete(ctx context.Context, id uint) error {
	tx := rp.dbService.Instance.WithContext(ctx)

	if err := tx.Delete(&entities.Schedule{}, id).Error; err != nil {
		return err
	}

	return nil
}
