package service

import (
	"context"

	"github.com/kunlanat/go-example/api/v1/books/model"
	"github.com/kunlanat/go-example/domain"
	"github.com/kunlanat/go-example/dto"
	"github.com/kunlanat/go-example/errs"
	"github.com/kunlanat/go-example/repository"
)

type BookService interface {
	GetAllBooks(ctx context.Context, query *dto.QuerySearch) ([]domain.Books, error)
	GetBooksById(ctx context.Context, id string) (*domain.Books, error)
	CreateBook(ctx context.Context, book model.DTOCreateBook) (*domain.Books, error)
	UpdateBookById(ctx context.Context, id string, book *model.DTOUpdateBook) (*domain.Books, error)
	DeleteBookByID(ctx context.Context, id string) error
}

func BookServiceImp(repo repository.BookRepository) BookService {
	return &bookServiceImp{repo: repo}
}

type bookServiceImp struct {
	repo repository.BookRepository
}

func (s *bookServiceImp) GetAllBooks(ctx context.Context, query *dto.QuerySearch) ([]domain.Books, error) {
	return s.repo.GetAllBooks(ctx, query.Search)
}

func (s *bookServiceImp) GetBooksById(ctx context.Context, id string) (*domain.Books, error) {
	if len(id) == 0 {
		return nil, errs.ErrIDRequired
	}
	return s.repo.GetBooksById(ctx, id)
}

func (s *bookServiceImp) CreateBook(ctx context.Context, book model.DTOCreateBook) (*domain.Books, error) {
	res, err := s.repo.CreateBook(ctx, &book)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (s *bookServiceImp) UpdateBookById(ctx context.Context, id string, book *model.DTOUpdateBook) (*domain.Books, error) {
	if len(id) == 0 {
		return nil, errs.ErrIDRequired
	}

	data, err := s.repo.GetBooksById(ctx, id)
	if err != nil {
		return nil, err
	}

	return s.repo.UpdateBookById(ctx, id, &domain.Books{
		Name:  data.Name,
		Desc:  data.Desc,
		Price: book.Price,
	})
}

func (s *bookServiceImp) DeleteBookByID(ctx context.Context, id string) error {
	if len(id) == 0 {
		return errs.ErrIDRequired
	}
	return s.repo.DeleteBooksById(ctx, id)
}
