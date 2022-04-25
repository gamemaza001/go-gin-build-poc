package main

import (
	"database/sql"
	"fmt"
	model "go-api-poc/model"
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

var books = []model.Book{
	{Name: "Harry Potter", Age: 24},
	{Name: "The Lord of the Rings", Age: 34},
	{Name: "The Wizard of Oz", Age: 15},
}

var bookstest = []model.Book{
	{Name: "The Wizard of Oz", Age: 15},
}

const (
	host     = "rptcomm.postgres.database.azure.com"
	port     = 5432
	user     = "rpt_read"
	password = "rpt1234!!"
	dbname   = "rpt_poc"
)

func main() {

	r := gin.New()
	r.GET("/books", func(c *gin.Context) {
		c.JSON(http.StatusOK, books)
		fmt.Println("Hello World!")
	})

	r.POST("/books", func(c *gin.Context) {

		// model
		var book model.Book

		// config db
		psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
			"password=%s dbname=%s",
			host, port, user, password, dbname)

		// connect DB
		db, err := sql.Open("postgres", psqlInfo)

		if err := c.ShouldBindJSON(&book); err != nil {
			fmt.Println(err)
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}
		// db.SetConnMaxLifetime(0)
		// db.SetMaxOpenConns(10)
		// db.SetMaxIdleConns(5)

		insertStatement := `INSERT INTO books (Name, Age) VALUES ($1, $2)`

		_, err = db.Exec(insertStatement, book.Name, book.Age)
		// rows, err := db.Query(insertStatement, book.Name, book.Age)

		if err != nil {
			panic(err)
		}

		// books = append(books, book)

		// close connection
		db.Close()

		c.JSON(http.StatusCreated, book)
	})

	r.Run()
}
