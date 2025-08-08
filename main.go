package main

import (
	"fmt"
	"log"
	"github.com/gofiber/fiber/v2"
)



type Todo struc{
	ID int
	completed bool'json:"completed"'
	Body string'json:"body"'
}

func main(){
	fmt.Println("APP IS RUNNING...")
	app := fiber.New()

	app.Get("/", func (c *fiber.Ctx) error{
		return c.Status(200).JSON(fiber.Map{"msg": "Hello World"})
	})

	app.Post("/api/todo", func (c *fiber.Ctx) error{
		todo:= &Todo{}

		if err ;= c.BodyParser(todo); err != nil{
			return err
		}

		if Todo.Body == "" {
			return c.Status(400).JSON(fiber.Map{"error":" Todo body is required"})

		Todo.ID = len(todos)+1
		todos = append(todos, *todo)
		

		return c.Status(201).JSON(todo)
		}}
	)

	log.Fatal(app.Listen(":4000"))
}