package validator

import (
	"testing"

	errs "github.com/Zoe-2Fu/ps-tag-onboarding-go/internal/constants"
	models "github.com/Zoe-2Fu/ps-tag-onboarding-go/internal/data"
	repo "github.com/Zoe-2Fu/ps-tag-onboarding-go/internal/repository"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func TestValidateUserDetails_ValidUserDetails(t *testing.T) {
	user := models.NewUser(primitive.NilObjectID, "John", "Doe", "good@example.com", 25)

	userRepoMock := new(repo.UserRepoMock)
	validator := &UserValidator{
		userRepo: userRepoMock,
	}

	expectedOutput := (*errs.ErrorMessage)(nil)

	userRepoMock.On("ValidateUserExisted", mock.Anything, mock.Anything).Return(false)

	result := validator.ValidateUserDetails(user)

	assert.Nil(t, result, expectedOutput)
}

func TestValidateUserDetails_UserIsExisted(t *testing.T) {
	user := models.NewUser(primitive.NilObjectID, "John", "Doe", "a@a.a", 20)

	userRepoMock := new(repo.UserRepoMock)
	validator := &UserValidator{
		userRepo: userRepoMock,
	}

	expectedOutput := errs.NewErrorMessage(errs.ResponseValidationFailed, errs.ErrorNameUnique, user.ToWithouUserID())
	expectedOutputPointer := &expectedOutput

	userRepoMock.On("ValidateUserExisted", mock.Anything, mock.Anything).Return(true)

	result := validator.ValidateUserDetails(user)

	assert.Equal(t, expectedOutputPointer, result)
}

func TestValidateUserDetails_UserNameIsMissing(t *testing.T) {
	user := models.NewUser(primitive.NilObjectID, "", "Doe", "a@a.a", 20)

	userRepoMock := new(repo.UserRepoMock)
	validator := &UserValidator{
		userRepo: userRepoMock,
	}

	expectedOutput := errs.NewErrorMessage(errs.ResponseValidationFailed, errs.ErrorNameRequired, user.ToWithouUserID())
	expectedOutputPointer := &expectedOutput

	userRepoMock.On("ValidaiteUserExisted", mock.Anything, mock.Anything).Return(false)

	result := validator.ValidateUserDetails(user)

	assert.Equal(t, expectedOutputPointer, result)
}

func TestValidateUserDetails_UserEmailIsMissing(t *testing.T) {
	user := models.NewUser(primitive.NewObjectID(), "John", "Doe", "", 20)

	userRepoMock := new(repo.UserRepoMock)
	validator := &UserValidator{
		userRepo: userRepoMock,
	}

	expectedOutput := errs.NewErrorMessage(errs.ResponseValidationFailed, errs.ErrorEmailRequired, user.ToWithouUserID())
	expectedOutputPointer := &expectedOutput

	userRepoMock.On("ValidateUserExisted", mock.Anything, mock.Anything).Return(false)

	result := validator.ValidateUserDetails(user)

	assert.Equal(t, expectedOutputPointer, result)
}

func TestValidateUserDetails_InvalidUserEmailFormat(t *testing.T) {
	user := models.NewUser(primitive.NewObjectID(), "John", "Doe", "aa.a", 20)

	userRepoMock := new(repo.UserRepoMock)
	validator := &UserValidator{
		userRepo: userRepoMock,
	}

	expectedOutput := errs.NewErrorMessage(errs.ResponseValidationFailed, errs.ErrorEmailFormat, user.ToWithouUserID())
	expectedOutputPointer := &expectedOutput

	userRepoMock.On("ValidateUserExisted", mock.Anything, mock.Anything).Return(false)

	result := validator.ValidateUserDetails(user)

	assert.Equal(t, expectedOutputPointer, result)
}

func TestValidateUserDetails_InvalidUserAge(t *testing.T) {
	user := models.NewUser(primitive.NilObjectID, "John", "Doe", "a@a.a", 16)

	userRepoMock := new(repo.UserRepoMock)
	validator := &UserValidator{
		userRepo: userRepoMock,
	}

	expectedOutput := errs.NewErrorMessage(errs.ResponseValidationFailed, errs.ErrorAgeMinimum, user.ToWithouUserID())
	expectedOutputPointer := &expectedOutput

	userRepoMock.On("ValidateUserExisted", mock.Anything, mock.Anything).Return(false)

	result := validator.ValidateUserDetails(user)

	assert.Equal(t, expectedOutputPointer, result)
}

func TestValidateUserDetails_MultipleUserDetailsErrors(t *testing.T) {
	user := models.NewUser(primitive.NilObjectID, "", "Doe", "aa.a", 20)

	userRepoMock := new(repo.UserRepoMock)
	validator := &UserValidator{
		userRepo: userRepoMock,
	}

	expectedUser := user.ToWithouUserID()
	expectedOutput := errs.ErrorMessage{
		Error:   errs.ResponseValidationFailed,
		Details: []string{errs.ErrorNameRequired, errs.ErrorEmailFormat},
		User:    &expectedUser,
	}
	expectedOutputPointer := &expectedOutput

	userRepoMock.On("ValidateUserExisted", mock.Anything, mock.Anything).Return(false)

	result := validator.ValidateUserDetails(user)

	assert.Equal(t, expectedOutputPointer, result)
}
