package handler

import (
	"context"
	"net/http"
	"time"

	"github.com/Zoe-2Fu/ps-tag-onboarding-go/models"
	errs "github.com/Zoe-2Fu/ps-tag-onboarding-go/models/error"
	"github.com/labstack/echo/v4"
)

type userRepo interface {
	Find(ctx echo.Context, id string) (models.User, error)
	Save(ctx context.Context, user models.User) error
}

type userValidator interface {
	ValidateUserDetails(c echo.Context, user models.User) *errs.ErrorMessage
}

type UserHandler struct {
	userRepo  userRepo
	validator userValidator
}

func NewUserHandler(repo userRepo, validator userValidator) *UserHandler {
	return &UserHandler{userRepo: repo, validator: validator}
}

func (h *UserHandler) Find(c echo.Context) error {
	id := c.Param("id")
	var user models.User

	user, err := h.userRepo.Find(c, id)
	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, errs.ErrorMessage{
			Error:   errs.ErrorStatusNotFound,
			Details: []string{"User not found"},
		})
	}
	return c.JSON(http.StatusOK, user)
}

func (h *UserHandler) Save(c echo.Context) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	user := new(models.User)
	if err := c.Bind(user); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, errs.ErrorMessage{
			Error:   errs.ErrorBadRequest,
			Details: []string{"Can't bind values"},
		})
	}

	if validationErr := h.validator.ValidateUserDetails(c, *user); validationErr != nil {
		return echo.NewHTTPError(http.StatusBadRequest, validationErr)
	}

	if err := h.userRepo.Save(ctx, *user); err != nil {
		return err
	}

	return c.JSON(http.StatusCreated, user)
}
