package main

import (
	"fmt"
	"log"
	"github.com/gofiber/fiber/v2"
	"context"
 	//"gorm.io/driver/sqlite"
 	//"gorm.io/gorm"
	"os"
	"github.com/jackc/pgx/v5"
)



type Todo struct {
	ID int	`json:"id"`
	Completed bool `json:"completed"`
	Body string `json:"body"`
}

func main(){
	fmt.Println("APP IS RUNNING...")
	app := fiber.New()
	
	todos := []Todo{}

	app.Get("/", func (c *fiber.Ctx) error{
		return c.Status(200).JSON(fiber.Map{"msg": "Hello World"})
	})

	app.Post("/api/todo", func (c *fiber.Ctx) error {
		todo := &Todo{}

		if err := c.BodyParser(todo); err != nil{
			return err
		}

		if todo.Body == "" {
			return c.Status(400).JSON(fiber.Map{"error":" Todo body is required"})
			 }

		todo.ID = len(todos)+1
		todos = append(todos, *todo)
		

		return c.Status(201).JSON(todo)
	})



	app.Put("/api/todos/:id", func (c *fiber.Ctx) error {
		id := c.Params("id")

		for i, todo := range todos {
			if fmt.Sprint(todo.ID) == id {
				todos[i].Completed = true
				return c.Status(200).JSON(todos[i])
			}
		}

		return c.Status(401).JSON(fiber.Map{"body": "Item not found"})
	})

	log.Fatal(app.Listen(":4000"))


	fmt.Println(os.Getenv("DATABASE_URL"))

	
	conn, err := pgx.Connect(context.Background(), os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatalf("Failed to connect to the database: %v", err)
	}
	defer conn.Close(context.Background())

	// Example query to test connection
	var version string
	if err := conn.QueryRow(context.Background(), "SELECT version()").Scan(&version); err != nil {
		log.Fatalf("Query failed: %v", err)
	}

	log.Println("Connected to:", version)
}



