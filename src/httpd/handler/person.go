package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Person A person.
type Person struct {
	FirstName string `json:"FirstName"`
	LastName  string `json:"LastName"`
	Age       int16  `json:"Age"`
	EyeColor  string `json:"EyeColor"`
	City      string `json:"City"`
	State     string `json:"State"`
}

// GetPersonDefault Returns a default person value
func GetPersonDefault(c *gin.Context) {
	p := Person{FirstName: "John", LastName: "Doe", Age: 18, EyeColor: "brown", City: "Jackson", State: "MI"}

	c.JSON(http.StatusOK, p)
}

// GetPersonByID returns the id searched.
func GetPersonByID(c *gin.Context) {
	id := c.Query("id")

	c.JSON(http.StatusOK, gin.H{
		"ID": id,
	})
}
