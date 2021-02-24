package controllers

// ORM GORM
import gorm "github.com/jinzhu/gorm"

// Menghubungkan dengan database ORM
type InDB struct {

	DB *gorm.DB

}