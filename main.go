package main

import (
	"fmt"
	"log"
	"github.com/gofiber/fiber/v2"
	"github.com/tupap1/gologin/server/repository"

)



func main() {
	fmt.Println("APP IS RUNNING...")

	// 1. Inicializa la base de datos.
	db, err := repository.InitDatabase()
	if err != nil {
		log.Fatalf("Error al inicializar la base de datos: %v", err)
	}

	// 2. Crea la instancia de Fiber.
	app := fiber.New()

	// 3. Pasa 'app' y la conexión 'db' a la función que configura las rutas.
	// Esta acción utiliza la variable 'db', resolviendo el error del compilador.
	fmt.Println(db)
	// 4. Inicia el servidor.
	log.Fatal(app.Listen(":3000"))
}




