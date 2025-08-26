package main

import (
	"fmt"
	//"log"
	"github.com/gofiber/fiber/v2"
	//"github.com/tupap1/gologin/server/repository"

)



func main() {
	fmt.Println("APP IS RUNNING...")

	// 1. Inicializa la base de datos.
	//repository.InitDatabase()


	// 2. Crea la instancia de Fiber.
	app := fiber.New()

	// 3. Pasa 'app' y la conexión 'db' a la función que configura las rutas.
	// Esta acción utiliza la variable 'db', resolviendo el error del compilador.
	// 4. Inicia el servidor.
	//log.Fatal(app.Listen(":8080"))

	app.Get("/hola", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!sss")
	})

	app.Listen(":8080")
}




