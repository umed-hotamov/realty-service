package app

import (
	"net/http"

	v1 "github.com/umed-hotamov/realty-service/internal/handler/api/v1"
	"github.com/umed-hotamov/realty-service/internal/app/config"
	"github.com/umed-hotamov/realty-service/internal/app/server"
	"github.com/umed-hotamov/realty-service/internal/repository"
	pg "github.com/umed-hotamov/realty-service/internal/repository/postgres"
	"github.com/umed-hotamov/realty-service/internal/service"
	"github.com/umed-hotamov/realty-service/internal/service/providers"
	database "github.com/umed-hotamov/realty-service/pkg/database/postgres"
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
    
      fx.Annotate(
        pg.NewUserRepo,
        fx.As(new(repository.IUserRepo)),        
      ),

      fx.Annotate(
        providers.NewUserService,
        fx.As(new(service.IUserService)),
      ),
    ),
    fx.Supply(cfg, logger),
    fx.Invoke(func(*http.Server) {}),
  ).Run()
}
