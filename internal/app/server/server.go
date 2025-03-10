package server

import (
	"context"
	"fmt"
	"net/http"
	"time"

	v1 "github.com/umed-hotamov/realty-service/internal/handler/api/v1"
	"github.com/umed-hotamov/realty-service/internal/app/config"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"go.uber.org/fx"
	"go.uber.org/zap"
)

func NewEchoRouter() *echo.Echo {
  e := echo.New()
  e.Use(middleware.Logger())
  e.GET("/ping", func(c echo.Context) error {
      return c.String(http.StatusOK, "Pong")
  })

  return e
}

type ServerParams struct {
  fx.In
  
  Router  *echo.Echo
  Cfg     *config.Config
  Logger  *zap.Logger
  Handler *v1.Handler
}

func NewServer(lc fx.Lifecycle, params ServerParams) *http.Server {
  server := &http.Server{
    Handler:      params.Router,
    Addr:         fmt.Sprintf(":%s", params.Cfg.Server.Port),
    ReadTimeout:  15 * time.Second,
    WriteTimeout: 15 * time.Second,
  }
  lc.Append(fx.Hook{
    OnStart: func(ctx context.Context) error {
      go func() {
        params.Logger.Info(fmt.Sprintf("server started on port: %s\n", params.Cfg.Server.Port))
        params.Logger.Fatal("stop server", zap.Error(server.ListenAndServe()))
      }()

      return nil
    },
    OnStop: func(ctx context.Context) error {
      return server.Shutdown(ctx) 
    },
  })

  return server
}
