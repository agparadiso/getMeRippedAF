package sqldb

import (
	"log"
	"strconv"

	"github.com/agparadiso/getMeRippedAF/pkg/models"
	"github.com/gin-gonic/gin"
)

func GetIngredient(c *gin.Context) {
	id := c.Params.ByName("id")
	ingredientID, err := strconv.ParseInt(id, 0, 64)
	if err != nil {
		log.Fatal("[Error] failed to parse ingredientID")
	}

	// Connection to the database
	db := InitDb()
	// Close connection database
	defer db.Close()

	var ingredients []models.Ingredient
	// SELECT * FROM ingredients
	db.First(&ingredients, ingredientID)

	// Display JSON result
	c.JSON(200, ingredients)

}

func PostIngredient(c *gin.Context) {
	db := InitDb()
	defer db.Close()

	var ingredient models.Ingredient
	c.Bind(&ingredient)

	if ingredient.Name != "" {
		db.Create(&ingredient)
		// Display error
		c.JSON(201, gin.H{"success": ingredient})
	} else {
		// Display error
		c.JSON(422, gin.H{"error": "Fields are empty"})
	}

	// http post http://localhost:8080/api/v1/ingredient name=egg protein=:11 carbohydrate:=11 fat:=0
}
