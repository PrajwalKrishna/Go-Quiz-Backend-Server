package userFunction

import (
   "fmt"
   //"github.com/gin-contrib/cors"                        // Why do we need this package?
   "github.com/gin-gonic/gin"
   "github.com/jinzhu/gorm"
   _ "github.com/jinzhu/gorm/dialects/sqlite"           // If you want to use mysql or any other db, replace this line
)

type User struct {
   ID uint `json:"id"`
   UserName string `gorm:"unique_index" json:"username"`
   PassWord string `json:"password"`
}

func DeleteUser(c *gin.Context) {
    db, err := gorm.Open("sqlite3", "./gorm.db")
    if err != nil {
        panic("failed to connect table")
    }
    defer db.Close()
   id := c.Params.ByName("id")
   var user User
   d := db.Where("id = ?", id).Delete(&user)
   fmt.Println(d)
   c.Header("access-control-allow-origin", "*")
   c.JSON(200, gin.H{"id #" + id: "deleted"})
}

func UpdateUser(c *gin.Context) {
    db, err := gorm.Open("sqlite3", "./gorm.db")
    if err != nil {
        panic("failed to connect table")
    }
    defer db.Close()

   var user User
   id := c.Params.ByName("id")
   if err := db.Where("id = ?", id).First(&user).Error; err != nil {
      c.AbortWithStatus(404)
      fmt.Println(err)
   }
   c.BindJSON(&user)
   db.Save(&user)
   c.Header("access-control-allow-origin", "*") // Why am I doing this? Find out. Try running with this line commented
   c.JSON(200, user)
}

func CreateUser(c *gin.Context) {
    db, err := gorm.Open("sqlite3", "./gorm.db")
    if err != nil {
        panic("failed to connect table")
    }
    defer db.Close()

   var user User
   c.BindJSON(&user)
   db.Create(&user)
   c.Header("access-control-allow-origin", "*") // Why am I doing this? Find out. Try running with this line commented
   c.JSON(200, user)
}

func GetUser(c *gin.Context) {
    db, err := gorm.Open("sqlite3", "./gorm.db")
    if err != nil {
        panic("failed to connect table")
    }
    defer db.Close()

   id := c.Params.ByName("id")
   var user User
   db.Where("id = ?", id).First(&user)
   c.Header("access-control-allow-origin", "*") // Why am I doing this? Find out. Try running with this line commented
   c.JSON(200, user)
}

func GetUsers(c *gin.Context) {
    db, err := gorm.Open("sqlite3", "./gorm.db")
    if err != nil {
        panic("failed to connect table")
    }
    defer db.Close()

   var users []User
   if err := db.Find(&users).Error; err != nil {
      c.Header("access-control-allow-origin", "*") // Why am I doing this? Find out. Try running with this line commented
      c.AbortWithStatus(404)
      fmt.Println(err)
   } else {
      c.Header("access-control-allow-origin", "*") // Why am I doing this? Find out. Try running with this line commented
      c.JSON(200, users)
   }
}
