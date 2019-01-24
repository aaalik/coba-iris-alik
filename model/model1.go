package model

import (
    "github.com/jinzhu/gorm"
    _ "github.com/jinzhu/gorm/dialects/mysql"
    "time"
    _ "../global"
)

type Anakonda struct {
  Id int
  Nama string
  Lahir time.Time
}

var Db *gorm.DB

func Buat(nama string, lahir time.Time) {
  Db.AutoMigrate(&Anakonda{})

  Db.Create(&Anakonda{Nama: nama, Lahir: lahir})
}

func Baca()([]*Anakonda){
  Db.AutoMigrate(&Anakonda{})

  users := make([]*Anakonda, 0)
  Db.Find(&users)

  return users

}

func Ubah(nama string, lahir time.Time, id int){
  Db.AutoMigrate(&Anakonda{})

  Db.Model(&Anakonda{}).Where("id = ?", id).Update(Anakonda{Nama: nama, Lahir: lahir})
  //db.Create(&Anakonda{Nama: nama, Lahir: lahir})
}

func Hapus(id int){
  //db, err := gorm.Open("mysql", "root:qweasdzxc@/mamah?charset=utf8&parseTime=True&loc=Local")

  //if err != nil {
  //  log.Panic(err)
  //}
  //defer db.Close()
  //log.Println("Connection Established")
  Db.AutoMigrate(&Anakonda{})

  Db.Where("id = ?", id).Delete(&Anakonda{})
  //db.Create(&Anakonda{Nama: nama, Lahir: lahir})
}