package server

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/umed-hotamov/house-rental/internal/app/config"

	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/echo/v4"
	"go.uber.org/fx"
)

func NewEchoRouter() *echo.Echo {
  e := echo.New()
  
  e.Use(middleware.Logger())
  e.GET("/ping", func(c echo.Context) error {
      time.Sleep(15 * time.Second)
      return c.String(http.StatusOK, "Pong")
  })

  return e
}

type ServerParams struct {
  fx.In

  Router *echo.Echo
  Cfg    *config.Config
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
        log.Printf("server started on port: %s\n", params.Cfg.Server.Port)
        server.ListenAndServe()
      }()

      return nil
    },
    OnStop: func(ctx context.Context) error {
      return server.Shutdown(ctx) 
    },
  })

  return server
}
