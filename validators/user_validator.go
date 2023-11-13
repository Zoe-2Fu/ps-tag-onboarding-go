package validator

import (
	"strings"

	"github.com/Zoe-2Fu/ps-tag-onboarding-go/model"
	errs "github.com/Zoe-2Fu/ps-tag-onboarding-go/model/error"
	"github.com/labstack/echo/v4"
)

type userRepo interface {
	ValidaiteUserExisted(ctx echo.Context, user model.User) bool
}

type UserValidator struct {
	userRepo userRepo
}

func NewUserValidator(repo userRepo) *UserValidator {
	return &UserValidator{userRepo: repo}
}

func (v *UserValidator) ValidateUserDetails(c echo.Context, user model.User) *errs.ErrorMessage {
	var errorDetails []string

	isExist := v.userRepo.ValidaiteUserExisted(c, user)
	if isExist {
		errMsg := errs.NewErrorMessage(errs.ResponseValidationFailed, errs.ErrorNameUnique)

		return &errMsg
	}

	// check if the user.Firstname && user.Lastname is not null
	if len(user.FirstName) == 0 || len(user.LastName) == 0 {
		errorDetails = append(errorDetails, errs.ErrorNameRequired)
	}

	if len(user.Email) == 0 {
		errorDetails = append(errorDetails, errs.ErrorEmailRequired)
	} else if !strings.Contains(user.Email, "@") {
		errorDetails = append(errorDetails, errs.ErrorEmailFormatT)
	}

	if user.Age < 18 {
		errorDetails = append(errorDetails, errs.ErrorAgeMinimum)
	}

	if len(errorDetails) > 0 {
		errMsg := errs.ErrorMessage{
			Error:   errs.ResponseValidationFailed,
			Details: errorDetails,
		}

		return &errMsg
	}

	return nil
}
