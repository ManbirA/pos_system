package database

import (
	"context"
	"log"
	"time"

	"github.com/ManbirA/pos-data-microservice/graph/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type DB struct {
	client *mongo.Client
}

func Connect() *DB {
	clientOptions := options.Client().
		ApplyURI("mongodb+srv://ManbirA:<password>@posdatacluster.4rn3t.mongodb.net/myFirstDatabase?retryWrites=true&w=majority")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal(err)
	}
	return &DB{
		client: client,
	}
}

func (db *DB) Save(input *model.NewProduct) *model.Product {
	collection := db.client.Database("pos_system").Collection("products")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	res, err := collection.InsertOne(ctx, input)
	if err != nil {
		log.Fatal(err)
	}
	return &model.Product{
		ID:        res.InsertedID.(primitive.ObjectID).Hex(),
		Name:      input.Name,
		Price:     input.Price,
		Inventory: input.Inventory,
	}
}

func (db *DB) FindByID(ID string) *model.Product {
	ObjectID, err := primitive.ObjectIDFromHex(ID)
	collection := db.client.Database("pos_system").Collection("products")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	res := collection.FindOne(ctx, bson.M{"_id": ObjectID})
	if err != nil {
		log.Fatal(err)
	}
	product := model.Product{}
	res.Decode(&product)
	return &product
}

func (db *DB) All() []*model.Product {
	collection := db.client.Database("pos_system").Collection("products")
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	cur, err := collection.Find(ctx, bson.D{})
	if err != nil {
		log.Fatal(err)
	}
	var products []*model.Product
	for cur.Next(ctx) {
		var product *model.Product
		err := cur.Decode(&product)
		if err != nil {
			log.Fatal(err)
		}
		products = append(products, product)
	}
	return products

}
