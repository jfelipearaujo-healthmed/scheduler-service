package delete_schedule

import (
	"strconv"

	delete_schedule_contract "github.com/jfelipearaujo-healthmed/scheduler-service/internal/core/domain/use_cases/schedule/delete_schedule"
	"github.com/jfelipearaujo-healthmed/scheduler-service/internal/core/infrastructure/shared/http_response"
	"github.com/labstack/echo/v4"
)

type handler struct {
	useCase delete_schedule_contract.UseCase
}

func NewHandler(useCase delete_schedule_contract.UseCase) *handler {
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

	if err := h.useCase.Execute(ctx, userId, uint(parsedScheduleId)); err != nil {
		return http_response.HandleErr(c, err)
	}

	return http_response.NoContent(c)
}
