package main

import (
	"strconv"
	"github.com/gin-gonic/gin"
    "github.com/jinzhu/gorm"

    _ "github.com/mattn/go-sqlite3"
)

type Todo struct {
    gorm.Model
    Text   string
    Status string
}

//initDB
func dbInit() {
    db, err := gorm.Open("sqlite3", "test.sqlite3")
    if err != nil {
        panic("dbInit）")
    }
    db.AutoMigrate(&Todo{})
    defer db.Close()
}

//InsertData
func dbInsert(text string, status string) {
    db, err := gorm.Open("sqlite3", "test.sqlite3")
    if err != nil {
        panic("dbInsert)")
    }
    db.Create(&Todo{Text: text, Status: status})
    defer db.Close()
}

//UpdateData
func dbUpdate(id int, text string, status string) {
    db, err := gorm.Open("sqlite3", "test.sqlite3")
    if err != nil {
        panic("（dbUpdate)")
    }
    var todo Todo
    db.First(&todo, id)
    todo.Text = text
    todo.Status = status
    db.Save(&todo)
    db.Close()
}

//DeleteDB
func dbDelete(id int) {
    db, err := gorm.Open("sqlite3", "test.sqlite3")
    if err != nil {
        panic("（dbDelete)")
    }
    var todo Todo
    db.First(&todo, id)
    db.Delete(&todo)
    db.Close()
}

//GetAllData
func dbGetAll() []Todo {
    db, err := gorm.Open("sqlite3", "test.sqlite3")
    if err != nil {
        panic("(dbGetAll())")
    }
    var todos []Todo
    db.Order("created_at desc").Find(&todos)
    db.Close()
    return todos
}

//GetData
func dbGetOne(id int) Todo {
    db, err := gorm.Open("sqlite3", "test.sqlite3")
    if err != nil {
        panic("(dbGetOne())")
    }
    var todo Todo
    db.First(&todo, id)
    db.Close()
    return todo
}

func main() {
    router := gin.Default()
	router.LoadHTMLGlob("templates/*.html")
	
	dbInit()

    router.GET("/", func(ctx *gin.Context){
        ctx.HTML(200, "index.html", gin.H{})
	})
	
	data := "hello from internal"
    router.GET("/data", func(ctx *gin.Context){
    ctx.HTML(200, "data.html", gin.H{"data": data})
	})
	
	router.GET("/ping", func(c *gin.Context) {
        c.JSON(200, gin.H{
            "message": "pong",
        })
	 })

    router.GET("/todos", func(ctx *gin.Context) {
        todos := dbGetAll()
        ctx.HTML(200, "todos.html", gin.H{
            "todos": todos,
        })
	})
	
    router.POST("/todos/new", func(ctx *gin.Context) {
        text := ctx.PostForm("text")
        status := ctx.PostForm("status")
        dbInsert(text, status)
        ctx.Redirect(302, "/todos")
	})
	
    router.GET("/todos/detail/:id", func(ctx *gin.Context) {
        n := ctx.Param("id")
        id, err := strconv.Atoi(n)
        if err != nil {
            panic(err)
        }
        todo := dbGetOne(id)
        ctx.HTML(200, "detail.html", gin.H{"todo": todo})
    })

    //Update
    router.POST("/todos/update/:id", func(ctx *gin.Context) {
        n := ctx.Param("id")
        id, err := strconv.Atoi(n)
        if err != nil {
            panic("ERROR")
        }
        text := ctx.PostForm("text")
        status := ctx.PostForm("status")
        dbUpdate(id, text, status)
        ctx.Redirect(302, "/todos")
    })

    //削除確認
    router.GET("/todos/delete_check/:id", func(ctx *gin.Context) {
        n := ctx.Param("id")
        id, err := strconv.Atoi(n)
        if err != nil {
            panic("ERROR")
        }
        todo := dbGetOne(id)
        ctx.HTML(200, "delete.html", gin.H{"todo": todo})
    })

    //Delete
    router.POST("/todos/delete/:id", func(ctx *gin.Context) {
        n := ctx.Param("id")
        id, err := strconv.Atoi(n)
        if err != nil {
            panic("ERROR")
        }
        dbDelete(id)
        ctx.Redirect(302, "/todos")

    })

    router.Static("/static", "./static")



    router.Run()
}
