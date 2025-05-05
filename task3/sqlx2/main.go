package main

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3" // 匿名导入 SQLite 驱动
)

type Book struct {
	Id     int64  `db:"id"`
	Title  string `db:"title"`
	Author string `db:"author"`
	Price  int    `db:"price"`
}

func main() {
	db, err := sqlx.Connect("sqlite3", "./practice.db")
	if err != nil {
		panic("无法连接数据库")
	}
	err = db.Ping()
	if err != nil {
		panic("数据库不可达到")
	}
	fmt.Println("成功连接到数据库")
	var books []Book
	query := `
		 SELECT id, title, author, price
		FROM books
		 WHERE price > ?
		 ORDER BY price DESC
		`
	err = db.Select(&books, query, 100)
	if err != nil {
		panic(err)
	}

}
