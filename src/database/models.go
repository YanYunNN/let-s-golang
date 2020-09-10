package database

import (
	"time"
)

type Inventory struct {
	ID       uint      `gorm:"not null;unique"`
	MODEL_ID uint      `gorm:"not null;"`
	QUANTITY uint      `gorm:"not null;"`
	DATE     time.Time `gorm`
}

// TableName 定义表名字
func (Inventory) table_inventory() string {
	return "inventory"
}

type JobLog struct {
	ID          uint   `json:"id" gorm:"column:ID;not null;primary_key;type:int(11)"`
	MODEL_ID    uint   `json:"model_id" gorm:"column:MODEL_ID;not null;type:int(11)"`
	RFID        string  `json:"rfid" gorm:"column:RFID;type:varchar(25)"`
	DATE        int64 `json:"date" gorm:"column:DATE;type:bigint(20)"`
	CREATE_TIME int64 `json:"create_time" gorm:"column:CREATE_TIME;"`
	LOCATION_ID string `json:"location_id" gorm:"column:LOCATION_ID;"`
	STATUS      uint   `json:"status" gorm:"column:STATUS;"`
}