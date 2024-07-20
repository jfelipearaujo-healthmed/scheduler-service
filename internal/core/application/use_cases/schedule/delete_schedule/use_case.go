package delete_schedule_uc

import (
	"context"

	schedule_repository_contract "github.com/jfelipearaujo-healthmed/scheduler-service/internal/core/domain/repositories/schedule"
	delete_schedule_contract "github.com/jfelipearaujo-healthmed/scheduler-service/internal/core/domain/use_cases/schedule/delete_schedule"
)

type useCase struct {
	repository schedule_repository_contract.Repository
}

func NewUseCase(repository schedule_repository_contract.Repository) delete_schedule_contract.UseCase {
	return &useCase{
		repository: repository,
	}
}

func (uc *useCase) Execute(ctx context.Context, doctorID uint, scheduleID uint) error {
	_, err := uc.repository.GetByID(ctx, doctorID, scheduleID)
	if err != nil {
		return err
	}

	return uc.repository.Delete(ctx, doctorID, scheduleID)
}
