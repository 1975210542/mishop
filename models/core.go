package models

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var DB *gorm.DB
var err error

func init() {
	DB,err = gorm.Open("mysql", "root:GaoQK0109_@tcp(rm-bp1574319r8j71m273o.mysql.rds.aliyuncs.com:3306)/micom?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		fmt.Println("mysql 链接失败:",err)
	}
}
