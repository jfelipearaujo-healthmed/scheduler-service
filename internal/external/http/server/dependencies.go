package server

import (
	schedule_repository_contract "github.com/jfelipearaujo-healthmed/scheduler-service/internal/core/domain/repositories/schedule"
	create_schedule_contract "github.com/jfelipearaujo-healthmed/scheduler-service/internal/core/domain/use_cases/schedule/create_shedule"
	"github.com/jfelipearaujo-healthmed/scheduler-service/internal/external/cache"
	"github.com/jfelipearaujo-healthmed/scheduler-service/internal/external/persistence"
)

type Dependencies struct {
	Cache     cache.Cache
	DbService *persistence.DbService

	ScheduleRepository schedule_repository_contract.Repository

	CreateScheduleUseCase create_schedule_contract.UseCase
}
