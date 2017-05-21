package main

import (
	"github.com/agparadiso/getMeRippedAF/pkg/sqldb"
	"github.com/gin-gonic/gin"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	r := gin.Default()

	v1 := r.Group("api/v1")
	{
		//	v1.POST("/users", PostUser)
		//	v1.GET("/users", GetUsers)
		//	v1.GET("/users/:id", GetUser)
		//	v1.PUT("/users/:id", UpdateUser)
		//	v1.DELETE("/users/:id", DeleteUser)
		v1.GET("/ingredient/:id", sqldb.GetIngredient)
		v1.POST("/ingredient", sqldb.PostIngredient)
		v1.GET("/food/:id", sqldb.GetFood)
		v1.POST("/food", sqldb.PostFood)
		//	v1.POST("/diary", AddFoodToDiary)
		//	v1.GET("/diarys", GetDiarys)
	}

	r.Run(":8080")
}
