package repository

import (
	"fmt"
	"log"

	"github.com/mahfuz110244/graphql-go-demo/graph/model"
	db "github.com/mahfuz110244/graphql-go-demo/internal/pkg/db/mysql"
)

//CreateAuthor create's author
func CreateAuthor(author model.Author) (int64, error) {

	stmt, err := db.Db.Prepare("INSERT INTO Authors(Name,Biography) VALUES(?,?)")
	if err != nil {
		log.Fatal(err)
		return 0, err
	}

	res, err := stmt.Exec(author.Name, author.Biography)
	if err != nil {
		log.Fatal(err)
		return 0, err
	}

	id, err := res.LastInsertId()
	if err != nil {
		log.Fatal(err)
		return 0, err
	}

	defer stmt.Close()
	log.Println("Row inserted!!")
	return id, nil
}

//GetAuthorByID return author with respective id
func GetAuthorByID(id *string) (*model.Author, error) {

	stmt, err := db.Db.Prepare("select * from Authors where id=?")
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	defer stmt.Close()
	rows, err := stmt.Query(id)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	defer rows.Close()
	var author model.Author
	for rows.Next() {
		err = rows.Scan(&author.ID, &author.Name, &author.Biography)
		if err != nil {
			log.Fatal(err)
		}
	}

	defer rows.Close()

	if err = rows.Err(); err != nil {
		log.Fatal(err)
		return nil, err
	}
	return &author, nil

}

//GetAllAuthors returns all authors
func GetAllAuthors() ([]*model.Author, error) {
	stmt, err := db.Db.Prepare("select * from Authors")
	if err != nil {
		return nil, err
	}

	rows, err := stmt.Query()
	if err != nil {
		return nil, err
	}

	var authors []*model.Author
	for rows.Next() {
		var author model.Author
		rows.Scan(&author.ID, &author.Name, &author.Biography)
		authors = append(authors, &author)
	}

	if err = rows.Err(); err != nil {
		log.Fatal(err)
		return nil, err
	}

	defer stmt.Close()
	defer rows.Close()

	return authors, err
}

//CreateBook creates new book
func CreateBook(book model.Book) (int64, error) {
	stmt, err := db.Db.Prepare("insert into Books(Title,Price,IsbnNo,AuthorID) VALUES(?,?,?,?)")
	fmt.Println(stmt)
	if err != nil {
		return 0, err
	}

	res, err := stmt.Exec(book.Title, book.Price, book.IsbnNo, book.Author.ID)
	if err != nil {
		return 0, err
	}

	id, err := res.LastInsertId()
	if err != nil {
		return 0, err
	}

	return id, nil
}

//GetBooksByID returns books by respective id
func GetBooksByID(id *string) (*model.Book, error) {
	stmt, err := db.Db.Prepare("select Books.ID,Books.Title,Books.Price,Books.IsbnNo,Authors.ID,Authors.Name,Authors.Biography from Books inner join Authors where Books.AuthorID = Authors.ID and Books.ID = ? ;")
	if err != nil {
		return nil, err
	}

	rows, err := stmt.Query(id)
	var bookID, title, isbn_no, authorID, name, biography string
	var price int
	if rows.Next() {
		err := rows.Scan(&bookID, &title, &price, &isbn_no, &authorID, &name, &biography)
		if err != nil {
			return nil, err
		}
	}

	book := &model.Book{
		ID:     bookID,
		Title:  title,
		Price:  price,
		IsbnNo: isbn_no,
		Author: &model.Author{
			ID:        authorID,
			Name:      name,
			Biography: biography,
		},
	}
	defer rows.Close()
	defer stmt.Close()
	return book, nil
}

//GetAllBooks returns all Books Data
func GetAllBooks() ([]*model.Book, error) {
	var books []*model.Book
	stmt, err := db.Db.Prepare("select Books.ID,Books.Title,Books.Price,Books.IsbnNo,Authors.ID,Authors.Name,Authors.Biography from Books inner join Authors where Books.AuthorID = Authors.ID;")
	if err != nil {
		return nil, err
	}

	rows, err := stmt.Query()
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var bookID, title, isbn_no, authorID, name, biography string
		var price int
		err := rows.Scan(&bookID, &title, &price, &isbn_no, &authorID, &name, &biography)
		if err != nil {
			return nil, err
		}

		book := &model.Book{
			ID:     bookID,
			Title:  title,
			Price:  price,
			IsbnNo: isbn_no,
			Author: &model.Author{
				ID:        authorID,
				Name:      name,
				Biography: biography,
			},
		}
		books = append(books, book)
	}

	return books, nil
}

//GetAllBooks returns all Books Data
func GetAllBooksByAuthorName(name string) ([]*model.Book, error) {
	var books []*model.Book
	stmt, err := db.Db.Prepare("select Books.ID,Books.Title,Books.Price,Books.IsbnNo,Authors.ID,Authors.Name,Authors.Biography from Books inner join Authors where Authors.Name = ? and Books.AuthorID = Authors.ID;")
	if err != nil {
		return nil, err
	}

	rows, err := stmt.Query(name)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var bookID, title, isbn_no, authorID, name, biography string
		var price int
		err := rows.Scan(&bookID, &title, &price, &isbn_no, &authorID, &name, &biography)
		if err != nil {
			return nil, err
		}

		book := &model.Book{
			ID:     bookID,
			Title:  title,
			Price:  price,
			IsbnNo: isbn_no,
			Author: &model.Author{
				ID:        authorID,
				Name:      name,
				Biography: biography,
			},
		}
		books = append(books, book)
	}

	return books, nil
}
