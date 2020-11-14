package main

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

type Todo struct {
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Status string `json:"status"`
}

func postHandler(c *gin.Context) {
	var rq Todo
	err := c.ShouldBindJSON(&rq)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	Insert(&rq)
	c.JSON(http.StatusCreated, rq)
}

func main() {

	r := gin.Default()

	r.POST("/todos", postHandler)
	r.Run(":1234")

}

func Insert(todo *Todo) {

	database := "postgres://dazttmvw:JfDp5IwFbTcVtchOpDHBDxbI65iB9BF3@arjuna.db.elephantsql.com:5432/dazttmvw"
	db, err := sql.Open("postgres", database)
	if err != nil {
		log.Fatal("Connect to database error", err)
	}
	defer db.Close()

	insert, err := db.Query("INSERT INTO todos VALUES ($1, $2) RETURNING id", todo.Title, todo.Status)
	var id int
	err = insert.Scan(&id)

	if err != nil {
		log.Fatal("Can't ican id", err)
	}

	todo.ID = id

}
