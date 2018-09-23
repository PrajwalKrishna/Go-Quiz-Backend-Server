package main

import (
   "fmt"
 //  "github.com/gin-contrib/cors"                        // Why do we need this package?
   "github.com/gin-gonic/gin"                           // Using gin as microframework
   "github.com/jinzhu/gorm"                             //Using gorm as orm
   _ "github.com/jinzhu/gorm/dialects/sqlite"           //Using sqlite as db

   genre  "./genreFunction"
   question "./questionFunction"
   quiz "./quizFunction"
)

var db *gorm.DB                                         // declaring the db globally
var err error

type Person struct{
       ID uint `json:"id"`
       FirstName string `json:"firstname"`
       LastName string `json:"lastname"`
       City string `json:"city"`
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
   db.AutoMigrate(&quiz.Quiz{})
   db.AutoMigrate(&question.Question{})
   //db.Model(&question.Question{}).AddForeignKey("genre_id", "genres(id)", "RESTRICT", "RESTRICT")
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

   //APIs related to quiz
   r.GET("/quizs/:genre_id",quiz.GetAllQuizs)
   r.GET("/quiz/:id",quiz.GetQuiz)
   r.POST("/quiz",quiz.AddQuiz)
   r.DELETE("/quiz/:id",quiz.DeleteQuiz)


   //APIs related to Questions
   r.GET("/questions/:quiz_id",question.GetAllQuestions)
   r.GET("/question/:id",question.GetQuestion)
   r.POST("/question",question.AddQuestion)
   r.DELETE("/question/:id",question.DeleteQuestion)

   r.Run(":8080")                                           // Run on port 8080
}
