package repository_test

import (
	"context"
	"testing"

	"github.com/kunlanat/go-example/api/v1/books/model"
	"github.com/kunlanat/go-example/domain"
	"github.com/kunlanat/go-example/repository"
	"github.com/kunlanat/go-example/repository/entities"
	"github.com/stretchr/testify/suite"
	"gorm.io/gorm"
)

type ExampleTestSuite struct {
	suite.Suite
	repo repository.BookRepository
	db   *gorm.DB
}

func (suite *ExampleTestSuite) SetupTest() {
	db := openTestDatabase()

	suite.db = db
	suite.repo = repository.BookRepositoryWithGORM(db)
}

func (suite *ExampleTestSuite) TestExample_GetAllBooks() {
	ctx := context.TODO()
	book := entities.Books{
		Name:  "Test",
		Desc:  "Test",
		Price: 100,
	}
	suite.Assert().NoError(suite.db.WithContext(ctx).Create(&book).Error)

	// Should be true
	{
		books, err := suite.repo.GetAllBooks(ctx, "")
		suite.Assert().NoError(err)
		suite.Assert().Equal(book.Name, books[0].Name)
	}

	// Should be true
	{
		books, err := suite.repo.GetAllBooks(ctx, "es")
		suite.Assert().NoError(err)
		suite.Assert().Equal(book.Name, books[0].Name)
	}

	// Should be nil
	{
		books, err := suite.repo.GetAllBooks(ctx, "example")
		suite.Assert().NoError(err)
		suite.Assert().Nil(books)
	}
}

func (suite *ExampleTestSuite) TestExample_CreateBook() {
	ctx := context.TODO()
	create := model.DTOCreateBook{
		Name:  "Test",
		Desc:  "Test",
		Price: 100,
	}

	// Should be created (CreateBook)
	output, err := suite.repo.CreateBook(ctx, &create)
	suite.Assert().NoError(err)
	suite.Assert().Equal(create.Name, output.Name)
}

func (suite *ExampleTestSuite) TestExample_GetBooksById() {
	ctx := context.TODO()
	create := entities.Books{
		Name:  "Test",
		Desc:  "Test",
		Price: 100,
	}
	suite.Assert().NoError(suite.db.WithContext(ctx).Create(&create).Error)

	{
		book, err := suite.repo.GetBooksById(ctx, create.ID)
		suite.Assert().NoError(err)
		suite.Assert().Equal(create.Name, book.Name)
	}
	{
		book, err := suite.repo.GetBooksById(ctx, "ID not found")
		suite.Assert().Error(err)
		suite.Assert().Nil(book)
	}
}

func (suite *ExampleTestSuite) TestExample_UpdateBookById() {
	ctx := context.TODO()
	create := entities.Books{
		Name:  "Test",
		Desc:  "Test",
		Price: 100,
	}
	suite.Assert().NoError(suite.db.WithContext(ctx).Create(&create).Error)

	update := domain.Books{
		Name: "UPDATE",
	}
	{
		out, err := suite.repo.UpdateBookById(ctx, create.ID, &update)
		suite.Assert().NoError(err)
		suite.Assert().Equal(update.Name, out.Name)
	}
	{
		out, err := suite.repo.UpdateBookById(ctx, "ID not found", &update)
		suite.Assert().Error(err)
		suite.Assert().Nil(out)
	}
}

func (suite *ExampleTestSuite) TestExample_DeleteBooksById() {
	ctx := context.TODO()
	create := entities.Books{
		Name:  "Test",
		Desc:  "Test",
		Price: 100,
	}
	suite.Assert().NoError(suite.db.WithContext(ctx).Create(&create).Error)

	{
		err := suite.repo.DeleteBooksById(ctx, create.ID)
		suite.Assert().NoError(err)
	}
}

func TestExampleTestSuite(t *testing.T) {
	suite.Run(t, new(ExampleTestSuite))
}
