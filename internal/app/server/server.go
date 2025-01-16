package server

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/umed-hotamov/house-rental/internal/app/config"

	"github.com/labstack/echo/v4"
	"go.uber.org/fx"
	// "go.uber.org/zap"
)

type Config struct {
  Host string
  Port string
}

func NewEchoRouter() *echo.Echo {
  e := echo.New()

  e.GET("/ping", func(c echo.Context) error {
      return c.String(http.StatusOK, "Pong")
  })

  return e
}

type ServerParams struct {
  fx.In

  Router *echo.Echo
  Cfg    *config.Config
  // Logger *zap.Logger
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
      ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt)
      defer stop()

      go func() {
        log.Printf("server started on port: ", params.Cfg.Server.Port)
        server.ListenAndServe()
      }()

      go func() {
        <-ctx.Done()
        fmt.Print("Intrrupt signal received, shutting down...")
        stop()
      }()

      return nil
    },
    OnStop: func(ctx context.Context) error {
      ctx, cancel := context.WithTimeout(ctx, 30 * time.Second)
      defer cancel()
      return server.Shutdown(ctx) 
    },
  })

  return server
}
