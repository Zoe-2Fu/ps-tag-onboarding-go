package main

import (
	handler "github.com/Zoe-2Fu/ps-tag-onboarding-go/handlers"
	"github.com/Zoe-2Fu/ps-tag-onboarding-go/mongo"
	"github.com/Zoe-2Fu/ps-tag-onboarding-go/routes"
	validator "github.com/Zoe-2Fu/ps-tag-onboarding-go/validators"
	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	mongoClient := mongo.ConnectMongoDB()
	db := mongoClient.NewMongoDB()

	userRepo := mongo.NewUserRepo(db)
	userValidator := validator.NewUserValidator(userRepo)
	userHandler := handler.NewUserHandler(userRepo, userValidator)

	routes.UserRoute(e, *userHandler)
	e.Use(handler.HandleError)

	e.Logger.Fatal(e.Start(":6000"))
}
