package server

import (
	schedule_repository_contract "github.com/jfelipearaujo-healthmed/scheduler-service/internal/core/domain/repositories/schedule"
	create_schedule_contract "github.com/jfelipearaujo-healthmed/scheduler-service/internal/core/domain/use_cases/schedule/create_shedule"
	delete_schedule_contract "github.com/jfelipearaujo-healthmed/scheduler-service/internal/core/domain/use_cases/schedule/delete_schedule"
	get_schedule_by_id_contract "github.com/jfelipearaujo-healthmed/scheduler-service/internal/core/domain/use_cases/schedule/get_schedule_by_id"
	list_schedules_contract "github.com/jfelipearaujo-healthmed/scheduler-service/internal/core/domain/use_cases/schedule/list_schedules"
	update_schedule_contract "github.com/jfelipearaujo-healthmed/scheduler-service/internal/core/domain/use_cases/schedule/update_schedule"
	"github.com/jfelipearaujo-healthmed/scheduler-service/internal/external/cache"
	"github.com/jfelipearaujo-healthmed/scheduler-service/internal/external/persistence"
)

type Dependencies struct {
	Cache     cache.Cache
	DbService *persistence.DbService

	ScheduleRepository schedule_repository_contract.Repository

	CreateScheduleUseCase  create_schedule_contract.UseCase
	ListSchedulesUseCase   list_schedules_contract.UseCase
	GetScheduleByIdUseCase get_schedule_by_id_contract.UseCase
	UpdateScheduleUseCase  update_schedule_contract.UseCase
	DeleteScheduleUseCase  delete_schedule_contract.UseCase
}
