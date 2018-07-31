package service

import (
	"log"
	"time"
)

type Book struct {
	ID        int64     `xorm:"id pk" json:"id"`
	Name      string    `xorm:"name" json:"name"`
	CreatedAt time.Time `xorm:"created"`
	UpdatedAt time.Time `xorm:"updated"`
}

func NewBook() *Book {
	return new(Book)
}

func (b *Book) FindByID(id int64) bool {
	b.ID = id
	has, err := engine.Get(b)
	if err != nil {
		log.Println(err)
		return false
	}
	return has
}

func (b *Book) GetAll() []Book {
	var books []Book
	err := engine.Find(&books)
	if err == nil {
		return books
	}
	return nil
}

func (b *Book) GetByUserID(id int64) []Book {
	var books []Book
	err := engine.
		Join("INNER", "user_book", "book.id = user_book.book_id").
		Join("INNER", "user", "user_book.user_id = user.id").
		Where("user.id = ?", id).
		Find(&books)
	if err != nil {
		log.Println(err)
		return nil
	}
	return books
}
