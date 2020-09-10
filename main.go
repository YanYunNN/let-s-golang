package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	time2 "time"
	"yanyun.com/src/database"
)

//var database *gorm.DB
var err error

func init() {
	r := gin.Default()
	r.GET("/users", GetAll)
	r.GET("/inventory", GetInventory)
	r.POST("/inventory", CreateInventory)
	r.PUT("/inventory/:id", UpdateInventory)
	r.DELETE("/inventory/:id", DeleteInventory)
	r.GET("/job_log/:id", GetJobLog)
	r.GET("/job/finished", getJobLogWhere)
	r.Run(":8080")
}

func main() {
	log.Println("start run-------")
}

func GetAll(c *gin.Context) {
	db := database.GetDB()
	var users []database.User
	if err := db.Find(&users).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, users)
	}
}

func GetInventory(c *gin.Context) {
	//db := database.DB.Self
	db := database.GetDB()
	var inventories []database.Inventory
	if err := db.Table("inventory").Find(&inventories).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, inventories)
	}
}

func GetJobLog(c *gin.Context) {
	id := c.Params.ByName("id")
	//db := database.DB.Self
	db := database.GetDB()
	var joblogs []database.JobLog
	if err := db.Table("job_log").Where("ID=?", id).
		First(&joblogs).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		for i := range joblogs {
			jobLog := joblogs[i]
			println(jobLog.ID)
		}
		c.JSON(200, joblogs)
	}
}

func CreateInventory(c *gin.Context) {
	var inventory database.Inventory
	db := database.GetDB()
	c.BindJSON(&inventory)
	db.Table("inventory").Create(&inventory)
	c.JSON(200, inventory)
}

func DeleteInventory(c *gin.Context) {
	db := database.GetDB()
	id := c.Params.ByName("id")
	var inventory database.Inventory
	d := db.Where("id = ?", id).Delete(&inventory)
	fmt.Println(d)
	c.JSON(200, gin.H{"id #" + id: "deleted"})
}

func UpdateInventory(c *gin.Context) {
	db := database.GetDB()
	var inventory database.Inventory
	id := c.Params.ByName("id")
	if err := db.Where("id = ?", id).First(&inventory).
		Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	}
	c.BindJSON(&inventory)
	db.Save(&inventory)
	c.JSON(200, inventory)
}

func getJobLogWhere(c *gin.Context) {
	db := database.GetDB()
	beginTime := time2.Now().AddDate(0, -4, 0).Unix()
	log.Println(beginTime * 1000)
	var joblogs []database.JobLog
	if err := db.Table("job_log").Where("CREATE_TIME >= ?", beginTime*1000).
		Where("LOCATION_ID =? ", "210.0.1.1").Find(&joblogs).
		Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	}
	c.JSON(200, joblogs)
}
