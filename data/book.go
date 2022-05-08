package data

import (
	"database/sql"
	"fmt"
	"log"
)

type Book struct {
	Book_id  int64
	Title    string
	Author   string
	Jenre    string
	Bookcase string
	Amount   int
}

type bookService struct {
	db *sql.DB
}

type BookService interface {
	GetBook() []Book
	InsertBook(newBook *Book)
	UpdateBook(book *Book) (int64, error)
	DeleteBook(book *Book) (int64, error)
}

func CreateBookService(db *sql.DB) BookService {
	return &bookService{db}
}

func (b *bookService) GetBook() []Book {
	rows, err := b.db.Query("SELECT Book_id, Title, Author, Jenre, Bookcase, Amount FROM book")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	instances := make([]Book, 0)
	for rows.Next() {
		instance := Book{}
		err := rows.Scan(&instance.Book_id, &instance.Title, &instance.Author, &instance.Jenre, &instance.Bookcase, &instance.Amount)
		if err != nil {
			log.Fatal(err)
		}
		instances = append(instances, instance)
	}
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}
	return instances
}

func (b *bookService) InsertBook(newBook *Book) {
	stmt, err := b.db.Prepare("INSERT INTO book (Title, Author, Jenre, Bookcase, Amount ) VALUES (?,?,?,?,?);")
	if err != nil {
		log.Fatal(err)
	}

	res, err := stmt.Exec(newBook.Title, newBook.Author, newBook.Jenre, newBook.Bookcase, newBook.Amount)
	if err != nil {
		log.Fatal(err)
	} else {
		id, _ := res.LastInsertId()
		newBook.Book_id = id
		fmt.Println("new book add to Database ")
	}
}

//func (b *bookService) DeleteBook(book *Book) (int64, error) {
//	res, err := b.db.Exec("DELETE FROM book WHERE Book_id = ?", book.Book_id)
//	if err != nil {
//		return 0, err
//	} else {
//		return res.RowsAffected()
//	}
//}

func (b *bookService) UpdateBook(book *Book) (int64, error) {
	res, err := b.db.Exec("UPDATE book SET  Author = ?, Jenre = ?, Bookcase = ?, Amount = ? WHERE Title = ?", book.Author, book.Jenre, book.Bookcase, book.Amount, book.Title)
	if err != nil {
		return 0, err
	} else {
		return res.RowsAffected()
	}
}

func (b *bookService) DeleteBook(book *Book) (int64, error) {
	res, err := b.db.Exec("DELETE FROM book WHERE Book_id = ?", book.Book_id)
	if err != nil {
		return 0, err
	} else {
		return res.RowsAffected()
	}
}
