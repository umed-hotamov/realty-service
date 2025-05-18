package app

import (
	"net/http"

	"github.com/umed-hotamov/realty-service/internal/app/config"
	"github.com/umed-hotamov/realty-service/internal/app/server"
	manager "github.com/umed-hotamov/realty-service/internal/authHelpers/authManager"
	"github.com/umed-hotamov/realty-service/internal/authHelpers/storage"
	v1 "github.com/umed-hotamov/realty-service/internal/handler/api/v1"
	"github.com/umed-hotamov/realty-service/internal/repository"
	pg "github.com/umed-hotamov/realty-service/internal/repository/postgres"
	"github.com/umed-hotamov/realty-service/internal/service"
	"github.com/umed-hotamov/realty-service/internal/service/providers"
	database "github.com/umed-hotamov/realty-service/pkg/database/postgres"
	"github.com/umed-hotamov/realty-service/pkg/database/redis"
	"github.com/umed-hotamov/realty-service/pkg/logger"

	"go.uber.org/fx"
)

func Run() {
	cfg := config.GetConfig()
	logger := logger.NewLogger(cfg)

	logger.Info("start app")

	fx.New(
		fx.Provide(
			server.NewEchoRouter,
			server.NewServer,
			v1.NewHandler,

			database.NewPostgresDB,
			redis.NewClient,

			storage.NewSessionStorage,
			manager.NewAuthManager,

			fx.Annotate(
				pg.NewUserRepo,
				fx.As(new(repository.IUserRepo)),
			),
			fx.Annotate(
				pg.NewPropertyRepo,
				fx.As(new(repository.IPropertyRepo)),
			),
			fx.Annotate(
				pg.NewPgLisingRepo,
				fx.As(new(repository.IListingRepo)),
			),

			fx.Annotate(
				providers.NewUserService,
				fx.As(new(service.IUserService)),
			),
			fx.Annotate(
				providers.NewPropertyService,
				fx.As(new(service.IPropertyService)),
			),
			fx.Annotate(
				providers.NewListingService,
				fx.As(new(service.IListingService)),
			),
			fx.Annotate(
				providers.NewAuthService,
				fx.As(new(service.IAuthService)),
			),
		),
		fx.Supply(cfg, logger),
		fx.Invoke(func(*http.Server) {}),
	).Run()
}
