package validator

import (
	"strings"

	"github.com/Zoe-2Fu/ps-tag-onboarding-go/model"
	errs "github.com/Zoe-2Fu/ps-tag-onboarding-go/model/error"
	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func ValidateUserDetails(c echo.Context, user model.User, userCollection *mongo.Collection) *errs.ErrorMessage {
	var errorDetails []string

	// check if the user exist
	existingUser := model.User{}
	filter := bson.M{"firstname": user.FirstName}

	err := userCollection.FindOne(c.Request().Context(), filter).Decode(&existingUser)
	if err == nil {
		errorDetails = append(errorDetails, errs.ErrorNameUnique)
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
