package users

import (
	"context"
	"encoding/json"
	"goapi/lib/response"
	"goapi/models"
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
		res.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(res).Encode(response.APIResponse{Status: false, Response: parseErr.Error()})
	}
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	inserted, err := userCollection.InsertOne(ctx, user)
	if err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(res).Encode(response.APIResponse{Status: false, Response: err.Error()})
	}
	// fmt.Println(inserted.InsertedID)
	res.WriteHeader(http.StatusCreated)
	message := "User created successfully: " + inserted.InsertedID.(primitive.ObjectID).Hex()
	json.NewEncoder(res).Encode(response.APIResponse{Status: true, Response: message})
}

// GetUserByID endpoint
func GetUserByID(res http.ResponseWriter, req *http.Request) {
}
