package update_schedule

import (
	"strconv"

	"github.com/jfelipearaujo-healthmed/scheduler-service/internal/core/domain/dtos/schedule_dto"
	update_schedule_contract "github.com/jfelipearaujo-healthmed/scheduler-service/internal/core/domain/use_cases/schedule/update_schedule"
	"github.com/jfelipearaujo-healthmed/scheduler-service/internal/core/infrastructure/shared/http_response"
	"github.com/labstack/echo/v4"
)

type handler struct {
	useCase update_schedule_contract.UseCase
}

func NewHandler(useCase update_schedule_contract.UseCase) *handler {
	return &handler{
		useCase: useCase,
	}
}

func (h *handler) Handle(c echo.Context) error {
	ctx := c.Request().Context()

	req := new(schedule_dto.UpdateScheduleRequest)

	if err := c.Bind(req); err != nil {
		return http_response.BadRequest(c, "unable to parse the request body", err)
	}

	if !req.IsValid() {
		return http_response.BadRequest(c, "invalid request body", nil)
	}

	userId := c.Get("userId").(uint)
	scheduleId := c.Param("scheduleId")

	parsedScheduleId, err := strconv.ParseUint(scheduleId, 10, 64)
	if err != nil {
		return http_response.BadRequest(c, "invalid schedule id", err)
	}

	schedule, err := h.useCase.Execute(ctx, userId, uint(parsedScheduleId), req)
	if err != nil {
		return http_response.HandleErr(c, err)
	}

	return http_response.Created(c, schedule_dto.MapFromDomain(schedule))
}
