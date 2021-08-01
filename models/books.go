package models

import (
	"time"

	"github.com/binhosemcrause/webapi-with-go/db"
)

type Book struct {
	Id          int
	Name        string
	Description string
	MediumPrice float32
	Author      string
	ImageURL    string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

func ShowAllBooks() []Book {
	db := db.ConectDataBase()

	selectBooks, err := db.Query("select * from books order by id asc")
	if err != nil {
		panic(err.Error())
	}

	b := Book{}
	books := []Book{}

	for selectBooks.Next() {
		var id int
		var name string
		var description string
		var medium_price float32
		var author string
		var image_url string
		var createat time.Time
		var updateat time.Time

		err = selectBooks.Scan(&id, &name, &description, &medium_price, &author, &image_url, &createat, &updateat)
		if err != nil {
			panic(err.Error())
		}

		b.Id = id
		b.Name = name
		b.Description = description
		b.Author = author
		b.ImageURL = image_url
		b.CreatedAt = createat
		b.UpdatedAt = updateat

		books = append(books, b)
	}

	defer db.Close()
	return books
}

func ShowBook(id string) Book {
	db := db.ConectDataBase()

	bookDB, err := db.Query("select * from books where id=$1", id)
	if err != nil {
		panic(err.Error())
	}

	println(bookDB)

	bookAtt := Book{}

	for bookDB.Next() {
		var id int
		var name string
		var description string
		var medium_price float32
		var author string
		var image_url string
		var createat time.Time
		var updateat time.Time

		err = bookDB.Scan(&id, &name, &description, &medium_price, &author, &image_url, &createat, &updateat)
		if err != nil {
			panic(err.Error())
		}

		bookAtt.Id = id
		bookAtt.Name = name
		bookAtt.Description = description
		bookAtt.Author = author
		bookAtt.ImageURL = image_url
		bookAtt.CreatedAt = createat
		bookAtt.UpdatedAt = updateat
	}

	defer db.Close()
	return bookAtt
}

func CreateBook(name string, description string, medium_price float32, author string, image_url string) {
	db := db.ConectDataBase()

	insertBook, err := db.Prepare("insert into books (name, description, medium_price, author, ImageURL) values ($1, $2, $3, $4, $5);")
	if err != nil {
		panic(err.Error())
	}

	insertBook.Exec(name, description, medium_price, author, image_url)
	defer db.Close()
}
