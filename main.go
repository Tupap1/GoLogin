package main

import (
	"fmt"
	"log"
	"github.com/gofiber/fiber/v2"
	//"github.com/tupap1/gologin/server/repository"

)



func main() {
	fmt.Println("APP IS RUNNING...")

	// Crea la instancia de Fiber.
	app := fiber.New()



	app.Get("/daddy", func(c *fiber.Ctx) error {
		repository.InitDatabase()
		return c.SendString("base de datos inicializada")

})


	app.Get("/db", func(c *fiber.Ctx) error {

		var tablenames []TableName
		result := db.Raw("SHOW TABLES").Scan(&tablenames)
		if result.Error != nil {
			return c.Status(fiber.StatusInternalServerError).SendString(fmt.Sprintf("Error al consultar las tablas: %v", result.Error))
		}

		// Extraer los nombres de las tablas en un slice de strings para una respuesta más limpia.
		var names []string
		for _, tableName := range tableNames {
			names = append(names, tableName.TableName)
		}

		// Devolver la lista de nombres de tablas como un JSON.
		return c.JSON(names)
})


	app.Get("/", func(c *fiber.Ctx) error {
	return c.SendString("que rico, que rico, que rico ese totito")
})

	app.Listen(":8080")
}




