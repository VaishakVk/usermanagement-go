package users

import (
	"context"
	"encoding/json"
	"goapi/lib/response"
	"goapi/models"
	"goapi/validator"
	"goapi/validator/schema"
	"net/http"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

// CreateUser endpoint
func CreateUser(res http.ResponseWriter, req *http.Request) {
	var user models.User
	var userCollection *mongo.Collection = user.GetCollection()

	parseErr := json.NewDecoder(req.Body).Decode(&user)
	if parseErr != nil {
		response.SendResponse(res, http.StatusBadRequest, false, parseErr.Error())
		return
	}
	var payload schema.CreateUserSchema = schema.CreateUserSchema(user)
	schemaErr := validator.Validate(payload)
	if schemaErr != nil {
		response.SendResponse(res, http.StatusBadRequest, false, schemaErr.Error())
		return
	}
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	inserted, err := userCollection.InsertOne(ctx, user)
	if err != nil {
		response.SendResponse(res, http.StatusInternalServerError, false, parseErr.Error())
		return
	}

	message := "User created successfully: " + inserted.InsertedID.(primitive.ObjectID).Hex()
	response.SendResponse(res, http.StatusCreated, true, message)
}

// GetUserByID endpoint
func GetUserByID(res http.ResponseWriter, req *http.Request) {
}
