package appconstant

import (
	"gorm.io/gorm"
)

var context *gorm.DB
var err error


const connectionString = "root:faizan123@tcp(127.0.0.1)/godb?charset=utf8mb4&parseTime=True&loc=Local"

