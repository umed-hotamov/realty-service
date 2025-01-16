package app

import (
	"net/http"

	"github.com/umed-hotamov/house-rental/internal/app/config"
	"github.com/umed-hotamov/house-rental/internal/app/server"

	"go.uber.org/fx"
)

func Run() {
  cfg := config.GetConfig()

  fx.New(
    fx.Provide(
      server.NewEchoRouter,
      server.NewServer,
    ),
    fx.Supply(cfg),
    fx.Invoke(func(*http.Server) {}),
  ).Run()
}
