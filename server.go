package main

import (
	"net/http"

	"github.com/arashage/todos-api/database"
	"github.com/arashage/todos-api/model"
	"github.com/gin-gonic/gin"
)

var todo model.Todo

func postHandler(c *gin.Context) {
	var rq model.Todo
	err := c.ShouldBindJSON(&rq)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	id := database.Insert(rq)
	rq.ID = id
	c.JSON(http.StatusCreated, rq)
}

func main() {

	r := gin.Default()

	r.POST("/todos", postHandler)
	r.Run(":1234")

}
