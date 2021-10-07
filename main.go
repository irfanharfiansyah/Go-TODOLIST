package main

import (
	"fmt"

	"codebrains.io/todolist/database"
	"codebrains.io/todolist/models"
	"github.com/gofiber/fiber/v2"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)
func helloWorld(c *fiber.Ctx) error {
	return c.SendString("Hello World")
}
func initDatabase() {
	var err error
	dsn := "root:@tcp(127.0.0.1:3306)/db_todolist?charset=utf8mb4&parseTime=True&loc=Local"
	database.DBConn, err =  gorm.Open(mysql.Open(dsn), &gorm.Config{})

	
	if err != nil {
		panic("Failed to connect to database!")
	}
	fmt.Println("Database connected!")
	database.DBConn.AutoMigrate(&models.Todo{})
	fmt.Println("Migrated DB")
}
func setupRoutes(app *fiber.App) {
	app.Get("/todos", models.GetTodos)
	app.Post("/todos", models.CreateTodo)
}
func main()  {
	app := fiber.New()
	initDatabase()
	app.Get("/", helloWorld)
	setupRoutes(app)
	app.Listen(":8000")
}