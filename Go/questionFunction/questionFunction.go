package questionFunction

import (
    "fmt"
    "github.com/gin-gonic/gin"                           // Using gin as microframework
    "github.com/jinzhu/gorm"                             //Using gorm as orm
    _ "github.com/jinzhu/gorm/dialects/sqlite"           //Using sqlite as db
)

type Question struct{
    ID uint `json:"id"`
    Question string `json:"question"`
    Answer string `gorm:"type:varchar(4)" json:"answer"`
    Quiz_id uint `json:"quiz_id"`
    Multi bool `json:"multi"`
}

func DataBaseOpener() *gorm.DB{
    db, err := gorm.Open("sqlite3", "./gorm.db")
    if err != nil {
        panic("failed to connect table")
    }
    defer db.Close()
    return db
}

func GetAllQuestions(c *gin.Context) {
    db, err := gorm.Open("sqlite3", "./gorm.db")
    if err != nil {
        panic("failed to connect table")
    }
    defer db.Close()
   quiz_id := c.Params.ByName("quiz_id")
   var question []Question
   if check := db.Where("quiz_id = ?", quiz_id).Find(&question).Error;
   check != nil {
      c.AbortWithStatus(404)
      fmt.Println(check)
   }else {
//      c.Header("access-control-allow-origin", "*") // Why am I doing this? Find out. Try running with this line commented
      c.JSON(200, question)
   }
}

func GetQuestion(c *gin.Context) {
    db, err := gorm.Open("sqlite3", "./gorm.db")
    if err != nil {
        panic("failed to connect table")
    }
    defer db.Close()
   id := c.Params.ByName("id")
   var question []Question
   if check := db.Where("id = ?", id).First(&question).Error;
   check != nil {
      c.AbortWithStatus(404)
      fmt.Println(check)
   }else {
//      c.Header("access-control-allow-origin", "*") // Why am I doing this? Find out. Try running with this line commented
      c.JSON(200, question)
   }
}

func AddQuestion(c *gin.Context) {
    db, err := gorm.Open("sqlite3", "./gorm.db")
    if err != nil {
        panic("failed to connect table")
    }
    defer db.Close()

   var question Question
   if err = c.BindJSON(&question); err != nil{
       c.JSON(400,err)
       return
   }
   fmt.Println(question)
   if check := db.Create(&question).Error;
   check != nil{
        c.AbortWithStatus(404)
        fmt.Println(check)
   }else{
       //c.Header("access-control-allow-origin", "*") // Why am I doing this? Find out. Try running with this line commented
       c.JSON(200, question)
   }
}

func DeleteQuestion(c *gin.Context) {
    db, err := gorm.Open("sqlite3", "./gorm.db")
    if err != nil {
        panic("failed to connect table")
    }
    defer db.Close()
   id := c.Params.ByName("id")
   var question Question
   check := db.Where("id = ?", id).Delete(&question)
   if check != nil{
       c.AbortWithStatus(404)     //To be decided
       fmt.Println(err)
   }
   //c.Header("access-control-allow-origin", "*")
   c.JSON(200, gin.H{"id #" + id: "deleted"})
}
