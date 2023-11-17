package handler

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/Zoe-2Fu/ps-tag-onboarding-go/models"
	errs "github.com/Zoe-2Fu/ps-tag-onboarding-go/models/error"
	"github.com/Zoe-2Fu/ps-tag-onboarding-go/mongo"
	validator "github.com/Zoe-2Fu/ps-tag-onboarding-go/validators"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestSave_StatusCreated(t *testing.T) {
	// arrange
	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, "/save", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	userRepoMock := new(mongo.UserRepoMock)
	validatorMock := new(validator.UserValidatorMock)

	userHandler := &UserHandler{
		userRepo:  userRepoMock,
		validator: validatorMock,
	}

	userRepoMock.On("Save", mock.Anything, mock.Anything).Return(nil)
	validatorMock.On("ValidateUserDetails", mock.Anything, mock.Anything).Return((*errs.ErrorMessage)(nil))

	// act
	err := userHandler.Save(c)

	// assert
	assert.NoError(t, err)
	assert.Equal(t, http.StatusCreated, rec.Code)
}

func TestFind_StatusOK(t *testing.T) {
	// arrange
	e := echo.New()

	userRepoMock := new(mongo.UserRepoMock)
	validatorMock := new(validator.UserValidatorMock)

	userHandler := &UserHandler{
		userRepo:  userRepoMock,
		validator: validatorMock,
	}

	userID := "123333"
	expectedUser := models.NewUser(userID, "John", "Doe", "a@a.a", 20)
	expectedReponseBody, _ := json.Marshal(expectedUser)

	userRepoMock.On("Find", mock.Anything, mock.Anything).Return(expectedUser, nil)
	validatorMock.On("ValidateUserDetails", mock.Anything, mock.Anything).Return((*errs.ErrorMessage)(nil))

	req := httptest.NewRequest(http.MethodPost, "/find/"+userID, nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	// act
	err := userHandler.Find(c)

	// assert
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, rec.Code)
	assert.JSONEq(t, string(expectedReponseBody), rec.Body.String())
}

func TestFind_StatusNotFound(t *testing.T) {
	// arrange
	e := echo.New()

	userRepoMock := new(mongo.UserRepoMock)
	validatorMock := new(validator.UserValidatorMock)

	userHandler := &UserHandler{
		userRepo:  userRepoMock,
		validator: validatorMock,
	}

	userID := "123"
	err := echo.NewHTTPError(http.StatusNotFound, errs.ErrorMessage{
		Error:   errs.ErrorStatusNotFound,
		Details: []string{"User not found"},
	})
	userRepoMock.On("Find", mock.Anything, mock.Anything).Return(models.User{}, err)

	req := httptest.NewRequest(http.MethodPost, "/find/"+userID, nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	// act
	handlerErr := userHandler.Find(c)

	// assert
	assert.Error(t, handlerErr)
}
