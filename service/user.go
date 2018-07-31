package service

import (
	"fmt"
	"log"
	"os"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
)

type User struct {
	ID        int64     `xorm:"id pk" json:"id"`
	Name      string    `xorm:"name" json:"name"`
	Password  string    `xorm:"password"`
	CreatedAt time.Time `xorm:"created"`
	UpdatedAt time.Time `xorm:"updated"`
}

var engine *xorm.Engine
var err error

func init() {
	driver := "mysql"
	dbUser := os.Getenv("SAMPLE_DB_USER")
	dbPassword := os.Getenv("SAMPLE_DB_PASS")
	dbHost := os.Getenv("SAMPLE_DB_HOST")
	dbName := os.Getenv("SAMPLE_DB_NAME")
	engine, err = xorm.NewEngine(driver, fmt.Sprintf("%s:%s@tcp(%s:3306)/%s", dbUser, dbPassword, dbHost, dbName))
	if err != nil {
		log.Println(err)
		panic("failed to connect database")
	}

	engine.ShowSQL(true)
}

func NewUser() *User {
	return new(User)
}

func (u *User) FindByID(id int64) bool {
	u.ID = id
	has, err := engine.Get(u)
	log.Println(err)
	log.Println(*u)
	log.Println(id)
	log.Println(has)
	return has
}

func (u *User) GetAll() []User {
	var users []User
	err := engine.Find(&users)
	if err == nil {
		return users
	}
	return nil
}
