package configs

import (
	"context"
	"net/http"

	"github.com/Zoe-2Fu/ps-tag-onboarding-go/model"
	errs "github.com/Zoe-2Fu/ps-tag-onboarding-go/model/error"
	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

var userCollection = "user"

type UserRepository struct {
	db *mongo.Database
}

func NewUserRepo(db *mongo.Database) *UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) Find(ctx echo.Context, id string) (model.User, error) {
	var user model.User

	err := r.db.Collection(userCollection).FindOne(ctx.Request().Context(), bson.M{"id": id}).Decode(&user)
	if err != nil {
		return model.User{}, err
	}
	return user, nil
}

func (r *UserRepository) Save(ctx context.Context, user model.User) error {
	userBSON, err := bson.Marshal(user)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, errs.ErrorMessage{
			Error:   errs.ErrorInternalServerError,
			Details: []string{"Failed to marshaling BSON"},
		})
	}

	_, err = r.db.Collection(userCollection).InsertOne(ctx, userBSON)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, errs.ErrorMessage{
			Error:   errs.ErrorInternalServerError,
			Details: []string{"Failed to save user"},
		})
	}

	return nil
}

func (r *UserRepository) ValidaiteUserExisted(ctx echo.Context, user model.User) bool {
	existingUser := model.User{}
	filter := bson.M{"firstname": user.FirstName, "lastname": user.LastName}

	err := r.db.Collection(userCollection).FindOne(ctx.Request().Context(), filter).Decode(&existingUser)
	if err != nil {
		return true
	}

	return false
}
