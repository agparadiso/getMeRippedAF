package main

import (
	"strconv"

	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/mattn/go-sqlite3"
)

type Users struct {
	ID        int    `gorm:"AUTO_INCREMENT" form:"id" json:"id"`
	Firstname string `gorm:"not null" form:"firstname" json:"firstname"`
	Lastname  string `gorm:"not null" form:"lastname" json:"lastname"`
}

type Diary struct {
	ID    int `gorm:"AUTO_INCREMENT" form:"id" json:"id"`
	User  Users
	Foods []Food
}

type Food struct {
	ID          int `gorm:"AUTO_INCREMENT" form:"id" json:"id"`
	Ingredients []Ingredient
}

type Ingredient struct {
	ID           int    `gorm:"AUTO_INCREMENT" form:"id" json:"id"`
	Name         string `gorm:"not null" form:"name" json:"name"`
	Protein      int    `form:"protein" json:"protein"`
	Carbohydrate int    `form:"carbohydrate" json:"carbohydrate"`
	Fat          int    `form:"fat" json:"fat"`
}

func main() {
	r := gin.Default()

	v1 := r.Group("api/v1")
	{
		v1.POST("/users", PostUser)
		v1.GET("/users", GetUsers)
		v1.GET("/users/:id", GetUser)
		v1.PUT("/users/:id", UpdateUser)
		v1.DELETE("/users/:id", DeleteUser)
		v1.GET("/ingredient/:id", GetIngredient)
		v1.GET("/food/:id", GetFood)
		v1.POST("/diary", AddFoodToDiary)
		v1.GET("/diarys", GetDiarys)
	}

	r.Run(":8080")
}

func InitDb() *gorm.DB {
	// Openning file
	db, err := gorm.Open("sqlite3", "./data.db")
	db.LogMode(true)
	// Error
	if err != nil {
		panic(err)
	}
	// Creating the table
	if !db.HasTable(&Users{}) {
		db.CreateTable(&Users{})
		db.Set("gorm:table_options", "ENGINE=InnoDB").CreateTable(&Users{})
	}

	if !db.HasTable(&Diary{}) {
		db.CreateTable(&Diary{})
		db.Set("gorm:table_options", "ENGINE=InnoDB").CreateTable(&Diary{})
	}

	return db
}

func PostUser(c *gin.Context) {
	db := InitDb()
	defer db.Close()

	var user Users
	c.Bind(&user)

	if user.Firstname != "" && user.Lastname != "" {
		// INSERT INTO "users" (name) VALUES (user.Name);
		db.Create(&user)
		// Display error
		c.JSON(201, gin.H{"success": user})
	} else {
		// Display error
		c.JSON(422, gin.H{"error": "Fields are empty"})
	}

	// curl -i -X POST -H "Content-Type: application/json" -d "{ \"firstname\": \"Thea\", \"lastname\": \"Queen\" }" http://localhost:8080/api/v1/users
	// http post http://127.0.0.1:8080/api/v1/users Content-Type=application/json firstname=Gabo lastname=Paradiso
}

func GetUsers(c *gin.Context) {
	// Connection to the database
	db := InitDb()
	// Close connection database
	defer db.Close()

	var users []Users
	// SELECT * FROM users
	db.Find(&users)

	// Display JSON result
	c.JSON(200, users)

	// curl -i http://localhost:8080/api/v1/users
}

func GetUser(c *gin.Context) {
	// Connection to the database
	db := InitDb()
	// Close connection database
	defer db.Close()

	id := c.Params.ByName("id")
	var user Users
	// SELECT * FROM users WHERE id = 1;
	db.First(&user, id)

	if user.ID != 0 {
		// Display JSON result
		c.JSON(200, user)
	} else {
		// Display JSON error
		c.JSON(404, gin.H{"error": "User not found"})
	}

	// curl -i http://localhost:8080/api/v1/users/1
}

