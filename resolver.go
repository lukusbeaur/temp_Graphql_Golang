package ProjectQ_Reg_Login

import (
	"context"

	"github.com/projects/ProjectQ_Reg_Login/api"
	"github.com/projects/ProjectQ_Reg_Login/api/dal"
) // THIS CODE IS A STARTING POINT ONLY. IT WILL NOT BE UPDATED WITH SCHEMA CHANGES.

type Resolver struct{}

func (r *Resolver) Mutation() MutationResolver {
	return &mutationResolver{r}
}
func (r *Resolver) Query() QueryResolver {
	return &queryResolver{r}
}

type mutationResolver struct{ *Resolver }

func (r *mutationResolver) CreateUser(ctx context.Context, input NewUser) (api.NewUser, error) {
	newUser := api.NewUser{
		Name:     input.Name,
		Username: input.Username,
		Email:    input.Email,
		Password: input.Password,
	}

	rows, err := dal.LogAndQuery(r.db, "INSERT INTO Users (name, username, email, password) VALUES($1, $2, $3, $4) RETURNING id",
		input.Name, input.Username, input.Email, input.Password)

	defer rows.Close()

	if err != nil || !rows.Next() {
		return api.User{}, err
	}
	if err := rows.Scan(&newUser.ID); err != nil {

	}
}

type queryResolver struct{ *Resolver }

func (r *queryResolver) Users(ctx context.Context, limit *int, offset *int) ([]User, error) {
	panic("not implemented")
}
