package controllers

import (
	"net/http"
	"restapigo/structs"
	"github.com/gin-gonic/gin"
)

func (idb *InDB) GetPerson(c *gin.Context) {
	var (
		person []structs.Person
		result gin.H
	)

	idb.DB.Find(&person)

	if len(person) <= 0 {

		result = gin.H{
			"result" : nil,
			"count" : 0,
		}

	} else {

		result = gin.H{
			"result": person,
			"count": len(person),
		}

	}

	c.JSON(http.StatusOK, result)
}

func (idb *InDB) PostPerson(c *gin.Context) {
	var (
		person structs.Person
		result gin.H
	)

	first_name := c.PostForm("firstname")
	last_name := c.PostForm("lastname")

	person.FirstName = first_name
	person.LastName = last_name

	idb.DB.Create(&person)

	result = gin.H{
		"results" : person,
	}

	c.JSON(http.StatusOK, result)
}