func UpdateUser(c *gin.Context) {
	// Connection to the database
	db := InitDb()
	// Close connection database
	defer db.Close()

	// Get id user
	id := c.Params.ByName("id")
	var user Users
	// SELECT * FROM users WHERE id = 1;
	db.First(&user, id)

	if user.Firstname != "" && user.Lastname != "" {

		if user.ID != 0 {
			var newUser Users
			c.Bind(&newUser)
			result := Users{
				ID:        user.ID,
				Firstname: newUser.Firstname,
				Lastname:  newUser.Lastname,
			}

			// UPDATE users SET firstname='newUser.Firstname', lastname='newUser.Lastname' WHERE id = user.Id;
			db.Save(&result)
			// Display modified data in JSON message "success"
			c.JSON(200, gin.H{"success": result})
		} else {
			// Display JSON error
			c.JSON(404, gin.H{"error": "User not found"})
		}

	} else {
		// Display JSON error
		c.JSON(422, gin.H{"error": "Fields are empty"})
	}

	// curl -i -X PUT -H "Content-Type: application/json" -d "{ \"firstname\": \"Thea\", \"lastname\": \"Merlyn\" }" http://localhost:8080/api/v1/users/1
}

func DeleteUser(c *gin.Context) {
	// Connection to the database
	db := InitDb()
	// Close connection database
	defer db.Close()

	// Get id user
	id := c.Params.ByName("id")
	var user Users
	// SELECT * FROM users WHERE id = 1;
	db.First(&user, id)

	if user.ID != 0 {
		// DELETE FROM users WHERE id = user.Id
		db.Delete(&user)
		// Display JSON result
		c.JSON(200, gin.H{"success": "User #" + id + " deleted"})
	} else {
		// Display JSON error
		c.JSON(404, gin.H{"error": "User not found"})
	}

	// curl -i -X DELETE http://localhost:8080/api/v1/users/1
}

func GetIngredient(c *gin.Context) {
	id := c.Params.ByName("id")
	ingredient_id, _ := strconv.ParseInt(id, 0, 64)

	if ingredient_id == 1 {
		content := gin.H{"id": ingredient_id, "name": "egg", "Protein": 11, "Carbohydrate": 1, "Fat": 0}
		c.JSON(200, content)
	} else if ingredient_id == 2 {
		content := gin.H{"id": ingredient_id, "name": "egg", "Protein": 5, "Carbohydrate": 25, "Fat": 1}
		c.JSON(200, content)
	} else {
		content := gin.H{"error": "ingredient with id#" + id + " not found"}
		c.JSON(404, content)
	}
}

func GetFood(c *gin.Context) {
	id := c.Params.ByName("id")
	food_id, _ := strconv.ParseInt(id, 0, 64)

	if food_id == 1 {
		content := gin.H{"id": food_id, "ingredients": 1}
		c.JSON(200, content)
	} else if food_id == 2 {
		content := gin.H{"id": food_id, "ingredients": 2}
		c.JSON(200, content)
	} else {
		content := gin.H{"error": "food with id#" + id + " not found"}
		c.JSON(404, content)
	}
}

func AddFoodToDiary(c *gin.Context) {
	db := InitDb()
	defer db.Close()

	var diary Diary

	//var food Food
	//food := db.First(&food, c.Params.ByName("Foods"))
	//diary.Foods = append(diary.Foods, food)

	var user Users
	fmt.Println(c.Params.ByName("User"))
	db.First(&user, c.Params.ByName("User"))

	diary.User = user

	c.Bind(&diary)
	// INSERT INTO "diary";
	db.Create(&diary)
	// Display error
	c.JSON(201, gin.H{"success": diary})

	//http post http://127.0.0.1:8080/api/v1/diary User:1 Foods:1
}

func GetDiarys(c *gin.Context) {
	// Connection to the database
	db := InitDb()
	// Close connection database
	defer db.Close()

	var diary []Diary
	// SELECT * FROM users
	db.Find(&diary)

	// Display JSON result
	c.JSON(200, diary)

	// curl -i http://localhost:8080/api/v1/users
}
