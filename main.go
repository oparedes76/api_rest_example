package main

import (
	"database/sql"
	"fmt"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
)

func main() {

	//db, err := sql.Open("postgres", "postgres://oparedes:ogPC.2k76@172.16.2.100:5432/testdb")
	db, err := sql.Open("postgres", "user=oparedes password=ogPC.2k76 dbname=test sslmode=disable")

	if err != nil {
		panic(err)
	}

	// Create a Gorm database instance.
	gormDB, err := gorm.Open("postgres", db)

	if err != nil {
		panic(err)
	}

	// Define your database models.
	type User struct {
		ID   int `gorm:"primary_key"`
		Name string
	}

	// Create a new Fiber instance.
	server := fiber.New()

	// Define a route for the "/" endpoint.
	// server.Get("/", func(c *fiber.Ctx) error {
	// 	// Send a simple message to the client.
	// 	return c.SendString("Hello, World!")
	// })

	// server.Get("/:name/:age/:gender?", func(c *fiber.Ctx) error {
	// 	msg := fmt.Sprintf("👴 %s is %s years old", c.Params("name"), c.Params("age"))
	// 	return c.SendString(msg) // => 👴 john is 75 years old
	// }).Name("")

	// Define a route for the "/" endpoint.
	server.Get("/:id/:name", func(c *fiber.Ctx) error {
		// Send a simple message to the client.
		// Create a new user.
		//user := User{ID: 3, Name: "pedro"}

		ui, err := strconv.Atoi(c.Params("id"))

		// Check for errors.
		if err != nil {
			fmt.Println("Error converting string to unsigned integer:", err)
			panic(err)
		}

		user := User{ID: ui, Name: c.Params("name")}
		c.BodyParser(&user)

		// Save the user to the database.
		gormDB.Create(&user)

		// Send the user to the client.
		c.JSON(user)

		return c.SendString("Registro Creado")
	}).Name("users")

	// Listen on port 3001.
	server.Listen(":3001")

}
