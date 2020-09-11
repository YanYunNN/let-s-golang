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
	r.GET("/test/inventMap", testGatherInventMap)
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
	inventories := selectInventory()
	if inventories == nil {
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
	joblogs := selectJobLog()
	if joblogs == nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	}
	c.JSON(200, joblogs)
}
func testGatherInventMap(c *gin.Context) {
	inventMap := gatherJobLog2Inventory()
	if inventMap == nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	}
	c.JSON(200, "Job_log恢复计数完成")
}

func selectJobLog() []database.JobLog {
	db := database.GetDB()
	beginTime := time2.Now().AddDate(0, -6, 0).Unix()
	var joblogs []database.JobLog
	if err := db.Table("job_log").Where("CREATE_TIME >= ?", beginTime*1000).
		Where("LOCATION_ID =? ", "210.0.1.1").Find(&joblogs).
		Error; err != nil {
		fmt.Println(err)
	}
	return joblogs
}

func selectInventory() []database.Inventory {
	db := database.GetDB()
	var inventories []database.Inventory
	if err := db.Table("inventory").Where("TYPE=?", 1).Find(&inventories).Error; err != nil {
		fmt.Println(err)
	}
	return inventories
}

/*产品对照表*/
func getModelMap() map[string]uint {
	modelMap := make(map[string]uint)
	db := database.GetDB()
	var models []database.Model
	if err := db.Table("model").Find(&models).Error; err != nil {
		fmt.Println(err)
	}
	for i, model := range models {
		i = i
		//fmt.Printf("key:%d  value:%s \n", i, model)
		modelMap[model.CODE] = model.ID
	}
	return modelMap
}

/*初始化库存洞*/
func gatherInventMap() map[database.InventKey]int {
	modelMap := getModelMap()
	inventories := selectInventory
	inventMap := make(map[database.InventKey]int)
	for i, inventory := range inventories() {
		i = i
		//log.Println("遍历i：", i, "inventory:%s", inventory)
		//获取所有inventKey
		inventKey := database.InventKey{}
		inventKey.SHIFT_BEGIN = inventory.SHIFT_BEGIN
		inventKey.SHIFT_END = inventory.SHIFT_END
		inventKey.MODEL_ID = modelMap[inventory.MODEL_CODE]
		inventKey.ID = inventory.ID
		//放入map,初始计数为0
		inventMap[inventKey] = 0
	}
	return inventMap
}

/*分类更新*/
func gatherJobLog2Inventory() map[database.InventKey]int {
	//得到基础所有数据
	inventMap := gatherInventMap()
	jobLogs := selectJobLog()
	for i, jobLog := range jobLogs {
		i = i
		//fmt.Printf("i:%d  jobLog: %s \n", i, jobLog)
		for key, count := range inventMap {
			if key.MODEL_ID == jobLog.MODEL_ID &&
				key.SHIFT_END <= jobLog.CREATE_TIME &&
				jobLog.CREATE_TIME <= key.SHIFT_END {
				count++
				continue
			}
		}
	}
	//入库操作
	for key := range inventMap {
		key = key
		UpdateInventory2DB(key.ID, inventMap[key])
	}
	return inventMap
}

func UpdateInventory2DB(id int, count int) {
	db := database.GetDB()
	var inventory database.Inventory
	if err := db.Table("inventory").Where("id = ?", id).First(&inventory).Error; err != nil {
		fmt.Println(err)
	} else {
		//更新完工计数
		if err := db.Table("inventory").Model(&inventory).Update("ACTUAL", count).Error; err != nil {
			fmt.Println(err)
		}
		log.Println("更新一条完工记录:", inventory)
	}
}
