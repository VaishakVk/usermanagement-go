package users

import (
	"context"
	"encoding/json"
	"goapi/helpers"
	"goapi/lib/response"
	"goapi/models"
	"goapi/validator"
	"goapi/validator/schema"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// CreateUser endpoint
func CreateUser(res http.ResponseWriter, req *http.Request) {
	var user models.User
	var userCollection *mongo.Collection = user.GetCollection()

	// Parse Request body to user
	parseErr := json.NewDecoder(req.Body).Decode(&user)
	if parseErr != nil {
		response.SendResponse(res, http.StatusBadRequest, false, parseErr.Error())
		return
	}

	// Validate user
	var payload schema.CreateUserSchema = schema.CreateUserSchema(user)
	schemaErr := validator.Validate(payload)
	if schemaErr != nil {
		response.SendResponse(res, http.StatusBadRequest, false, schemaErr.Error())
		return
	}

	hash, err := helpers.HashPassword(user.Password)
	if err != nil {
		response.SendResponse(res, http.StatusInternalServerError, false, schemaErr.Error())
		return
	}
	user.Password = string(hash)

	// Insert into Mongo
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
	params := mux.Vars(req)
	id := params["id"]
	var user models.User
	var userCollection *mongo.Collection = user.GetCollection()
	projection := bson.D{
		{"first_name", 1}, {"email", 1},
	}
	docID, _ := primitive.ObjectIDFromHex(id)
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	if err := userCollection.FindOne(ctx,
		bson.M{"_id": docID},
		options.FindOne().SetProjection(projection)).Decode(&user); err != nil {
		response.SendResponse(res, http.StatusInternalServerError, false, err.Error())
	}
	response.SendResponse(res, http.StatusOK, true, user)
}

// GetAllUsers endpoint
func GetAllUsers(res http.ResponseWriter, req *http.Request) {
	var user models.User
	userCollection := user.GetCollection()
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	projection := bson.D{
		{"first_name", 1}, {"email", 1},
	}
	cursor, dbError := userCollection.Find(ctx, bson.M{}, options.Find().SetProjection(projection))
	if dbError != nil {
		response.SendResponse(res, http.StatusInternalServerError, false, dbError.Error())
	}
	var users []models.User
	if err := cursor.All(ctx, &users); err != nil {
		response.SendResponse(res, http.StatusInternalServerError, false, err.Error())

	}
	response.SendResponse(res, http.StatusOK, true, users)
}

// UpdateUser Endpoint
func UpdateUser(res http.ResponseWriter, req *http.Request) {
	var user, out models.User
	var updateData map[string]interface{}
	userCollection := user.GetCollection()
	params := mux.Vars(req)
	id := params["id"]
	docID, _ := primitive.ObjectIDFromHex(id)
	json.NewDecoder(req.Body).Decode(&updateData)
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	filter := bson.M{"_id": docID}
	update := bson.M{"$set": updateData}
	projection := bson.D{
		{"first_name", 1}, {"email", 1},
	}
	optionsUpdate := options.FindOneAndUpdate().SetReturnDocument(1).SetProjection(projection)
	result := userCollection.FindOneAndUpdate(ctx, filter, update, optionsUpdate).Decode(&out)
	if result != nil {
		response.SendResponse(res, http.StatusNotFound, false, "User does not exist")
	} else {
		response.SendResponse(res, http.StatusOK, true, out)
	}
}

func LoginUser(res http.ResponseWriter, req *http.Request) {
	var user, reqBody models.User
	userCollection := user.GetCollection()
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	parseErr := json.NewDecoder(req.Body).Decode(&reqBody)
	if parseErr != nil {
		response.SendResponse(res, http.StatusBadRequest, false, parseErr.Error())
	}

	userCollection.FindOne(ctx, bson.M{"email": reqBody.Email}).Decode(&user)
	err := helpers.CompareHash(user.Password, reqBody.Password)
	if err != nil {
		response.SendResponse(res, http.StatusBadRequest, false, "Passwords do not match")
		return
	}
	claims := jwt.MapClaims{}
	claims["email"] = reqBody.Email
	token, jwtErr := helpers.SignToken(claims)
	if jwtErr != nil {
		response.SendResponse(res, http.StatusInternalServerError, false, jwtErr.Error())
		return
	}
	response.SendResponse(res, http.StatusOK, true, token)
}

func GetMe(res http.ResponseWriter, req *http.Request) {
	token := req.Header.Get("token")
	if len(token) == 0 {
		response.SendResponse(res, http.StatusForbidden, false, "Token is missing")
		return
	}
	tokenDecoded, err := helpers.ValidateToken(token)
	if err != nil {
		response.SendResponse(res, http.StatusInternalServerError, false, err.Error())
		return
	}
	response.SendResponse(res, http.StatusOK, true, tokenDecoded)
}
