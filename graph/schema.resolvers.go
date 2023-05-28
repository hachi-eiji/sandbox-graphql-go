package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.31

import (
	"context"
	"http-go-sandbox/graph/model"
	"http-go-sandbox/sqls"
	"strconv"

	pgx "github.com/jackc/pgx/v5"
)

// Authors is the resolver for the authors field.
func (r *queryResolver) Authors(ctx context.Context) ([]*model.Author, error) {
	conn, err := pgx.Connect(context.Background(), databaseConn())

	if err != nil {
		return nil, err
	}

	defer conn.Close(ctx)

	db := sqls.New(conn)

	authors, err := db.ListAuthors(ctx)
	if err != nil {
		return nil, err
	}

	for _, a := range authors {
		author := &model.Author{ID: strconv.FormatInt(a.ID, 10), Name: a.Name}
		r.authors = append(r.authors, author)
	}

	return r.authors, nil
}

// Query returns QueryResolver implementation.
func (r *Resolver) Query() QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }

// !!! WARNING !!!
// The code below was going to be deleted when updating resolvers. It has been copied here so you have
// one last chance to move it out of harms way if you want. There are two reasons this happens:
//   - When renaming or deleting a resolver the old code will be put in here. You can safely delete
//     it when you're done.
//   - You have helper methods in this file. Move them out to keep these resolver files clean.
func databaseConn() string {
	return "postgres://postgres:example@127.0.0.1:5431/postgres"
}
