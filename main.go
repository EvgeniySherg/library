package main

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/labstack/echo"
	"library/config"
	"library/data"
)

func main() {
	config := config.GetConfiguration()
	db := data.CreateConnectionToDatabase(config.DatabaseSource)
	defer db.Close()

	bookController := data.CreateBookController(db)

	e := echo.New()
	e.GET("/books", bookController.GetBooks)
	e.POST("/books", bookController.CreateBook)
	e.PUT("/books", bookController.UpdateBook)
	e.DELETE("/books", bookController.DeleteBook)
	e.Logger.Fatal(e.Start(":8080"))
}
