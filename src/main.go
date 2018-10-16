package main

import (
   "fmt"
  "github.com/gin-contrib/cors"                        // Why do we need this package?
   "github.com/gin-gonic/gin"                           // Using gin as microframework
   "github.com/jinzhu/gorm"                             //Using gorm as orm
   _ "github.com/jinzhu/gorm/dialects/sqlite"           //Using sqlite as db

   user "./userFunction"
   genre  "./genreFunction"
   question "./questionFunction"
   quiz "./quizFunction"
   leaderboard "./leaderboardFunction"
)

var db *gorm.DB                                         // declaring the db globally
var err error

func main() {
   fmt.Println("Starting")
   db, err = gorm.Open("sqlite3", "./gorm.db")
   if err != nil {
       panic("failed to connect table")
   }
   defer db.Close()

   //Commands for database
   db.AutoMigrate(&user.User{})
   db.AutoMigrate(&genre.Genre{})
   db.AutoMigrate(&quiz.Quiz{})
   db.AutoMigrate(&question.Question{})
   db.AutoMigrate(&leaderboard.Leaderboard{})
   //db.Model(&question.Question{}).AddForeignKey("genre_id", "genres(id)", "RESTRICT", "RESTRICT")
   r := gin.Default()                                        //Starting gin server

   //APIs for User
   r.GET("/user/", user.GetUsers)
   r.GET("/user/:id", user.GetUser)
   r.POST("/user", user.CreateUser)
   r.POST("/validateUser", user.ValidateUser)
   r.PUT("/user/:id", user.UpdateUser)
   r.DELETE("/user/:id", user.DeleteUser)



   //APIs related to Genre
   r.GET("/genreList/",genre.GetAllGenre)
   r.GET("/genre/:id",genre.GetGenre)
   r.POST("/genre",genre.AddGenre)
   r.DELETE("/genre/:id",genre.DeleteGenre)

   //APIs related to quiz
   r.GET("/quizes/:genre_id",quiz.GetAllQuizs)
   r.GET("/quiz/:id",quiz.GetQuiz)
   r.POST("/quiz",quiz.AddQuiz)
   r.DELETE("/quiz/:id",quiz.DeleteQuiz)


   //APIs related to Questions
   r.GET("/questions/:quiz_id",question.GetAllQuestions)
   r.GET("/question/:id",question.GetQuestion)
   r.POST("/question",question.AddQuestion)
   r.PUT("/question/:id", question.UpdateQuestion)
   r.DELETE("/question/:id",question.DeleteQuestion)

   //APIs related to leaderboard
   r.GET("/hometable/:user_id",leaderboard.ShowQuizesForUser)
   r.POST("/leaderboard/add",leaderboard.AddToLeaderBoard)
   r.PUT("/leaderboard/:id",leaderboard.UpdateScore)

   //APIs related to global leaderboard
   r.GET("/leaderboard/",leaderboard.GetGlobalLeaderBoard);
   r.GET("/leaderboard/:genre_id",leaderboard.GetGenreLeaderBoard);
   r.GET("/leaderboardDisplay/:user_id/:quiz_id",leaderboard.GetQuizLeaderBoard);


   r.Use((cors.Default()))
   r.Run(":8080")                                           // Run on port 8080
}
