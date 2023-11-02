package handler

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/Zoe-2Fu/ps-tag-onboarding-go/configs"
	"github.com/Zoe-2Fu/ps-tag-onboarding-go/model"
	validator "github.com/Zoe-2Fu/ps-tag-onboarding-go/validators"
	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

var userCollection *mongo.Collection = configs.GetCollection(configs.DB, "userdetails")

func Find(c echo.Context) error {
	id := c.Param("id")
	var user model.User

	err := userCollection.FindOne(c.Request().Context(), bson.M{"id": id}).Decode(&user)
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"message": "User not found"})
	}
	return c.JSON(http.StatusOK, user)
}

func Save(c echo.Context) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	user := new(model.User)
	if err := c.Bind(user); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": "Can't bind values"})
	}

	log.Printf("Received user data: %+v\n", user)

	validationErr := validator.ValidateUserDetails(c, *user, userCollection)
	if validationErr != nil {
		return echo.NewHTTPError(http.StatusBadRequest, validationErr)
	}

	userBSON, err := bson.Marshal(user)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Failed to marshaling BSON"})
	}

	_, err = userCollection.InsertOne(ctx, userBSON)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Failed to save user"})
	}

	return c.JSON(http.StatusCreated, user)
}
