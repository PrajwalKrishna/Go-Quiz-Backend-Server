package main

import (
   "fmt"
 //  "github.com/gin-contrib/cors"                        // Why do we need this package?
   "github.com/gin-gonic/gin"                           // Using gin as microframework
   "github.com/jinzhu/gorm"                             //Using gorm as orm
   _ "github.com/jinzhu/gorm/dialects/sqlite"           //Using sqlite as db

   genre  "./genreFunction"
)

var db *gorm.DB                                         // declaring the db globally
var err error

type Person struct{
       ID uint `json:"id"`
       FirstName string `json:"firstname"`
       LastName string `json:"lastname"`
       City string `json:"city"`
}


type Question struct{
    ID uint `json:"id"`
    question string `json:"question"`
    answer string `json:"answer"`
}

func main() {
   fmt.Println("Starting")
   db, err = gorm.Open("sqlite3", "./gorm.db")
   if err != nil {
       panic("failed to connect table")
   }
   defer db.Close()

   //Commands for database
   //db.AutoMigrate(&Person{})
   db.AutoMigrate(&genre.Genre{})
   //db.AutoMigrate(&Question{})

   r := gin.Default()                                        //Starting gin server

   /*r.GET("/people/", GetPeople)                             // Creating routes for each functionality
   r.GET("/people/:id", GetPerson)
   r.POST("/people", CreatePerson)
   r.PUT("/people/:id", UpdatePerson)
   r.DELETE("/people/:id", DeletePerson)*/
   //r.Use((cors.Default()))



   //APIs related to Genre
   r.GET("/genreList/",genre.GetAllGenre)
   r.GET("/genre/:id",genre.GetGenre)
   r.POST("/genre",genre.AddGenre)
   r.DELETE("/genre/:id",genre.DeleteGenre)


   r.Run(":8080")                                           // Run on port 8080
}


/*
func GetAllGenre(c *gin.Context) {
   var genre []Genre
   if err := db.Find(&genre).Error; err != nil {
      c.AbortWithStatus(404)
      fmt.Println(err)
   }else {
//      c.Header("access-control-allow-origin", "*") // Why am I doing this? Find out. Try running with this line commented
      c.JSON(200, genre)
   }
}

func DeleteGenre(c *gin.Context) {
   id := c.Params.ByName("id")
   var genre Genre
   err := db.Where("id = ?", id).Delete(&genre)
   if err != nil{
       c.AbortWithStatus(404)     //To be decided
       fmt.Println(err)
   }
   //c.Header("access-control-allow-origin", "*")
   c.JSON(200, gin.H{"id #" + id: "deleted"})
}

func AddGenre(c *gin.Context) {
   var genre Genre
   c.BindJSON(&genre)
   fmt.Println(genre)
   db.Create(&genre)
   //c.Header("access-control-allow-origin", "*") // Why am I doing this? Find out. Try running with this line commented
   c.JSON(200, genre)
}

func GetGenre(c *gin.Context) {
   id := c.Params.ByName("id")
   var genre Genre
   if err := db.Where("id = ?", id).First(&genre).Error; err != nil {
      c.AbortWithStatus(404)
      fmt.Println(err)
   } else {
      //c.Header("access-control-allow-origin", "*") // Why am I doing this? Find out. Try running with this line commented
      c.JSON(200, genre)
   }
}
*/
