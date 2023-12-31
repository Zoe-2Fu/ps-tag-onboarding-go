package repo

import (
	"context"
	"net/http"

	errs "github.com/Zoe-2Fu/ps-tag-onboarding-go/internal/constants"
	models "github.com/Zoe-2Fu/ps-tag-onboarding-go/internal/data"
	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
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

	objID, _ := primitive.ObjectIDFromHex(id)
	err := r.db.Collection(userCollection).FindOne(ctx.Request().Context(), bson.M{"_id": objID}).Decode(&user)
	if err != nil {
		return models.User{}, err
	}
	return user, nil
}

func (r *UserRepository) Save(ctx context.Context, user models.User) (primitive.ObjectID, error) {
	userBSON, err := bson.Marshal(user)
	if err != nil {
		return primitive.NilObjectID, echo.NewHTTPError(http.StatusInternalServerError, errs.ErrorMessage{
			Error:   errs.ErrorInternalServerError,
			Details: []string{"Failed to marshaling BSON"},
		})
	}
	result, err := r.db.Collection(userCollection).InsertOne(ctx, userBSON)
	if err != nil {
		return primitive.NilObjectID, echo.NewHTTPError(http.StatusInternalServerError, errs.ErrorMessage{
			Error:   errs.ErrorInternalServerError,
			Details: []string{"Failed to save user"},
		})
	}

	insertedID, ok := result.InsertedID.(primitive.ObjectID)
	if !ok {
		return primitive.NilObjectID, echo.NewHTTPError(http.StatusInternalServerError, errs.ErrorMessage{
			Error:   errs.ErrorInternalServerError,
			Details: []string{"Failed to find objectID"},
		})
	}

	return insertedID, err
}

func (r *UserRepository) ValidateUserExisted(user models.User) bool {
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
