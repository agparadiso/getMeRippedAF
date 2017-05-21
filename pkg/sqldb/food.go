package sqldb

import (
	"log"
	"strconv"

	"github.com/agparadiso/getMeRippedAF/pkg/models"
	"github.com/gin-gonic/gin"
)

func GetFood(c *gin.Context) {
	id := c.Params.ByName("id")
	foodID, err := strconv.ParseInt(id, 0, 64)
	if err != nil {
		log.Fatal("[Error] failed to parse foodID")
	}

	// Connection to the database
	db := InitDb()
	// Close connection database
	defer db.Close()

	var food models.Food
	// SELECT * FROM food
	db.First(&food, foodID)

	// Display JSON result
	c.JSON(200, food)

}

func PostFood(c *gin.Context) {
	db := InitDb()
	defer db.Close()

	var food models.Food
	c.Bind(&food)

	if food.Name != "" {
		db.Create(&food)
		// Display error
		c.JSON(201, gin.H{"success": food})
	} else {
		// Display error
		c.JSON(422, gin.H{"error": "Fields are empty"})
	}

	// http post http://localhost:8080/api/v1/food name=pancakes ingredients:='[1, 2]'
}
