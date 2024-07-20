package list_schedules

import (
	"github.com/jfelipearaujo-healthmed/scheduler-service/internal/core/domain/dtos/schedule_dto"
	list_schedules_contract "github.com/jfelipearaujo-healthmed/scheduler-service/internal/core/domain/use_cases/schedule/list_schedules"
	"github.com/jfelipearaujo-healthmed/scheduler-service/internal/core/infrastructure/shared/http_response"
	"github.com/labstack/echo/v4"
)

type handler struct {
	useCase list_schedules_contract.UseCase
}

func NewHandler(useCase list_schedules_contract.UseCase) *handler {
	return &handler{
		useCase: useCase,
	}
}

func (h *handler) Handle(c echo.Context) error {
	ctx := c.Request().Context()

	userId := c.Get("userId").(uint)

	schedules, err := h.useCase.Execute(ctx, userId)
	if err != nil {
		return http_response.HandleErr(c, err)
	}

	return http_response.OK(c, schedule_dto.MapFromDomainSlice(schedules))
}
