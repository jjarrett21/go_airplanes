// Package Main
package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/iron-io/iron_go/worker"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

var db *gorm.DB
var err error

// Airplane object.
type Airplane struct {
	ID               int    `json:"id"`
	Manufacturer     string `json:"manufacturer"`
	Model            string `json:"model"`
	Year             int    `json:"year"`
	SinglePilotRated bool   `json:"singlepilotrated"`
	TopSpeed         int    `json:"topspeed"`
	Engine           string `json:"engine"`
}

// Main function that setups our database and http requests
func main() {

	worker.ParseFlags() //To use the docker dependency

	// Using a sqlite databse for lightweight and easy access.
	db, err = gorm.Open("sqlite3", "./gorm.db")
	if err != nil {
		fmt.Println("Failed to connect to databse")
	}
	defer db.Close()

	// Handles the data in our table in prod we would manage the SCHEMA$ better but for prototyping this is good.
	db.AutoMigrate(&Airplane{})

	r := gin.Default()
	r.POST("/airplane", CreateAirplane)
	r.PUT("/airplane/:id", UpdateAirplane)
	r.GET("/airplane/:id", GetAirplane)
	r.GET("/airplanes/", GetAirplanes)
	r.DELETE("/airplane/:id", DestroyAirplanes)

	r.Run(":8080")
}

// CreateAirplane : Allows the user to create an airplane in the db
// System will only populate provided values and will leave the rest blank
// Binding the object to Json and returning the expected '200' status
func CreateAirplane(c *gin.Context) {
	var airplane Airplane
	c.BindJSON(&airplane)
	db.Create(&airplane)
	c.JSON(200, airplane)
}

// UpdateAirplane : Allows the user to update an airplane instance that is already in the db
func UpdateAirplane(c *gin.Context) {
	var airplane Airplane
	id := c.Params.ByName("id")

	if err := db.Where("id = ?", id).First(&airplane).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	}
	c.BindJSON(&airplane)

	db.Save(&airplane)
	c.JSON(200, airplane)

}

// GetAirplane : Allows the user to request an airplane based on id #.
func GetAirplane(c *gin.Context) {
	id := c.Params.ByName("id")
	var airplane Airplane
	if err := db.Where("id = ?", id).First(&airplane).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, airplane)
	}
}

// GetAirplanes : Allows the user to get a list of all airplanes.
func GetAirplanes(c *gin.Context) {
	var planes []Airplane
	if err := db.Find(&planes).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, planes)
	}
}

// DestroyAirplanes : Allows user to delete airplanes from the databse by Id #.
func DestroyAirplanes(c *gin.Context) {
	id := c.Params.ByName("id")
	var airplane Airplane
	d := db.Where("id = ?", id).Delete(&airplane)
	fmt.Println(d)
	c.JSON(200, gin.H{"id #" + id: "deleted"})
}
