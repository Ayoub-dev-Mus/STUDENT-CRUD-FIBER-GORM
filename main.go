package main

import (
	"STUDENTPROJECT/Database"
	"STUDENTPROJECT/Models"
	"STUDENTPROJECT/Services"
	"fmt"

	"github.com/gofiber/fiber"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func helloWorld(c *fiber.Ctx) {
	c.Send("Hello, World!")
}

func setupRoutes(app *fiber.App) {

	app.Get("/Studentapi/v1/AllStudent", Services.GetStudents)
	app.Get("/Studentapi/v1/CustomStudent/:id", Services.GetStudent)
	app.Post("/Studentapi/v1/AddStudent", Services.NewStudent)
	app.Delete("/Studentapi/v1/DeleteStudent/:id", Services.DeleteStudent)
	app.Put("/Studentapi/v1/UpdateStudent/:id", Services.UpdateStudent)
}

func initDatabase() {
	var err error
	Database.DBConn, err = gorm.Open(sqlite.Open("student.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	Database.DBConn.AutoMigrate(&Models.Student{})
	fmt.Println("Connection Opened to Database")
}

func main() {
	app := fiber.New()
	initDatabase()

	setupRoutes(app)
	app.Listen(3000)

}
