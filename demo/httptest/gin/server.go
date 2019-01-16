package main

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

type Person struct {
	ID        uint   `json:”id”`
	FirstName string `json:”firstname”`
	LastName  string `json:”lastname”`
}

func main() {
	// gin.SetMode(gin.ReleaseMode)

	db, _ := gorm.Open("sqlite3", "./gorm.db")
	defer db.Close()

	db.AutoMigrate(&Person{})
	// db.Create(&p1)
	// db.Create(&p2)
	var p3 Person // identify a Person type for us to store the results in
	db.First(&p3) // Find the first record in the Database and store it in p3

	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": p3.LastName,
		})
	})
	r.Run() // listen and serve on 0.0.0.0:8080
}
