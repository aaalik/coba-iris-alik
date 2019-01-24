package main

import (
    "github.com/jinzhu/gorm"
    _ "github.com/jinzhu/gorm/dialects/mysql"
    "github.com/kataras/iris"
    . "./controller"
    _ "./global"
    "./model"
    "log"
)


func main() {
    initDB()
    app := iris.Default()

    // Simple group: v1.
    v1 := app.Party("/v1")
    {
        v1.Get("/tambah", Tambah)
        v1.Get("/kurang", Kurang)
        v1.Get("/kali", Kali)
        v1.Get("/bagi", Bagi)
        v1.Post("/create", Create)
        v1.Get("/read", Read)
        v1.Post("/update", Update)
        v1.Delete("/delete", Delete)

    }

    app.Run(iris.Addr(":8080"))
}

func initDB(){
    var Db, err = gorm.Open("mysql", "root:qweasdzxc@/mamah?charset=utf8&parseTime=True&loc=Local")

    model.Db = Db
    if err != nil {
        log.Panic(err)
    }
    defer Db.Close()
    log.Println("Connection Established")
}