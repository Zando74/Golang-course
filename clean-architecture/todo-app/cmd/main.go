package main

import (
	"todo-app/internal/app/adapter/controller"

	"github.com/labstack/echo/v4"
	"github.com/spf13/viper"
)

func main() {
	echoContext := echo.New()

	// for local development
	viper.SetDefault("PGHOST", "0.0.0.0")
	viper.SetDefault("PGUSER", "postgres")
	viper.SetDefault("PGPASSWORD", "postgres")

	viper.BindEnv("PGHOST", "PGHOST")
	viper.BindEnv("PGUSER", "PGUSER")
	viper.BindEnv("PGPASSWORD", "PGPASSWORD")

	controller.TaskRouter(echoContext)
	echoContext.Logger.Info(echoContext.Start(":3000"))
}
