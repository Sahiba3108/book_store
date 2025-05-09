package routes

import (
	"log"
	"test/services"

	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	app.Get("/api/books", getBooks)
	app.Get("/api/books/search", searchBooks)
	app.Post("/api/books", addBook)
	app.Put("/api/books", updateBook)
	app.Delete("/api/books", deleteBook)
}

func getBooks(c *fiber.Ctx) error {
	books := services.Store.GetBooks()
	return c.JSON(books)
}

func addBook(c *fiber.Ctx) error {
	var payload struct {
		Title  string `json:"title"`
		Author string `json:"author"`
	}
	if err := c.BodyParser(&payload); err != nil || payload.Title == "" || payload.Author == "" {
		return c.Status(fiber.StatusBadRequest).SendString("Title and Author are required.")
	}
	newBook := services.Store.AddBook(payload.Title, payload.Author)
	return c.Status(fiber.StatusCreated).JSON(newBook)
}

func updateBook(c *fiber.Ctx) error {
	var payload struct {
		Id     int    `json:"id"`
		Title  string `json:"title"`
		Author string `json:"author"`
	}
	if err := c.BodyParser(&payload); err != nil || payload.Id == 0 || payload.Title == "" || payload.Author == "" {
		return c.Status(fiber.StatusBadRequest).SendString("ID, Title and Author are required.")
	}
	updatedBook, err := services.Store.UpdateBook(payload.Id, payload.Title, payload.Author)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Error updating book")
	}
	return c.JSON(updatedBook)
}

func deleteBook(c *fiber.Ctx) error {
	var payload struct {
		Id int `json:"id"`
	}
	if err := c.BodyParser(&payload); err != nil || payload.Id == 0 {
		log.Println("deleteBook: invalid request payload")
		return c.Status(fiber.StatusBadRequest).SendString("ID is required.")
	}
	log.Printf("deleteBook: attempting to delete book with id %d", payload.Id)
	if err := services.Store.DeleteBook(payload.Id); err != nil {
		log.Printf("deleteBook: error deleting book with id %d: %v", payload.Id, err)
		return c.Status(fiber.StatusInternalServerError).SendString("Error deleting book")
	}
	log.Printf("deleteBook: successfully deleted book with id %d", payload.Id)
	return c.SendStatus(fiber.StatusNoContent)
}

func searchBooks(c *fiber.Ctx) error {
	name := c.Query("name")
	if name == "" {
		return c.Status(fiber.StatusBadRequest).SendString("Search parameter 'name' is required.")
	}
	books := services.Store.SearchBooks(name)
	return c.JSON(books)
}
