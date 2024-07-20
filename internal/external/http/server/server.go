package server

import (
	"context"
	"fmt"
	"log/slog"
	"net/http"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	awsConfig "github.com/aws/aws-sdk-go-v2/config"
	create_schedule_uc "github.com/jfelipearaujo-healthmed/scheduler-service/internal/core/application/use_cases/schedule/create_schedule"
	delete_schedule_uc "github.com/jfelipearaujo-healthmed/scheduler-service/internal/core/application/use_cases/schedule/delete_schedule"
	get_schedule_by_id_uc "github.com/jfelipearaujo-healthmed/scheduler-service/internal/core/application/use_cases/schedule/get_schedule_by_id"
	list_schedules_uc "github.com/jfelipearaujo-healthmed/scheduler-service/internal/core/application/use_cases/schedule/list_schedules"
	update_schedule_uc "github.com/jfelipearaujo-healthmed/scheduler-service/internal/core/application/use_cases/schedule/update_schedule"
	"github.com/jfelipearaujo-healthmed/scheduler-service/internal/core/infrastructure/config"
	schedule_repository "github.com/jfelipearaujo-healthmed/scheduler-service/internal/core/infrastructure/repositories/schedule"
	"github.com/jfelipearaujo-healthmed/scheduler-service/internal/external/cache"
	"github.com/jfelipearaujo-healthmed/scheduler-service/internal/external/http/handlers/health"
	"github.com/jfelipearaujo-healthmed/scheduler-service/internal/external/http/handlers/schedule/create_schedule"
	"github.com/jfelipearaujo-healthmed/scheduler-service/internal/external/http/handlers/schedule/delete_schedule"
	"github.com/jfelipearaujo-healthmed/scheduler-service/internal/external/http/handlers/schedule/get_schedule_by_id"
	"github.com/jfelipearaujo-healthmed/scheduler-service/internal/external/http/handlers/schedule/list_schedules"
	"github.com/jfelipearaujo-healthmed/scheduler-service/internal/external/http/handlers/schedule/update_schedule"
	"github.com/jfelipearaujo-healthmed/scheduler-service/internal/external/http/middlewares/logger"
	"github.com/jfelipearaujo-healthmed/scheduler-service/internal/external/http/middlewares/role"
	"github.com/jfelipearaujo-healthmed/scheduler-service/internal/external/http/middlewares/token"
	"github.com/jfelipearaujo-healthmed/scheduler-service/internal/external/persistence"
	"github.com/jfelipearaujo-healthmed/scheduler-service/internal/external/secret"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type Server struct {
	Config *config.Config

	Dependencies
}

func NewServer(ctx context.Context, config *config.Config) (*Server, error) {
	cloudConfig, err := awsConfig.LoadDefaultConfig(ctx)
	if err != nil {
		slog.ErrorContext(ctx, "error getting aws config", "error", err)
		return nil, err
	}

	if config.CloudConfig.IsBaseEndpointSet() {
		cloudConfig.BaseEndpoint = aws.String(config.CloudConfig.BaseEndpoint)
	}

	secretService := secret.NewService(cloudConfig)

	dbUrl, err := secretService.GetSecret(ctx, config.DbConfig.UrlSecretName)
	if err != nil {
		slog.ErrorContext(ctx, "error getting secret", "secret_name", config.DbConfig.UrlSecretName, "error", err)
		return nil, err
	}

	cacheUrl, err := secretService.GetSecret(ctx, config.CacheConfig.HostSecretName)
	if err != nil {
		slog.ErrorContext(ctx, "error getting secret", "secret_name", config.CacheConfig.HostSecretName, "error", err)
		return nil, err
	}

	config.DbConfig.Url = dbUrl
	config.CacheConfig.Host = cacheUrl

	dbService := persistence.NewDbService()

	if err := dbService.Connect(config); err != nil {
		slog.ErrorContext(ctx, "error connecting to database", "error", err)
		return nil, err
	}

	cache := cache.NewRedisCache(ctx, config)

	scheduleRepository := schedule_repository.NewRepository(cache, dbService)

	return &Server{
		Config: config,
		Dependencies: Dependencies{
			Cache:     cache,
			DbService: dbService,

			ScheduleRepository: scheduleRepository,

			CreateScheduleUseCase:  create_schedule_uc.NewUseCase(scheduleRepository, config.ApiConfig.Location),
			ListSchedulesUseCase:   list_schedules_uc.NewUseCase(scheduleRepository),
			GetScheduleByIdUseCase: get_schedule_by_id_uc.NewUseCase(scheduleRepository),
			UpdateScheduleUseCase:  update_schedule_uc.NewUseCase(scheduleRepository, config.ApiConfig.Location),
			DeleteScheduleUseCase:  delete_schedule_uc.NewUseCase(scheduleRepository),
		},
	}, nil
}

func (s *Server) GetServer() *http.Server {
	return &http.Server{
		Addr:         fmt.Sprintf(":%d", s.Config.ApiConfig.Port),
		Handler:      s.RegisterRoutes(),
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}
}

func (s *Server) RegisterRoutes() http.Handler {
	e := echo.New()
	e.Use(logger.Middleware())
	e.Use(middleware.Recover())

	s.addHealthCheckRoutes(e)

	api := e.Group(fmt.Sprintf("/api/%s", s.Config.ApiConfig.ApiVersion))

	api.Use(token.Middleware())
	api.Use(role.Middleware(role.Doctor))
	s.addScheduleRoutes(api)

	return e
}

func (s *Server) addHealthCheckRoutes(e *echo.Echo) {
	healthHandler := health.NewHandler(s.DbService)

	e.GET("/health", healthHandler.Handle)
}

func (s *Server) addScheduleRoutes(g *echo.Group) {
	createScheduleHandler := create_schedule.NewHandler(s.CreateScheduleUseCase)
	listSchedulesHandler := list_schedules.NewHandler(s.ListSchedulesUseCase)
	getScheduleByIdHandler := get_schedule_by_id.NewHandler(s.GetScheduleByIdUseCase)
	updateScheduleHandler := update_schedule.NewHandler(s.UpdateScheduleUseCase)
	deleteScheduleHandler := delete_schedule.NewHandler(s.DeleteScheduleUseCase)

	g.POST("/schedules", createScheduleHandler.Handle)
	g.GET("/schedules", listSchedulesHandler.Handle)
	g.GET("/schedules/:scheduleId", getScheduleByIdHandler.Handle)
	g.PUT("/schedules/:scheduleId", updateScheduleHandler.Handle)
	g.DELETE("/schedules/:scheduleId", deleteScheduleHandler.Handle)
}
