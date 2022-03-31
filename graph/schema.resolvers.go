package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"strconv"

	"github.com/mahfuz110244/graphql-go-demo/graph/generated"
	"github.com/mahfuz110244/graphql-go-demo/graph/model"
	"github.com/mahfuz110244/graphql-go-demo/repository"
)

func (r *mutationResolver) CreatAuthor(ctx context.Context, name string, biography string) (*model.Author, error) {
	var author model.Author
	author.Name = name
	author.Biography = biography
	id, err := repository.CreateAuthor(author)
	if err != nil {
		return nil, err
	} else {
		return &model.Author{ID: strconv.FormatInt(id, 10), Name: author.Name, Biography: author.Biography}, nil
	}
}

func (r *mutationResolver) CreateBook(ctx context.Context, title string, price int, isbnNo string, author string) (*model.Book, error) {
	var book model.Book
	book.Title = title
	book.Price = price
	book.IsbnNo = isbnNo
	book.Author = &model.Author{
		ID: author,
	}
	id, err := repository.CreateBook(book)
	if err != nil {
		return nil, err
	}
	idStr := strconv.Itoa(int(id))
	createdBook, _ := repository.GetBooksByID(&idStr)
	return createdBook, nil
}

func (r *queryResolver) BookByID(ctx context.Context, id *string) (*model.Book, error) {
	book, err := repository.GetBooksByID(id)
	if err != nil {
		return nil, err
	}
	return book, nil
}

func (r *queryResolver) AllBooks(ctx context.Context) ([]*model.Book, error) {
	books, err := repository.GetAllBooks()
	if err != nil {
		return nil, err
	}

	return books, nil
}

func (r *queryResolver) AuthorByID(ctx context.Context, id *string) (*model.Author, error) {
	author, err := repository.GetAuthorByID(id)
	if err != nil {
		return nil, err
	} else {
		return author, nil
	}
}

func (r *queryResolver) AllAuthors(ctx context.Context) ([]*model.Author, error) {
	authors, err := repository.GetAllAuthors()
	if err != nil {
		return nil, err
	} else {
		return authors, err
	}
}

func (r *queryResolver) Authors(ctx context.Context, name string) (*model.Books, error) {
	books, err := repository.GetAllBooksByAuthorName(name)
	if err != nil {
		return nil, err
	}
	booksData := &model.Books{
		Books: books,
	}
	return booksData, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
