package Models


import(
	"gorm.io/gorm"

)

type Student struct {
	gorm.Model
	FirstName  string `json:"firstname"`
	LastName string `json:"lastname"`
	Age int    `json:"age"`
}