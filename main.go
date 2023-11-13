package main

import (
	"github.com/Zoe-2Fu/ps-tag-onboarding-go/configs"
	"github.com/Zoe-2Fu/ps-tag-onboarding-go/handler"
	"github.com/Zoe-2Fu/ps-tag-onboarding-go/routes"
	validator "github.com/Zoe-2Fu/ps-tag-onboarding-go/validators"
	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	mongoClient := configs.ConnectMongoDB()
	db := mongoClient.NewMongoDB()

	userRepo := configs.NewUserRepo(db)
	userValidator := validator.NewUserValidator(userRepo)
	userHandler := handler.NewUserHandler(userRepo, userValidator)

	routes.UserRoute(e, *userHandler)
	e.Use(handler.HandleError)

	e.Logger.Fatal(e.Start(":6000"))
}
