package database

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"log"
	"time"
)

type User struct {
	gorm.Model
	//ID       uint      `gorm:"not null;unique"`
	Name     string    `gorm:"not null;"`
	Age      uint      `gorm:"not null;"`
	Birthday time.Time `gorm`
}


func initDB() {
	db, err := gorm.Open("mysql", "root:123456@/redpacket?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		fmt.Println(err)
	}
	defer db.Close()
	log.Println(err)
	// 自动迁移模式
	db.AutoMigrate(&User{})
	user := User{Name: "Jinzhu", Age: 18, Birthday: time.Now()}
	db.Create(&user)
	println(user.ID)
}
