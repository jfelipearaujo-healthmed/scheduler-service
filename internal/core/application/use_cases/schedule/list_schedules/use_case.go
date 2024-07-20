package list_schedules_uc

import (
	"context"

	"github.com/jfelipearaujo-healthmed/scheduler-service/internal/core/domain/entities"
	schedule_repository_contract "github.com/jfelipearaujo-healthmed/scheduler-service/internal/core/domain/repositories/schedule"
	list_schedules_contract "github.com/jfelipearaujo-healthmed/scheduler-service/internal/core/domain/use_cases/schedule/list_schedules"
)

type useCase struct {
	repository schedule_repository_contract.Repository
}

func NewUseCase(repository schedule_repository_contract.Repository) list_schedules_contract.UseCase {
	return &useCase{
		repository: repository,
	}
}

func (uc *useCase) Execute(ctx context.Context, doctorID uint) ([]entities.Schedule, error) {
	filter := &schedule_repository_contract.ListFilter{
		DoctorID: &doctorID,
	}

	return uc.repository.List(ctx, filter)
}
