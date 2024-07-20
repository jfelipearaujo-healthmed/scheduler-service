package get_schedule_by_id

import (
	"strconv"

	"github.com/jfelipearaujo-healthmed/scheduler-service/internal/core/domain/dtos/schedule_dto"
	get_schedule_by_id_contract "github.com/jfelipearaujo-healthmed/scheduler-service/internal/core/domain/use_cases/schedule/get_schedule_by_id"
	"github.com/jfelipearaujo-healthmed/scheduler-service/internal/core/infrastructure/shared/http_response"
	"github.com/labstack/echo/v4"
)

type handler struct {
	useCase get_schedule_by_id_contract.UseCase
}

func NewHandler(useCase get_schedule_by_id_contract.UseCase) *handler {
	return &handler{
		useCase: useCase,
	}
}

func (h *handler) Handle(c echo.Context) error {
	ctx := c.Request().Context()

	userId := c.Get("userId").(uint)
	scheduleId := c.Param("scheduleId")

	parsedScheduleId, err := strconv.ParseUint(scheduleId, 10, 64)
	if err != nil {
		return http_response.BadRequest(c, "invalid schedule id", err)
	}

	schedule, err := h.useCase.Execute(ctx, userId, uint(parsedScheduleId))
	if err != nil {
		return http_response.HandleErr(c, err)
	}

	return http_response.OK(c, schedule_dto.MapFromDomain(schedule))
}
