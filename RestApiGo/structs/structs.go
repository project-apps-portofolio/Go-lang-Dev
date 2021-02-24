package structs

// ORM GORM
import gorm "github.com/jinzhu/gorm"

type Person struct {

	gorm.Model
	FirstName string
	LastName string

}