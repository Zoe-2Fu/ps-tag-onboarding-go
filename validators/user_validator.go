package validator

import (
	"strings"

	"github.com/Zoe-2Fu/ps-tag-onboarding-go/models"
	errs "github.com/Zoe-2Fu/ps-tag-onboarding-go/models/error"
)

type userRepo interface {
	ValidaiteUserExisted(user models.User) bool
}

type UserValidator struct {
	userRepo userRepo
}

func NewUserValidator(repo userRepo) *UserValidator {
	return &UserValidator{userRepo: repo}
}

func (v *UserValidator) ValidateUserDetails(user models.User) *errs.ErrorMessage {
	var errorDetails []string

	if len(user.FirstName) == 0 || len(user.LastName) == 0 {
		errorDetails = append(errorDetails, errs.ErrorNameRequired)
	} else {
		isExist := v.userRepo.ValidaiteUserExisted(user)
		if isExist {
			errMsg := errs.NewErrorMessage(errs.ResponseValidationFailed, errs.ErrorNameUnique)

			return &errMsg
		}
	}

	if len(user.Email) == 0 {
		errorDetails = append(errorDetails, errs.ErrorEmailRequired)
	} else if !strings.Contains(user.Email, "@") {
		errorDetails = append(errorDetails, errs.ErrorEmailFormat)
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
