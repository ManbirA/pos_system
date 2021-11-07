package main

import (
	"net/http"
	"io/ioutil"
	"context"
	"time"
	"fmt"
	"github.com/julienschmidt/httprouter"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"github.com/ManbirA/pos_system/controllers"
)

func main() {
	
	r := httprouter.New()
	uc := controllers.NewUserController(getSession())
	pc := controllers.NewProductController(getSession())

	r.GET("/user/:id", uc.GetUser)
	r.POST("/user", uc.CreateUser)
	r.DELETE("/user/:id", uc.DeleteUser)

	r.GET("/product/:id", pc.GetProduct)
	r.GET("/products", pc.GetAllProducts)
	r.POST("/product", pc.CreateProduct)


	http.ListenAndServe("localhost: 8080", r);

}

func getSession() *mongo.Client {
	pass, err := ioutil.ReadFile("./creds.txt")
	if err != nil {
        fmt.Print(err)
    }
	mongoURI := fmt.Sprintf("mongodb+srv://ManbirA:%s@posdatacluster.4rn3t.mongodb.net/myFirstDatabase?retryWrites=true&w=majority", string(pass))

	clientOptions := options.Client().
		ApplyURI(mongoURI)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		panic(err)
	}
	return(client)
}
