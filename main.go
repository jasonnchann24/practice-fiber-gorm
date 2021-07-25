package main

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/jasonnchann24/fiber-gorm/book"
	"github.com/jasonnchann24/fiber-gorm/db"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func setupRoutes(app *fiber.App) {
	app.Get("/api/v1/books", book.GetBooks)
	app.Get("/api/v1/books/:id", book.GetBook)
	app.Post("/api/v1/books", book.CreateBook)
	app.Put("/api/v1/books/:id", book.UpdateBook)
	app.Delete("/api/v1/books/:id", book.DeleteBook)
}

func initDatabase() {
	var err error
	db.DBConn, err = gorm.Open(sqlite.Open("books.db"))
	if err != nil {
		panic("Failed to connect to database")
	}

	fmt.Println("Db connection success")

	db.DBConn.AutoMigrate(&book.Book{})
	fmt.Println("DB Migrated")
}

func main() {
	app := fiber.New()
	initDatabase()
	sqlDB, _ := db.DBConn.DB()
	defer sqlDB.Close()

	setupRoutes(app)

	app.Listen(":3000")

}
