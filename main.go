package main

import (
	"fmt"
	"log"
	"time"
	"github.com/gofiber/fiber/v2"
	"gorm.io/driver/postgres"
 	"gorm.io/gorm"

)



func main(){
	dsn := "host=aws-0-sa-east-1.pooler.supabase.com user=postgres.pwrlfwrdsfoslavbyvde password=@Ndres137 dbname=postgres port=6543 sslmode=disable TimeZone=Asia/Shanghai"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to connect database: %v", err)
	}


	sqlDB, err := db.DB()
	if err != nil {
		log.Fatalf("failed to get sqlDB from gorm: %v", err)
	}

	// Configuración del pool de conexiones
	sqlDB.SetMaxIdleConns(10) // Conexiones inactivas
	sqlDB.SetMaxOpenConns(100) // Máximo de conexiones abiertas
	sqlDB.SetConnMaxLifetime(time.Hour) // Tiempo de vida de la conexión

	log.Println("Successfully connected to the database!")


	err = db.AutoMigrate(&Test{})
	if err != nil {
		log.Fatalf("fallo al migrar la base de datos: %v", err)
	}
	fmt.Println("Migración de la tabla 'users' completada.") 

	// 2. Crear un nuevo usuario (CREATE)
	newUser := Test{Name: "Juan Perez", Email: "juan.perez@example.com"}
	result := db.Create(&newUser)
	if result.Error != nil {
		log.Fatalf("fallo al crear usuario: %v", result.Error)
	}
	fmt.Printf("Usuario creado con ID: %d\n", newUser.ID)

	// 3. Consultar todos los usuarios (READ)
	var users []Test
	db.Find(&users)
	fmt.Println("\n--- Todos los usuarios ---")
	for _, user := range users {
		fmt.Printf("ID: %d, Nombre: %s, Email: %s\n", user.ID, user.Name, user.Email)
	}


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


	



}



