package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	restful "github.com/emicklei/go-restful/v3"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

type Todo struct {
	ID int `json:"id"`
	Title string `json:"title"`
	Done bool `json:"done"`
	Body string `json:"body"`
}



func getAllTodos(c *gin.Context){

	connectionString := "root:root@tcp(localhost:3306)/test"

	db,err := sql.Open("mysql", connectionString)

	if err != nil {
		log.Fatal(err)
	}
	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to the database")

	rows,err := db.Query("SELECT * from todos")


	if err != nil {
		log.Println("Error",err)
	}
	
	defer rows.Close()

	var todos []Todo

	for rows.Next() {
		var todo Todo
		errs := rows.Scan(&todo.ID, &todo.Title, &todo.Done)
		if errs != nil {
			log.Println("Error scanning rows:", errs)
			// return
		}
		todos = append(todos, todo)
	}
	 c.IndentedJSON(http.StatusOK, todos)
}


func main() {

	ws := new(restful.WebService)

	ws.
		Path("/users").
		Consumes(restful.MIME_XML, restful.MIME_JSON).
		Produces(restful.MIME_JSON, restful.MIME_XML) // you can specify this per route as well

	tags := []string{"users"}



    router.Run("localhost:8080")


}
