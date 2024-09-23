package books

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/kunlanat/go-example/api/v1/books/model"
	"github.com/kunlanat/go-example/dto"
	"github.com/kunlanat/go-example/service"
)

func BookController(serv service.BookService) *fiber.App {
	app := fiber.New()
	ctl := bookCtl{
		serv: serv,
	}

	app.Post("/books", ctl.CreateBook)
	app.Get("/books", ctl.GetAllBooks)
	app.Get("/books/:id", ctl.GetBookByID)
	app.Patch("/books/:id", ctl.UpdateBookByID)
	app.Delete("/books/:id", ctl.DeleteBookByID)

	return app
}

type bookCtl struct {
	serv service.BookService
}

func (s *bookCtl) CreateBook(ctx *fiber.Ctx) error {
	book := model.DTOCreateBook{}
	if err := ctx.BodyParser(&book); err != nil {
		return err
	}
	data, err := s.serv.CreateBook(ctx.Context(), book)
	if err != nil {
		return err
	}

	return ctx.JSON(data)
}

func (s *bookCtl) GetAllBooks(ctx *fiber.Ctx) error {
	query := dto.QuerySearch{}
	ctx.QueryParser(&query)
	data, err := s.serv.GetAllBooks(ctx.Context(), &query)
	if err != nil {
		return err
	}
	return ctx.JSON(data)
}

func (s *bookCtl) GetBookByID(ctx *fiber.Ctx) error {
	id := ctx.Params("id")

	data, err := s.serv.GetBooksById(ctx.Context(), id)
	if err != nil {
		return err
	}
	return ctx.JSON(data)
}

func (s *bookCtl) UpdateBookByID(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	book := model.DTOUpdateBook{}
	if err := ctx.BodyParser(&book); err != nil {
		return err
	}
	data, err := s.serv.UpdateBookById(ctx.Context(), id, &book)
	if err != nil {
		return err
	}
	return ctx.JSON(data)
}

func (s *bookCtl) DeleteBookByID(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	if err := s.serv.DeleteBookByID(ctx.Context(), id); err != nil {
		return err
	}
	return ctx.Status(http.StatusOK).JSON("Deleted")
}
