package database

type Inventory struct {
	ID         int    `json:"id" gorm:"column:ID;not null;primary_key"`
	MODEL_CODE string `json:"model_id" gorm:"column:MODEL_CODE;"`
	QUANTITY   int    `gorm:"not null;column:QUANTITY"`
	UNIT       string `json:"unit" gorm:"column:UNIT"`
	TIME       int64  `json:"time" gorm:"column:TIME"`
	//TYPE        int    `json:"type" gorm:"column:TYPE"`
	TYPE        []uint8 `json:"type" gorm:"column:TYPE"`
	SHIFT_BEGIN int64   `json:"shift_begin" gorm:"column:SHIFT_BEGIN"`
	SHIFT_END   int64   `json:"shift_end" gorm:"column:SHIFT_END"`
	SHIFT_NAME  string  `json:"shift_name" gorm:"column:SHIFT_NAME"`
	ACTUAL      int     `json:"actual" gorm:"column:ACTUAL"`
	EXPECT      int     `json:"expect" gorm:"column:EXPECT"`
}

// TableName 定义表名字
func (Inventory) table_inventory() string {
	return "inventory"
}

type JobLog struct {
	ID          uint   `json:"id" gorm:"column:ID;not null;primary_key;type:int(11)"`
	MODEL_ID    uint   `json:"model_id" gorm:"column:MODEL_ID;not null;type:int(11)"`
	RFID        string `json:"rfid" gorm:"column:RFID;type:varchar(25)"`
	DATE        int64  `json:"date" gorm:"column:DATE;type:bigint(20)"`
	CREATE_TIME int64  `json:"create_time" gorm:"column:CREATE_TIME;"`
	LOCATION_ID string `json:"location_id" gorm:"column:LOCATION_ID;"`
	STATUS      uint   `json:"status" gorm:"column:STATUS;"`
}

type Model struct {
	ID            uint   `json:"id" gorm:"column:ID;not null;primary_key"`
	NAME          string `json:"name" gorm:"column:NAME"`
	SERIAL_NUMBER string `json:"serial_number" gorm:"column:SERIAL_NUMBER"`
	CODE          string `json:"code" gorm:"column:CODE"`
}

type InventKey struct {
	ID          int   `json:"id"`
	MODEL_ID    uint  `json:"model_id"`
	SHIFT_BEGIN int64 `json:"shift_begin"`
	SHIFT_END   int64 `json:"shift_end"`
}
