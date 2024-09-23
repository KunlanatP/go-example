package repository

import (
	"context"

	"github.com/kunlanat/go-example/api/v1/books/model"
	"github.com/kunlanat/go-example/domain"
	"github.com/kunlanat/go-example/repository/entities"
	"gorm.io/gorm"
)

type BookRepository interface {
	CreateBook(ctx context.Context, book *model.DTOCreateBook) (*domain.Books, error)
	GetAllBooks(ctx context.Context, search string) ([]domain.Books, error)
	GetBooksById(ctx context.Context, id string) (*domain.Books, error)
	UpdateBookById(ctx context.Context, id string, book *domain.Books) (*domain.Books, error)
	DeleteBooksById(ctx context.Context, id string) error
}

func BookRepositoryWithGORM(db *gorm.DB) BookRepository {
	return &bookRepo{
		db: db,
	}
}

type bookRepo struct {
	db *gorm.DB
}

func (r *bookRepo) CreateBook(ctx context.Context, book *model.DTOCreateBook) (*domain.Books, error) {
	created := &entities.Books{
		Name:  book.Name,
		Desc:  book.Desc,
		Price: book.Price,
	}
	result := r.db.Create(created)

	if err := result.Error; err != nil {
		return nil, err
	}
	return created.ToDomain(), nil
}

func (r *bookRepo) GetAllBooks(ctx context.Context, search string) (out []domain.Books, err error) {
	books := []entities.Books{}
	qry := r.db.Model(&entities.Books{})
	if len(search) > 0 {
		qry.Where("name LIKE ?", "%"+search+"%")
	}

	if err := qry.Find(&books).Error; err != nil {
		return nil, err
	}

	for _, book := range books {
		out = append(out, *book.ToDomain())
	}

	return
}

func (r *bookRepo) GetBooksById(ctx context.Context, id string) (*domain.Books, error) {
	book := entities.Books{}
	qry := r.db.Model(&entities.Books{}).Where("id = ?", id).First(&book)

	if err := qry.Error; err != nil {
		return nil, err
	}

	return book.ToDomain(), nil
}

func (r *bookRepo) UpdateBookById(ctx context.Context, id string, book *domain.Books) (*domain.Books, error) {
	tx := r.db.Begin()
	update := entities.Books{
		Name:  book.Name,
		Desc:  book.Desc,
		Price: book.Price,
	}

	updated := r.db.Model(&update).Where("id = ?", id).Updates(update)

	if err := updated.Error; err != nil {
		return nil, err
	}

	if err := tx.First(&update, "id=?", id).Error; err != nil {
		return nil, err
	}

	tx.Commit()

	return update.ToDomain(), nil
}

func (r *bookRepo) DeleteBooksById(ctx context.Context, id string) error {
	return r.db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Delete(&entities.Books{}, "id=?", id).Error; err != nil {
			return err
		}
		return nil
	})
}
