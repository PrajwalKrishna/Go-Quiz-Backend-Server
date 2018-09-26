package leaderboardFunction

import (
    "fmt"
    //"github.com/gin-contrib/cors"                        // Why do we need this package?
    "github.com/gin-gonic/gin"                           // Using gin as microframework
    "github.com/jinzhu/gorm"                             //Using gorm as orm
    _ "github.com/jinzhu/gorm/dialects/sqlite"           //Using sqlite as db
)

type Leaderboard struct{
    ID uint `json:"id"`
    User_id uint `gorm:"unique_index:idx_name_code" json:"user_id"`
    Quiz_id uint `gorm:"unique_index:idx_name_code" json:"quiz_id"`
    Score int `json:"score"`
}

func AddToLeaderBoard(c *gin.Context) {
    db, err := gorm.Open("sqlite3", "./gorm.db")
    if err != nil {
        panic("failed to connect table")
    }
    defer db.Close()

   var leaderboard Leaderboard
   c.BindJSON(&leaderboard)
   err = db.Create(&leaderboard).Error;
   if(err != nil){
       e := db.Where("user_id = ? AND quiz_id = ?",leaderboard.User_id,leaderboard.Quiz_id).First(&leaderboard)
       if e != nil {
           c.AbortWithStatus(404)
           fmt.Println(e)
       }
       leaderboard.Score = 0
       db.Save(&leaderboard)
   }
   c.Header("access-control-allow-origin", "*") // Why am I doing this? Find out. Try running with this line commented
   c.JSON(200, leaderboard)
}

func GetGlobalLeaderBoard(c *gin.Context) {
    db, err := gorm.Open("sqlite3", "./gorm.db")
    if err != nil {
        panic("failed to connect table")
    }
    defer db.Close()
   type Result struct{
       Total int `json:"total"`
       User_name string `json:"user_name"`
       User_id uint `json:"id"`
   }
   var result []Result
   err = db.Raw("select sum(score) as total,user_name,user_id from leaderboards join users where user_id = users.id group by user_id order by total desc").Scan(&result).Error
   if err != nil {
      c.Header("access-control-allow-origin", "*") // Why am I doing this? Find out. Try running with this line commented
      c.AbortWithStatus(404)
      fmt.Println(err)
   }else {
      c.Header("access-control-allow-origin", "*") // Why am I doing this? Find out. Try running with this line commented
      c.JSON(200, result)
   }
}

func ShowQuizesForUser(c *gin.Context) {
    db, err := gorm.Open("sqlite3", "./gorm.db")
    if err != nil {
        panic("failed to connect table")
    }
    defer db.Close()
   user_id := c.Params.ByName("user_id")
   type Result struct{
       Title string `json:title`
       Quiz_id uint `json:"quiz_id"`
       Score int `json:"score"`
   }
   var result []Result
   err = db.Raw("select score,title,quiz_id from leaderboards join quizzes where quiz_id = quizzes.id and user_id=? order by score desc;",user_id).Scan(&result).Error;
   if err != nil {
      c.Header("access-control-allow-origin", "*") // Why am I doing this? Find out. Try running with this line commented
      c.AbortWithStatus(404)
      fmt.Println(err)
   }else {
      c.Header("access-control-allow-origin", "*") // Why am I doing this? Find out. Try running with this line commented
      c.JSON(200, result)
   }
}

func GetGenreLeaderBoard(c *gin.Context) {
    db, err := gorm.Open("sqlite3", "./gorm.db")
    if err != nil {
        panic("failed to connect table")
    }
    defer db.Close()
   genre_id := c.Params.ByName("genre_id")
   type Result struct{
       User_name string `json:"user_name"`
       User_id uint `json:"id"`
       Total int `json:"total"`
   }
   var result []Result
   err = db.Raw("select sum(score) as total,user_id,user_name from leaderboards join quizzes on quiz_id = quizzes.id join users on user_id=users.id where quizzes.genre_id = ? group by user_id order by total desc;",genre_id).Scan(&result).Error
   if err != nil {
      c.Header("access-control-allow-origin", "*") // Why am I doing this? Find out. Try running with this line commented
      c.AbortWithStatus(404)
      fmt.Println(err)
   }else {
      c.Header("access-control-allow-origin", "*") // Why am I doing this? Find out. Try running with this line commented
      c.JSON(200, result)
   }
}
func GetQuizLeaderBoard(c *gin.Context) {
    db, err := gorm.Open("sqlite3", "./gorm.db")
    if err != nil {
        panic("failed to connect table")
    }
    defer db.Close()

   quiz_id := c.Params.ByName("quiz_id")
   user_id := c.Params.ByName("user_id")
   var leaderboard Leaderboard
   db.Where("user_id = ? AND quiz_id = ?", user_id,quiz_id).First(&leaderboard)
   if err != nil {
      c.Header("access-control-allow-origin", "*") // Why am I doing this? Find out. Try running with this line commented
      c.AbortWithStatus(404)
      fmt.Println(err)
   }else {
      c.Header("access-control-allow-origin", "*") // Why am I doing this? Find out. Try running with this line commented
      c.JSON(200, leaderboard)
   }
}
func UpdateScore(c *gin.Context) {
    db, err := gorm.Open("sqlite3", "./gorm.db")
    if err != nil {
        panic("failed to connect table")
    }
    defer db.Close()

   id := c.Params.ByName("id")

   var leaderboard Leaderboard
   db.Where("id = ? ",id).First(&leaderboard)
   c.BindJSON(&leaderboard)
   db.Save(&leaderboard)
   c.Header("access-control-allow-origin", "*") // Why am I doing this? Find out. Try running with this line commented
   c.JSON(200, leaderboard)
}
