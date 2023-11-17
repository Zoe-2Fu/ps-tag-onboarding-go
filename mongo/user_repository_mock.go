package mongo

import (
	"context"

	"github.com/Zoe-2Fu/ps-tag-onboarding-go/models"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/mock"
)

type UserRepoMock struct {
	mock.Mock
}

func (m *UserRepoMock) Find(c echo.Context, id string) (models.User, error) {
	args := m.Called(c, id)
	return args.Get(0).(models.User), args.Error(1)
}

func (m *UserRepoMock) Save(c context.Context, user models.User) error {
	args := m.Called(c, user)
	return args.Error(0)
}

func (m *UserRepoMock) ValidaiteUserExisted(user models.User) bool {
	args := m.Called(user)
	return args.Bool(0)
}
