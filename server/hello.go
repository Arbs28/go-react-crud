package main

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)

type Todo struct {
	ID int `json:"id"`
	Title string `json:"title"`
	Done bool `json:"done"`
	Body string `json:"body"`
}

func main() {

	fmt.Print("Hello world")

	app := fiber.New()

	todos := []Todo{}


	app.Get("/test" ,func(c *fiber.Ctx) error{
		return c.SendString("Ok")
	})

	//get all Todos
	app.Get("/api/todos" ,func(c *fiber.Ctx) error{
		return c.JSON(todos)
	})


	//post a new todo pass title and body as json
	app.Post("/api/todos" ,func(c *fiber.Ctx) error{

		todo := &Todo{}
		
		if err := c.BodyParser(todo); err != nil {
			return err
		}

		 todo.ID = len(todos) + 1

		 todos = append(todos, *todo)

		 return c.JSON(todos)
	})


	//mark a todo done
	app.Patch("/api/todos/:id/done" ,func(c *fiber.Ctx) error {
		id,err := c.ParamsInt("id")

		if err != nil{
			return c.Status(401).SendString("Invalid Id")
		}

		for i,t := range todos {
			if t.ID == id {
				todos[i].Done = true
				break
			}
		}

		return c.JSON(todos)
	})

	//delete a todo
	app.Delete("/api/todos/:id" ,func(c *fiber.Ctx) error {
		id,err := c.ParamsInt("id")


		if err != nil{
			return c.Status(401).SendString("Invalid Id")
		}

		for i := len(todos) - 1; i >= 0; i-- {


			if todos[i].ID == id {
				todos = append(todos[:i],todos[i+1:]...)
			}
		}

		return c.JSON(todos)
	})



	app.Listen(":4000")
}