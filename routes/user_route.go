package routes

import (
	"github.com/Zoe-2Fu/ps-tag-onboarding-go/handler"
	"github.com/labstack/echo/v4"
)

func UserRoute(e *echo.Echo) {
	e.GET("/user/:id", handler.Find)
	e.POST("/save", handler.Save)
}
