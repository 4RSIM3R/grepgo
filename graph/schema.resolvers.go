package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"strconv"

	"github.com/ilzam/hackernews/graph/generated"
	"github.com/ilzam/hackernews/graph/model"
	"github.com/ilzam/hackernews/model/dawuhs"
)

func (r *mutationResolver) CreateDawuh(ctx context.Context, input model.NewDawuh) (*model.Dawuh, error) {
	// var dawuh model.Dawuh
	// var author model.User
	// dawuh.Dawuh = input.Dawuh
	// dawuh.Qoil = input.Qoil
	// author.Name = "Cak Suku"
	// dawuh.Author = &author
	// return &dawuh, nil
	var dawuh dawuhs.Dawuh
	dawuh.Dawuh = input.Dawuh
	dawuh.Qoil = input.Qoil
	dawuhID := dawuh.Save()
	return &model.Dawuh{ID: strconv.FormatInt(dawuhID, 10), Dawuh: dawuh.Dawuh, Qoil: dawuh.Qoil}, nil
}

func (r *mutationResolver) CreateUser(ctx context.Context, input model.NewUser) (string, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) Login(ctx context.Context, input model.Login) (string, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) RefreshToken(ctx context.Context, input model.RefreshTokenInput) (string, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Dawuh(ctx context.Context) ([]*model.Dawuh, error) {
	// var dawuh dawuhs.Dawuh
	// var dawuhs []*model.Dawuh
	var res []*model.Dawuh
	var re []dawuhs.Dawuh
	// dawuhs = append(dawuhs, &model.Dawuh{Dawuh: "Meneng, Ngalah, Menang", Qoil: "KH Dimyathi Romly", Author: &model.User{Name: "Cak Suku"}})
	// dawuhs = append(dawuhs, &model.Dawuh{Dawuh: "Berdikir Kuat, Berfikir Cepat, Bertindak Tepat", Qoil: "KH As'ad Umar", Author: &model.User{Name: "Cak Suku"}})
	// return dawuhs, nil
	re = dawuhs.GetAll()
	for _, reslt := range re {
		res = append(res, &model.Dawuh{Dawuh: reslt.Dawuh, Qoil: reslt.Qoil, ID: reslt.ID})
	}
	return res, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
