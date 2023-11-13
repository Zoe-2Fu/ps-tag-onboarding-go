package configs

import (
	"context"

	"github.com/Zoe-2Fu/ps-tag-onboarding-go/model"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/mock"
)

type UserRepoMock struct {
	mock.Mock
}

func (m *UserRepoMock) Find(c echo.Context, id string) (model.User, error) {
	args := m.Called(c, id)
	return args.Get(0).(model.User), args.Error(1)
}

func (m *UserRepoMock) Save(c context.Context, user model.User) error {
	args := m.Called(c, user)
	return args.Error(0)
}

func (m *UserRepoMock) ValidaiteUserExisted(c echo.Context, user model.User) bool {
	args := m.Called(c, user)
	return args.Bool(0)
}
