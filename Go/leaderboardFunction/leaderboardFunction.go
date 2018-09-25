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
       User_name string `json:"user_name"`
       User_id uint `json:"id"`
       Total int `json:"total"`
   }
   var result []Result
   err = db.Table("leaderboards").Select("user_id,sum(score) as total").Group("user_id").Order("total desc").Scan(&result).Error
   //err = db.Table("leaderboards").Joins("left join users on users.id = leaderboards.user_id").Select("users.user_name,user_id,sum(leaderboards.score) as total").Group("user_id").Order("total desc").Scan(&leaderboard).Error
   if err != nil {
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
       User_id uint `json:"user_id"`
       Quiz_id uint `json:"quiz_id"`
       Score int `json:"score"`
   }
   var result []Result
   err = db.Table("leaderboards").Where("user_id = ?", user_id).Scan(&result).Error;
   if err != nil {
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
   type Result struct{
       User_name string `json:"user_name"`
       User_id uint `json:"id"`
       Total int `json:"total"`
   }
   var result []Result
   err = db.Table("leaderboards").Select("user_id,sum(score) as total").Group("user_id").Order("total desc").Scan(&result).Error
   //err = db.Table("leaderboards").Joins("left join users on users.id = leaderboards.user_id").Select("users.user_name,user_id,sum(leaderboards.score) as total").Group("user_id").Order("total desc").Scan(&leaderboard).Error
   if err != nil {
      c.AbortWithStatus(404)
      fmt.Println(err)
   }else {
      c.Header("access-control-allow-origin", "*") // Why am I doing this? Find out. Try running with this line commented
      c.JSON(200, result)
   }
}
/*func UpdateAddScore(c *gin.Context) {
    db, err := gorm.Open("sqlite3", "./gorm.db")
    if err != nil {
        panic("failed to connect table")
    }
    defer db.Close()

   var leaderboard Leaderboard
   addent := c.Params.ByName("addent")
   if err := db.Where("id = ?", id).First(&leaderboard).Error; err != nil {
      c.AbortWithStatus(404)
      fmt.Println(err)
   }
   c.BindJSON(&leaderboard)
   db.Save(&leaderboard)
   c.Header("access-control-allow-origin", "*") // Why am I doing this? Find out. Try running with this line commented
   c.JSON(200, user)
}
*/
