package data

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/labstack/echo"
	"net/http"
)

type bookController struct {
	service BookService
}

func CreateBookController(db *sql.DB) *bookController {
	newService := CreateBookService(db)
	return &bookController{service: newService}
}


func (bc *bookController) GetBooks(c echo.Context) error {
	books := bc.service.GetBook()
	return c.JSON(http.StatusOK, books)
}

func (bc *bookController) CreateBook(c echo.Context) error {
	a := c.Request().Body
	var book Book
	json.NewDecoder(a).Decode(&book)
	bc.service.InsertBook(&book)
	return c.JSON(http.StatusOK, book)
}

func (bc *bookController) UpdateBook(c echo.Context) error {
	var book Book
	err := json.NewDecoder(c.Request().Body).Decode(&book)
	if err != nil {
		fmt.Println(err)
	}
	a, _ := bc.service.UpdateBook(&book)
	fmt.Printf("строк изменено %d ", a)
	return c.JSON(http.StatusOK, book)
}

func (bc *bookController) DeleteBook(c echo.Context) error {
	var book Book
	err := json.NewDecoder(c.Request().Body).Decode(&book)
	if err != nil {
		fmt.Println(err)
	}
	a, _ := bc.service.DeleteBook(&book)
	fmt.Printf("строк удалено %d ", a)
	return c.JSON(http.StatusOK, book)
}


