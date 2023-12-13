package constants

import models "github.com/Zoe-2Fu/ps-tag-onboarding-go/internal/data"

var (
	ErrorAgeMinimum          = "User does not meet minimum age requirement"
	ErrorEmailFormat         = "User email must be properly formatted"
	ErrorEmailRequired       = "User email required"
	ErrorNameRequired        = "User first/last names required"
	ErrorNameUnique          = "User with the same first and last name already exists"
	ResponseUserNotFound     = "User not found"
	ResponseValidationFailed = "User did not pass validation"
	ErrorBadRequest          = "Bad Request"
	ErrorStatusNotFound      = "Status Not Found"
	ErrorInternalServerError = "Internal Server Error"
)

type ErrorMessage struct {
	Error   string      `json:"error"`
	Details []string    `json:"details"`
	User    models.User `json:"userdetails"`
}

func NewErrorMessage(errMsg string, detail string, user models.User) ErrorMessage {
	details := []string{detail}
	return ErrorMessage{errMsg, details, user}
}
