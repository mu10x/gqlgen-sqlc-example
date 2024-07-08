package pg

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"
)

type Repository interface {
	CreateAgent(ctx context.Context, arg CreateAgentParams) (Agent, error)
	DeleteAgent(ctx context.Context, id int64) (Agent, error)
	GetAgent(ctx context.Context, id int64) (Agent, error)
	ListAgents(ctx context.Context) ([]Agent, error)
	UpdateAgent(ctx context.Context, arg UpdateAgentParams) (Agent, error)

	CreateAuthor(ctx context.Context, arg CreateAuthorParams) (Author, error)
	DeleteAuthor(ctx context.Context, id int64) (Author, error)
	GetAuthor(ctx context.Context, id int64) (Author, error)
	ListAuthors(ctx context.Context) ([]Author, error)
	UpdateAuthor(ctx context.Context, arg UpdateAuthorParams) (Author, error)

	CreateBook(ctx context.Context, arg CreateBookParams, authorIDs []int64) (*Book, error)
	UpdateBook(ctx context.Context, arg UpdateBookParams, authorIDs []int64) (*Book, error)
	DeleteBook(ctx context.Context, id int64) (Book, error)
	GetBook(ctx context.Context, id int64) (Book, error)
	ListBooks(ctx context.Context) ([]Book, error)
}

type repoSvc struct {
	*Queries
	db *pgxpool.Pool
}

func (r *repoSvc) withTx(ctx context.Context, txFn func(*Queries) error) error {
	tx, err := r.db.Begin(context.Background())
	if err != nil {
		return err
	}
	q := New(tx)
	err = txFn(q)
	if err != nil {
		if rbErr := tx.Rollback(ctx); rbErr != nil {
			err = fmt.Errorf("tx failed: %v, unable to rollback: %v", err, rbErr)
		}
	} else {
		err = tx.Commit(ctx)
	}
	return err
}

func (r *repoSvc) CreateBook(ctx context.Context, arg CreateBookParams, authorIDs []int64) (*Book, error) {
	book := new(Book)
	err := r.withTx(ctx, func(q *Queries) error {
		res, err := q.CreateBook(ctx, arg)
		if err != nil {
			return err
		}
		for _, authorID := range authorIDs {
			if err := q.SetBookAuthor(ctx, SetBookAuthorParams{
				BookID:   res.ID,
				AuthorID: authorID,
			}); err != nil {
				return err
			}
		}
		book = &res
		return nil
	})
	return book, err
}

func (r *repoSvc) UpdateBook(ctx context.Context, arg UpdateBookParams, authorIDs []int64) (*Book, error) {
	book := new(Book)
	err := r.withTx(ctx, func(q *Queries) error {
		res, err := q.UpdateBook(ctx, arg)
		if err != nil {
			return err
		}
		if err = q.UnsetBookAuthors(ctx, res.ID); err != nil {
			return err
		}
		for _, authorID := range authorIDs {
			if err := q.SetBookAuthor(ctx, SetBookAuthorParams{
				BookID:   res.ID,
				AuthorID: authorID,
			}); err != nil {
				return err
			}
		}
		book = &res
		return nil
	})
	return book, err
}

func NewRepository(db *pgxpool.Pool) Repository {
	return &repoSvc{
		Queries: New(db),
		db:      db,
	}
}

func Open(dataSourceName string) (*pgxpool.Pool, error) {
	return pgxpool.New(context.Background(), dataSourceName)
}
