package create_schedule

import (
	"github.com/jfelipearaujo-healthmed/scheduler-service/internal/core/domain/dtos/schedule_dto"
	create_schedule_contract "github.com/jfelipearaujo-healthmed/scheduler-service/internal/core/domain/use_cases/schedule/create_shedule"
	"github.com/jfelipearaujo-healthmed/scheduler-service/internal/core/infrastructure/shared/http_response"
	"github.com/jfelipearaujo-healthmed/scheduler-service/internal/core/infrastructure/shared/validator"
	"github.com/labstack/echo/v4"
)

type handler struct {
	useCase create_schedule_contract.UseCase
}

func NewHandler(useCase create_schedule_contract.UseCase) *handler {
	return &handler{
		useCase: useCase,
	}
}

func (h *handler) Handle(c echo.Context) error {
	ctx := c.Request().Context()

	req := new(schedule_dto.CreateScheduleRequest)

	if err := c.Bind(req); err != nil {
		return http_response.BadRequest(c, "unable to parse the request body", err)
	}

	if err := validator.Validate(req); err != nil {
		return http_response.UnprocessableEntity(c, "invalid request body", err)
	}

	schedule, err := h.useCase.Execute(ctx, req)
	if err != nil {
		return http_response.HandleErr(c, err)
	}

	return http_response.Created(c, schedule_dto.MapFromDomain(schedule))
}
