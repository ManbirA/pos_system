package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/ManbirA/pos-data-microservice/database"
	"github.com/ManbirA/pos-data-microservice/graph/generated"
	"github.com/ManbirA/pos-data-microservice/graph/model"
)

var db = database.Connect()

func (r *mutationResolver) CreateProduct(ctx context.Context, input *model.NewProduct) (*model.Product, error) {
	return db.Save(input), nil
}

func (r *queryResolver) Product(ctx context.Context, id string) (*model.Product, error) {
	return db.FindByID(id), nil
}

func (r *queryResolver) Products(ctx context.Context) ([]*model.Product, error) {
	return db.All(), nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
