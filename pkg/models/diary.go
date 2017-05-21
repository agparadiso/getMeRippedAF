package models

type Diary struct {
	ID    int `gorm:"AUTO_INCREMENT" form:"id" json:"id"`
	Foods []Food
}

/*
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
*/
