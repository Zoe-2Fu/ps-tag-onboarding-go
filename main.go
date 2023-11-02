package main

import (
	"github.com/Zoe-2Fu/ps-tag-onboarding-go/configs"
	"github.com/Zoe-2Fu/ps-tag-onboarding-go/handler"
	"github.com/Zoe-2Fu/ps-tag-onboarding-go/routes"
	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	e.Use(handler.HandleError)

	configs.ConnectMongoDB()
	routes.UserRoute(e)

	e.Logger.Fatal(e.Start(":6000"))
}
