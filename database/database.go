package database

import (
	"fiber-gorm/models/orderModel"
	"fiber-gorm/models/productModel"
	"fiber-gorm/models/userModel"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

type DbInstance struct {
	Db *gorm.DB
}

var Database DbInstance

func ConnectDb() {
	dsn := "root:@tcp(127.0.0.1:3306)/fiber_gorm?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to the database")
	}
	// migration tables
	_ = db.AutoMigrate(userModel.Users{}, productModel.Products{}, orderModel.Orders{})
	Database = DbInstance{
		Db: db,
	}
}
