package gqlgen

// THIS CODE IS A STARTING POINT ONLY. IT WILL NOT BE UPDATED WITH SCHEMA CHANGES.

import (
	"context"

	"github.com/mu10x/gqlgen-sqlc-example/pg"
)

type Resolver struct {
	Repository pg.Repository
}

// Authors is the resolver for the authors field.
func (r *agentResolver) Authors(ctx context.Context, obj *pg.Agent) ([]pg.Author, error) {
	return r.Repository.ListAuthorsByAgentID(ctx, obj.ID)
}

// Website is the resolver for the website field.
func (r *authorResolver) Website(ctx context.Context, obj *pg.Author) (*string, error) {
	panic("not implemented")
}

// Agent is the resolver for the agent field.
func (r *authorResolver) Agent(ctx context.Context, obj *pg.Author) (*pg.Agent, error) {
	panic("not implemented")
}

// Books is the resolver for the books field.
func (r *authorResolver) Books(ctx context.Context, obj *pg.Author) ([]pg.Book, error) {
	panic("not implemented")
}

// Authors is the resolver for the authors field.
func (r *bookResolver) Authors(ctx context.Context, obj *pg.Book) ([]pg.Author, error) {
	panic("not implemented")
}

// CreateAgent is the resolver for the createAgent field.
func (r *mutationResolver) CreateAgent(ctx context.Context, data AgentInput) (*pg.Agent, error) {
	agent, err := r.Repository.CreateAgent(ctx, pg.CreateAgentParams{
		Name:  data.Name,
		Email: data.Email,
	})
	if err != nil {
		return nil, err
	}
	return &agent, nil
}

// UpdateAgent is the resolver for the updateAgent field.
func (r *mutationResolver) UpdateAgent(ctx context.Context, id int64, data AgentInput) (*pg.Agent, error) {
	panic("not implemented")
}

// DeleteAgent is the resolver for the deleteAgent field.
func (r *mutationResolver) DeleteAgent(ctx context.Context, id int64) (*pg.Agent, error) {
	panic("not implemented")
}

// CreateAuthor is the resolver for the createAuthor field.
func (r *mutationResolver) CreateAuthor(ctx context.Context, data AuthorInput) (*pg.Author, error) {
	panic("not implemented")
}

// UpdateAuthor is the resolver for the updateAuthor field.
func (r *mutationResolver) UpdateAuthor(ctx context.Context, id int64, data AuthorInput) (*pg.Author, error) {
	panic("not implemented")
}

// DeleteAuthor is the resolver for the deleteAuthor field.
func (r *mutationResolver) DeleteAuthor(ctx context.Context, id int64) (*pg.Author, error) {
	panic("not implemented")
}

// CreateBook is the resolver for the createBook field.
func (r *mutationResolver) CreateBook(ctx context.Context, data BookInput) (*pg.Book, error) {
	panic("not implemented")
}

// UpdateBook is the resolver for the updateBook field.
func (r *mutationResolver) UpdateBook(ctx context.Context, id int64, data BookInput) (*pg.Book, error) {
	panic("not implemented")
}

// DeleteBook is the resolver for the deleteBook field.
func (r *mutationResolver) DeleteBook(ctx context.Context, id int64) (*pg.Book, error) {
	panic("not implemented")
}

// Agent is the resolver for the agent field.
func (r *queryResolver) Agent(ctx context.Context, id int64) (*pg.Agent, error) {
	panic("not implemented")
}

// Agents is the resolver for the agents field.
func (r *queryResolver) Agents(ctx context.Context) ([]pg.Agent, error) {
	return r.Repository.ListAgents(ctx)
}

// Author is the resolver for the author field.
func (r *queryResolver) Author(ctx context.Context, id int64) (*pg.Author, error) {
	panic("not implemented")
}

// Authors is the resolver for the authors field.
func (r *queryResolver) Authors(ctx context.Context) ([]pg.Author, error) {
	panic("not implemented")
}

// Book is the resolver for the book field.
func (r *queryResolver) Book(ctx context.Context, id int64) (*pg.Book, error) {
	panic("not implemented")
}

// Books is the resolver for the books field.
func (r *queryResolver) Books(ctx context.Context) ([]pg.Book, error) {
	panic("not implemented")
}

// Agent returns AgentResolver implementation.
func (r *Resolver) Agent() AgentResolver { return &agentResolver{r} }

// Author returns AuthorResolver implementation.
func (r *Resolver) Author() AuthorResolver { return &authorResolver{r} }

// Book returns BookResolver implementation.
func (r *Resolver) Book() BookResolver { return &bookResolver{r} }

// Mutation returns MutationResolver implementation.
func (r *Resolver) Mutation() MutationResolver { return &mutationResolver{r} }

// Query returns QueryResolver implementation.
func (r *Resolver) Query() QueryResolver { return &queryResolver{r} }

type agentResolver struct{ *Resolver }
type authorResolver struct{ *Resolver }
type bookResolver struct{ *Resolver }
type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
