package controllers

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
	"log"

	"github.com/ManbirA/pos_system/models"
	"github.com/julienschmidt/httprouter"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)
type UserController struct {
	session *mongo.Client
}

func NewUserController(s *mongo.Client) *UserController{
	return &UserController{s}
}

func (uc UserController) GetUser (w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	id := p.ByName("id")
	ObjectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		log.Fatal(err)
	}
	u := models.User{}
	collection := uc.session.Database("pos_system").Collection("User")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	result := collection.FindOne(ctx, bson.M{"_id": ObjectID})
	result.Decode(&u)
	rtrUser, errUser := json.Marshal(&u);

	if errUser != nil {
		log.Fatal(err)
	}
	w.Header().Set("Content-Type", "applications/json")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "%s\n", rtrUser)
}

func (uc UserController) CreateUser(w http.ResponseWriter, r *http.Request, _ httprouter.Params){
	u := models.NewUser{}
	json.NewDecoder(r.Body).Decode(&u);

	collection := uc.session.Database("pos_system").Collection("User")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	res, err := collection.InsertOne(ctx, u)
	if err != nil {
		log.Fatal(err)
	}

	resUsr := &models.User{
		ID: res.InsertedID.(primitive.ObjectID).Hex(),
		Name: u.Name,
		Gender: u.Gender,
		Age: u.Age,
	}
	rtrUser, errUser := json.Marshal(resUsr);

	if errUser != nil {
		log.Fatal(err)
	}
	w.Header().Set("Content-Type", "applications/json")
	w.WriteHeader(http.StatusCreated)
	fmt.Fprintf(w, "%s\n", rtrUser)

}

func (uc UserController) DeleteUser(w http.ResponseWriter, r *http.Request, p httprouter.Params){
	id := p.ByName("id")
	ObjectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		log.Fatal(err)
	}

	collection := uc.session.Database("pos_system").Collection("User")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	
	res, resErr := collection.DeleteOne(ctx, bson.M{"_id": ObjectID})
	if resErr != nil {
		w.WriteHeader(http.StatusNotFound)
	}

	if (*res).DeletedCount > 0 {
		w.WriteHeader(http.StatusOK)

		fmt.Fprintf(w, "Deleted User: %s\n", id)
	} else {
		w.WriteHeader(http.StatusNotFound)
	}

}