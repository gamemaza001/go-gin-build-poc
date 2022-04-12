package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	model "go-api-poc/model"
	_ "github.com/lib/pq"
	"database/sql"
)

var books = []model.Book{
	{Unique_id: "1", Name: "Harry Potter", Age: 24},
	{Unique_id: "2", Name: "The Lord of the Rings", Age: 34},
	{Unique_id: "3", Name: "The Wizard of Oz", Age: 15},
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
		psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s",
		host, port, user, password, dbname)
		db, err := sql.Open("postgres", psqlInfo)
		if err != nil {
			panic(err)
		}
		defer db.Close()

		err = db.Ping()
		if err != nil {
			panic(err)
		}
		fmt.Println("Successfully connected!")
		fmt.Println("Hello World!")
	})

	r.POST("/books", func(c *gin.Context) {
		var book model.Book
	
		if err := c.ShouldBindJSON(&book); err != nil {
			fmt.Println(err);
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}
	
		books = append(books, book)
	
		c.JSON(http.StatusCreated, book)
	})

	r.Run()
}