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

//func (bc *bookController) GetBooks() func(http.ResponseWriter, *http.Request) {
//	return func(w http.ResponseWriter, r *http.Request) {
//		books := bc.service.GetBook()
//		_, err := fmt.Fprintln(w, books)
//		if err != nil {
//			fmt.Println(err)
//		}
//	}
//} // старый вариант через http package

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

//
//func (bc *bookController) CreateBook() func(http.ResponseWriter, *http.Request) {
//	return func(w http.ResponseWriter, r *http.Request) {
//		method := r.Method
//		if method == "POST" {
//			var book Book
//			json.NewDecoder(r.Body).Decode(&book)
//			bc.service.InsertBook(&book)
//			_, err := fmt.Fprintln(w, book)
//			if err != nil {
//				fmt.Println(err)
//			}
//		}
//	}
//} // старый вариант через http package

//func (bc *bookController) UpdateBook() func(http.ResponseWriter, *http.Request) {
//	return func(w http.ResponseWriter, r *http.Request) {
//		newBook := bc.getBookFromRequest(r)
//		_, err := bc.service.UpdateBook(newBook)
//		if err != nil {
//			fmt.Println(err)
//		}
//	}
//}

//func (bc *bookController) getBookFromRequest(r *http.Request) *Book {
//	query := r.URL.Query()
//	i, _ := strconv.Atoi(query["amount"][0])
//	return &Book{
//		Title:    query["title"][0],
//		Author:   query["author"][0],
//		Jenre:    query["jenre"][0],
//		Bookcase: query["bookcase"][0],
//		Amount:   i,
//	}
//} Старый вариант для ввода всех параметров через поисковую строку
