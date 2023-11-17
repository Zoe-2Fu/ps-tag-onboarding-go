package validator

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/Zoe-2Fu/ps-tag-onboarding-go/models"
	errs "github.com/Zoe-2Fu/ps-tag-onboarding-go/models/error"
	"github.com/Zoe-2Fu/ps-tag-onboarding-go/mongo"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func SetUpContext() echo.Context {
	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, "/user", nil)
	rec := httptest.NewRecorder()
	return e.NewContext(req, rec)
}

func TestValidateUserDetails_ValidUserDetail(t *testing.T) {
	c := SetUpContext()
	user := models.NewUser("233333", "John", "Doe", "a@a.a", 20)

	userRepoMock := new(mongo.UserRepoMock)
	validator := &UserValidator{
		userRepo: userRepoMock,
	}

	expectedOutput := (*errs.ErrorMessage)(nil)

	userRepoMock.On("ValidaiteUserExisted", mock.Anything, mock.Anything).Return(false)

	result := validator.ValidateUserDetails(c, user)

	assert.Equal(t, expectedOutput, result)
}

func TestValidateUserDetails_UserIsExisted(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, "/user", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	user := models.NewUser("233333", "John", "Doe", "a@a.a", 20)

	userRepoMock := new(mongo.UserRepoMock)
	validator := &UserValidator{
		userRepo: userRepoMock,
	}

	expectedOutput := errs.NewErrorMessage(errs.ResponseValidationFailed, errs.ErrorNameUnique)
	expectedOutputPointer := &expectedOutput

	userRepoMock.On("ValidaiteUserExisted", mock.Anything, mock.Anything).Return(true)

	result := validator.ValidateUserDetails(c, user)

	assert.Equal(t, expectedOutputPointer, result)
}

func TestValidateUserDetails_UserNameIsMissing(t *testing.T) {
	c := SetUpContext()
	user := models.NewUser("233333", "", "Doe", "a@a.a", 20)

	userRepoMock := new(mongo.UserRepoMock)
	validator := &UserValidator{
		userRepo: userRepoMock,
	}

	expectedOutput := errs.NewErrorMessage(errs.ResponseValidationFailed, errs.ErrorNameRequired)
	expectedOutputPointer := &expectedOutput

	userRepoMock.On("ValidaiteUserExisted", mock.Anything, mock.Anything).Return(false)

	result := validator.ValidateUserDetails(c, user)

	assert.Equal(t, expectedOutputPointer, result)
}

func TestValidateUserDetails_UserEmailIsMissing(t *testing.T) {
	c := SetUpContext()
	user := models.NewUser("233333", "John", "Doe", "", 20)

	userRepoMock := new(mongo.UserRepoMock)
	validator := &UserValidator{
		userRepo: userRepoMock,
	}

	expectedOutput := errs.NewErrorMessage(errs.ResponseValidationFailed, errs.ErrorEmailRequired)
	expectedOutputPointer := &expectedOutput

	userRepoMock.On("ValidaiteUserExisted", mock.Anything, mock.Anything).Return(false)

	result := validator.ValidateUserDetails(c, user)

	assert.Equal(t, expectedOutputPointer, result)
}

func TestValidateUserDetails_InvaildUserEmailFormat(t *testing.T) {
	c := SetUpContext()
	user := models.NewUser("233333", "John", "Doe", "aa.a", 20)

	userRepoMock := new(mongo.UserRepoMock)
	validator := &UserValidator{
		userRepo: userRepoMock,
	}

	expectedOutput := errs.NewErrorMessage(errs.ResponseValidationFailed, errs.ErrorEmailFormat)
	expectedOutputPointer := &expectedOutput

	userRepoMock.On("ValidaiteUserExisted", mock.Anything, mock.Anything).Return(false)

	result := validator.ValidateUserDetails(c, user)

	assert.Equal(t, expectedOutputPointer, result)
}

func TestValidateUserDetails_InvaildUserAge(t *testing.T) {
	c := SetUpContext()
	user := models.NewUser("233333", "John", "Doe", "a@a.a", 16)

	userRepoMock := new(mongo.UserRepoMock)
	validator := &UserValidator{
		userRepo: userRepoMock,
	}

	expectedOutput := errs.NewErrorMessage(errs.ResponseValidationFailed, errs.ErrorAgeMinimum)
	expectedOutputPointer := &expectedOutput

	userRepoMock.On("ValidaiteUserExisted", mock.Anything, mock.Anything).Return(false)

	result := validator.ValidateUserDetails(c, user)

	assert.Equal(t, expectedOutputPointer, result)
}

func TestValidateUserDetails_MultipleUserDeatilErros(t *testing.T) {
	c := SetUpContext()
	user := models.NewUser("233333", "", "Doe", "aa.a", 20)

	userRepoMock := new(mongo.UserRepoMock)
	validator := &UserValidator{
		userRepo: userRepoMock,
	}

	expectedOutput := errs.ErrorMessage{
		Error:   errs.ResponseValidationFailed,
		Details: []string{errs.ErrorNameRequired, errs.ErrorEmailFormat},
	}
	expectedOutputPointer := &expectedOutput

	userRepoMock.On("ValidaiteUserExisted", mock.Anything, mock.Anything).Return(false)

	result := validator.ValidateUserDetails(c, user)

	assert.Equal(t, expectedOutputPointer, result)
}
