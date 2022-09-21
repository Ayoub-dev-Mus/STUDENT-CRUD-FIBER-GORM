package Services

import (
	"STUDENTPROJECT/Database"
	"STUDENTPROJECT/Models"

	"github.com/gofiber/fiber"
)

func GetStudents(c *fiber.Ctx) {
	db := Database.DBConn
	var students []Models.Student
	db.Find(&students)
	c.JSON(students)
}

func GetStudent(c *fiber.Ctx) {
	id := c.Params("id")
	db := Database.DBConn
	var students []Models.Student
	db.Find(&students, id)
	c.JSON(students)
}

func NewStudent(c *fiber.Ctx) {
	db := Database.DBConn
	student := new(Models.Student)
	if err := c.BodyParser(student); err != nil {
		c.Status(503).Send(err)
		return
	}
	db.Create(&student)
	c.JSON(student)
}

func UpdateStudent(c *fiber.Ctx) {

	student := new(Models.Student)
	id := c.Params("id")
	db := Database.DBConn
	if err := c.BodyParser(student); err != nil {
		c.Status(503).SendString(err.Error())
	}
	db.Where("id = ?", id).Updates(&student)
	c.Status(200).JSON(student)
}

func DeleteStudent(c *fiber.Ctx) {
	id := c.Params("id")
	db := Database.DBConn

	var student Models.Student
	db.First(&student, id)
	if student.FirstName == "" {
		c.Status(500).Send("No Student Found with ID")
		return
	}
	db.Delete(&student)
	c.Send("Student Successfully deleted")
}
