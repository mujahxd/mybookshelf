package main

import (
	"github.com/labstack/echo/v4"
	"github.com/mujahxd/mybookshelf/auth"
	"github.com/mujahxd/mybookshelf/config"
	"github.com/mujahxd/mybookshelf/features/user"
	"github.com/mujahxd/mybookshelf/routes"
	"github.com/mujahxd/mybookshelf/utils/database"
)

func main() {
	e := echo.New()
	loadConfig := config.InitConfig()
	db := database.ConnectionDB(loadConfig)

	// database
	database.Migrate(db)
	userRepository := user.NewRepository(db)

	userService := user.NewService(userRepository)

	authService := auth.NewService()

	userHandler := user.NewHandler(userService, authService)
	routes.InitRoute(e, userHandler)

	e.Logger.Fatal(e.Start(":8000"))
}
