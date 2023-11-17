package mongo

import (
	"context"
	"net/http"

	"github.com/Zoe-2Fu/ps-tag-onboarding-go/models"
	errs "github.com/Zoe-2Fu/ps-tag-onboarding-go/models/error"
	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

var userCollection = "userdetails"

type UserRepository struct {
	db *mongo.Database
}

func NewUserRepo(db *mongo.Database) *UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) Find(ctx echo.Context, id string) (models.User, error) {
	var user models.User

	err := r.db.Collection(userCollection).FindOne(ctx.Request().Context(), bson.M{"id": id}).Decode(&user)
	if err != nil {
		return models.User{}, err
	}
	return user, nil
}

func (r *UserRepository) Save(ctx context.Context, user models.User) error {
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

func (r *UserRepository) ValidaiteUserExisted(user models.User) bool {
	filter := bson.M{"firstname": user.FirstName, "lastname": user.LastName}

	count, err := r.db.Collection(userCollection).CountDocuments(context.Background(), filter)
	if err != nil {
		return true
	}

	if count > 0 {
		return true
	}

	return false
}
