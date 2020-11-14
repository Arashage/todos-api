package database

import (
	"database/sql"
	"log"
	"os"

	"github.com/arashage/todos-api/model"
	_ "github.com/lib/pq"
)

func Insert(todo model.Todo) (returnID int) {

	database := os.Getenv("DATABASE_URL")
	//database := "postgres://dazttmvw:JfDp5IwFbTcVtchOpDHBDxbI65iB9BF3@arjuna.db.elephantsql.com:5432/dazttmvw"
	db, err := sql.Open("postgres", database)
	if err != nil {
		log.Fatal("Connect to database error", err)
	}
	defer db.Close()

	ins, err := db.Query("INSERT INTO todos VALUES ($1, $2) RETURNING id", todo.Title, todo.Status)
	var id int
	err = ins.Scan(&id)

	if err != nil {
		log.Fatal("Can't ican id", err)
	}

	return id

}
