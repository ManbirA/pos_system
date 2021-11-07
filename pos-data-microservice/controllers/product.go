package controllers

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/ManbirA/pos_system/models"
	"github.com/julienschmidt/httprouter"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type ProductController struct {
	session *mongo.Client
}

func NewProductController(s *mongo.Client) *ProductController{
	return &ProductController{s}
}

func (uc ProductController) GetProduct (w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	// Parse ID
	id := p.ByName("id")
	ObjectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		log.Fatal(err)
		w.WriteHeader(http.StatusBadRequest)
	}

	// Make mongoDB query
	product := models.Product{}
	collection := uc.session.Database("pos_system").Collection("products")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	findErr := collection.FindOne(ctx, bson.M{"_id": ObjectID}).Decode(&product)

	if findErr != nil {
		log.Fatal(findErr)
		w.WriteHeader(http.StatusNotFound)
	}
	// decode result
	productJson, errProductJson := json.Marshal(&product);
	if errProductJson != nil {
		log.Fatal(errProductJson)
		w.WriteHeader(http.StatusBadRequest)
	}

	// write response
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "%s\n", productJson)
}

func (uc ProductController) GetAllProducts (w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	// Make mongoDB query
	collection := uc.session.Database("pos_system").Collection("products")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	results, findErr := collection.Find(ctx, bson.D{})
	if findErr != nil {
		log.Fatal(findErr)
		w.WriteHeader(http.StatusNotFound)
	}

	// decode result
	products := []*models.Product{}
	for results.Next(ctx) {
		var product *models.Product
		err := results.Decode(&product)
		if err != nil {
			log.Fatal(err)
			w.WriteHeader(http.StatusBadRequest)
		} 
		products = append(products, product)
	}

	productsJson, errProductsJson := json.Marshal(&products);
	if errProductsJson != nil {
		log.Fatal(errProductsJson)
		w.WriteHeader(http.StatusBadRequest)

	}

	// write response
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "%s\n", productsJson)
}

func (uc ProductController) CreateProduct(w http.ResponseWriter, r *http.Request, _ httprouter.Params){
	// Get product parameters
	product := models.NewProduct{}
	json.NewDecoder(r.Body).Decode(&product);

	// Make databse mutation
	collection := uc.session.Database("pos_system").Collection("products")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	res, err := collection.InsertOne(ctx, product)
	if err != nil {
		log.Fatal(err)
		w.WriteHeader(http.StatusNotFound)
	}

	resProduct := &models.Product{
		ID: res.InsertedID.(primitive.ObjectID).Hex(),
		Name: product.Name,
		Inventory: product.Inventory,
		Price: product.Price,
	}
	rtrProduct, errProduct := json.Marshal(resProduct);

	if errProduct != nil {
		log.Fatal(err)
		w.WriteHeader(http.StatusBadRequest)
	}
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.WriteHeader(http.StatusCreated)
	fmt.Fprintf(w, "%s\n", rtrProduct)

}
