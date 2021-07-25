package book

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jasonnchann24/fiber-gorm/db"
	"gorm.io/gorm"
)

type Book struct {
	gorm.Model
	Title  string `json:"title"`
	Author string `json:"author"`
	Rating int    `json:"rating"`
}

func GetBooks(ctx *fiber.Ctx) error {
	db := db.DBConn
	var books []Book
	db.Find(&books)
	return ctx.JSON(books)
}

func GetBook(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	db := db.DBConn
	var book Book
	db.Find(&book, id)
	return ctx.JSON(book)
}

func CreateBook(ctx *fiber.Ctx) error {
	db := db.DBConn
	book := new(Book)
	if err := ctx.BodyParser(book); err != nil {
		return ctx.Status(fiber.StatusBadRequest).SendString(err.Error())
	}

	db.Create(&book)
	return ctx.JSON(book)
}

func UpdateBook(ctx *fiber.Ctx) error {
	return ctx.SendString("Update Book")
}

func DeleteBook(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	db := db.DBConn

	var book Book
	db.First(&book, id)
	if book.Title == "" {
		return ctx.Status(fiber.StatusNotFound).SendString("Not found")
	}

	db.Delete(&book)
	return ctx.SendString("book deleted")
}
