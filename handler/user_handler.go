package handler

import (
	"context"
	"github.com/Zoe-2Fu/ps-tag-onboarding-go/configs"
	"github.com/Zoe-2Fu/ps-tag-onboarding-go/model"
	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"net/http"
	"time"
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
	var user model.User
	defer cancel()

	if err := c.Bind(&user); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": "Error"})
	}

	newUser := model.User{
		ID:        primitive.NewObjectID(),
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Email:     user.Email,
		Age:       user.Age,
	}

	result, err := userCollection.InsertOne(ctx, newUser)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Failed to save user"})
	}

	return c.JSON(http.StatusCreated, result.InsertedID)
}
