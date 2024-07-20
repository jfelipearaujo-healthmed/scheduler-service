package get_schedule_by_id_uc

import (
	"context"

	"github.com/jfelipearaujo-healthmed/scheduler-service/internal/core/domain/entities"
	schedule_repository_contract "github.com/jfelipearaujo-healthmed/scheduler-service/internal/core/domain/repositories/schedule"
	get_schedule_by_id_contract "github.com/jfelipearaujo-healthmed/scheduler-service/internal/core/domain/use_cases/schedule/get_schedule_by_id"
)

type useCase struct {
	repository schedule_repository_contract.Repository
}

func NewUseCase(repository schedule_repository_contract.Repository) get_schedule_by_id_contract.UseCase {
	return &useCase{
		repository: repository,
	}
}

func (uc *useCase) Execute(ctx context.Context, doctorID uint, scheduleID uint) (*entities.Schedule, error) {
	return uc.repository.GetByID(ctx, doctorID, scheduleID)
}
