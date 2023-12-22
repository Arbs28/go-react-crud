func main() {

	connectDB()

	defer db.Close()




	// Fetch all Todos from the database
	app.Get("/api/todos", func(c *fiber.Ctx) error {

        log.Println("all gud")
		
		
		
		rows, err := db.Query("SELECT * FROM todos")
        log.Println("all gud")
    	if err != nil {
        log.Println("Error executing query:", err)
        return c.Status(500).SendString("Internal Server Error")
	    }

		defer rows.Close()

		// var todos []Todo

   		 todos := []Todo{}
    
		 for rows.Next() {
        var todo Todo
        
		err := rows.Scan(&todo.ID, &todo.Title, &todo.Done)

        if err != nil {
            log.Println("Error scanning rows:", err)
            return c.Status(500).SendString("Internal Server Error")
        }
        todos = append(todos, todo)

		return c.JSON(todos)
    }

    if err := rows.Err(); err != nil {
        log.Println("Error iterating over rows:", err)
        return c.Status(500).SendString("Internal Server Error")
    }

    return c.JSON(todos)
})


//post a new todo pass title and body as json
app.Post("/api/todos", func(c *fiber.Ctx) error {
	todo := &Todo{}
	if err := c.BodyParser(todo); err != nil {
		return err
	}

	result, err := db.Exec("INSERT INTO todos (Title, Done) VALUES (?, ?)", todo.Title, todo.Done)
	if err != nil {
		log.Println(err)
		return c.Status(500).SendString("Internal Server Error")
	}

	lastInsertID, _ := result.LastInsertId()
	todo.ID = int(lastInsertID)

	return c.JSON(todo)
})


//mark a todo done
app.Patch("/api/todos/:id/done", func(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(401).SendString("Invalid Id")
	}

	_, err = db.Exec("UPDATE todos SET Done = true WHERE ID = ?", id)
	if err != nil {
		log.Println(err)
		return c.Status(500).SendString("Internal Server Error")
	}

	return c.SendString("Todo marked as done")
})

//delete a todo
app.Delete("/api/todos/:id", func(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(401).SendString("Invalid Id")
	}

	_, err = db.Exec("DELETE FROM todos WHERE ID = ?", id)
	if err != nil {
		log.Println(err)
		return c.Status(500).SendString("Internal Server Error")
	}

	return c.SendString("Todo deleted")
})



app.Listen(":4000")
}