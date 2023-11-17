package validator

import (
	"github.com/Zoe-2Fu/ps-tag-onboarding-go/models"
	errs "github.com/Zoe-2Fu/ps-tag-onboarding-go/models/error"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/mock"
)

type UserValidatorMock struct {
	mock.Mock
}

func (m *UserValidatorMock) ValidateUserDetails(ctx echo.Context, user models.User) *errs.ErrorMessage {
	args := m.Called(ctx, user)
	return args.Get(0).(*errs.ErrorMessage)
}
