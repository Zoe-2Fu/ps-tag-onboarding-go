package routes

import (
	handler "github.com/Zoe-2Fu/ps-tag-onboarding-go/handlers"
	"github.com/labstack/echo/v4"
)

func UserRoute(e *echo.Echo, handler handler.UserHandler) {
	e.GET("/find/:id", handler.Find)
	e.POST("/save", handler.Save)
}
