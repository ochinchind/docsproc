// Package app configures and runs application.
package app

import (
	"fmt"
	"github.com/go-redis/redis"
	"os"
	"os/signal"
	"syscall"

	"github.com/gin-gonic/gin"

	"github.com/ochinchind/docsproc/config"
	v1 "github.com/ochinchind/docsproc/internal/controller/http/v1"
	"github.com/ochinchind/docsproc/internal/usecase"
	"github.com/ochinchind/docsproc/internal/usecase/repo"
	"github.com/ochinchind/docsproc/internal/usecase/webapi"
	"github.com/ochinchind/docsproc/pkg/casbin"
	"github.com/ochinchind/docsproc/pkg/httpserver"
	"github.com/ochinchind/docsproc/pkg/logger"
	"github.com/ochinchind/docsproc/pkg/postgres"
)

// Run creates objects via constructors.
func Run(cfg *config.Config) {
	l := logger.New(cfg.Log.Level)

	// Repository
	pg, err := postgres.New(cfg.PG.PGURL)
	if err != nil {
		l.Fatal(fmt.Errorf("app - Run - postgres.New: %w", err))
	}

	// Use case
	googleOAuthUseCase := usecase.NewGoogleOAuthUseCase(
		webapi.New(),
		repo.New(pg),
	)

	userUseCase := usecase.NewUserUseCase(
		repo.New(pg),
	)

	specialtyUseCase := usecase.NewSpecialtyUseCase(
		repo.NewSpecialtyRepo(pg),
	)

	qualificationUseCase := usecase.NewQualificationUseCase(
		repo.NewQualificationRepo(pg),
		repo.NewSpecialtyRepo(pg),
	)

	competencyUseCase := usecase.NewCompetencyUseCase(
		repo.NewCompetencyRepo(pg),
		repo.NewSpecialtyRepo(pg),
	)

	disciplineUseCase := usecase.NewDisciplineUseCase(
		repo.NewDisciplineRepo(pg),
		repo.NewQualificationRepo(pg),
		repo.NewCompetencyRepo(pg),
	)

	disciplineStudyPlanUseCase := usecase.NewDisciplineStudyPlanUseCase(
		repo.NewDisciplineStudyPlanRepo(pg),
		repo.NewDisciplineRepo(pg),
	)

	disciplineModuleUseCase := usecase.NewDisciplineModuleUseCase(
		repo.NewDisciplineModuleRepo(pg),
		repo.NewDisciplineRepo(pg),
	)

	disciplineModuleChapterUseCase := usecase.NewDisciplineModuleChapterUseCase(
		repo.NewDisciplineModuleChapterRepo(pg),
		repo.NewDisciplineModuleRepo(pg),
	)

	disciplineModuleChapterTopicUseCase := usecase.NewDisciplineModuleChapterTopicUseCase(
		repo.NewDisciplineModuleChapterTopicRepo(pg),
		repo.NewDisciplineModuleChapterRepo(pg),
	)

	authUseCase := usecase.NewAuthUseCase(
		repo.New(pg),
	)

	services := usecase.NewServices(
		googleOAuthUseCase,
		userUseCase,
		authUseCase,
		specialtyUseCase,
		qualificationUseCase,
		competencyUseCase,
		disciplineUseCase,
		disciplineStudyPlanUseCase,
		disciplineModuleUseCase,
		disciplineModuleChapterUseCase,
		disciplineModuleChapterTopicUseCase,
	)

	err = pg.Connect(cfg)
	if err != nil {
		l.Fatal(fmt.Errorf("app - Run - postgres.connect: %w", err))
	}

	err = pg.Migrate()
	if err != nil {
		l.Fatal(fmt.Errorf("app - Run - postgres.Migrate: %w", err))
	}

	// init casbin
	casbinEnforcer, err := casbin.InitCasbin()
	if err != nil {
		l.Fatal(err)
	}

	// init redis
	rd := redis.NewClient(&redis.Options{
		Addr:     cfg.Redis.RedisURL,      // Redis address
		Password: cfg.Redis.RedisPassword, // No password by default
		DB:       0,                       // Default DB
	})

	// RabbitMQ RPC Server
	//rmqRouter := amqprpc.NewRouter(services)
	//
	//rmqServer, err := server.New(cfg.RMQ.RMQURL, cfg.RMQ.ServerExchange, rmqRouter, l)
	//if err != nil {
	//	l.Fatal(fmt.Errorf("app - Run - rmqServer - server.New: %w", err))
	//}

	_, err = rd.Ping().Result()
	if err != nil {
		l.Fatal("Could not connect to Redis: %v", err)
	}

	// HTTP Server
	handler := gin.New()
	handler.Static("/uploads", "./uploads")
	handler.MaxMultipartMemory = 200 << 20
	v1.NewRouter(handler, l, services, casbinEnforcer, rd)
	httpServer := httpserver.New(handler, httpserver.Port(cfg.HTTP.Port))

	// Waiting signal
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)

	select {
	case s := <-interrupt:
		l.Info("app - Run - signal: " + s.String())
	case err = <-httpServer.Notify():
		l.Error(fmt.Errorf("app - Run - httpServer.Notify: %w", err))
		//case err = <-rmqServer.Notify():
		//	l.Error(fmt.Errorf("app - Run - rmqServer.Notify: %w", err))

	}

	// Shutdown
	err = httpServer.Shutdown()
	if err != nil {
		l.Error(fmt.Errorf("app - Run - httpServer.Shutdown: %w", err))
	}

	//err = rmqServer.Shutdown()
	//if err != nil {
	//	l.Error(fmt.Errorf("app - Run - rmqServer.Shutdown: %w", err))
	//}

}
