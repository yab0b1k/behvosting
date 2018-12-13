package lib

import (
	"behvosting/models"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/beego/bee/logger"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"time"
)

var db *gorm.DB

func init() {
	var err error
	dbHost := beego.AppConfig.String("dbHost")
	dbPort := beego.AppConfig.String("dbPort")
	dbUser := beego.AppConfig.String("dbUser")
	dbName := beego.AppConfig.String("dbName")
	dbPass := beego.AppConfig.String("dbPass")
	dbConnString := fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=disable password=%s", dbHost, dbPort, dbUser, dbName, dbPass)
	db, err = gorm.Open("postgres", dbConnString)
	if err != nil {
		beeLogger.Log.Fatal("Error connect to pg db " + err.Error())
	}
	db.DB().SetMaxIdleConns(10)
	db.DB().SetMaxOpenConns(100)
	db.LogMode(true)
	db.AutoMigrate(&models.User{}, &models.Image{})
}

func GetUserImages(userName string) []models.Image {
	var user models.User
	found := db.Where("name = ?", userName).Preload("Images").First(&user).RowsAffected
	if found == 0 {
		beeLogger.Log.Error("user " + userName + " not found")
		return nil
	}
	return user.Images
}

func CreateNewUser(userName string, pass string) models.User {
	user := models.User{
		Name: userName,
		Pass: pass,
	}
	err := db.Where("name = ?", userName).FirstOrCreate(&user).Error
	if err != nil {
		beeLogger.Log.Errorf("%+v", err)
	}
	return user
}

func CheckUserPassword(userName string, pass string) bool {
	return db.Model(&models.User{}).Where("name = ? and pass = ?", userName, pass).Select("name").RowsAffected > 0
}

func GetUserInfo(userName string) models.User {
	user := models.User{}
	db.Where("name = ?", userName).First(&user)
	return user
}

func AddImageToUser(userName string, link string) {
	var user models.User
	found := db.Where("name = ?", userName).Preload("Images").First(&user).RowsAffected
	if found == 0 {
		beeLogger.Log.Error("user " + userName + "not found")
		return
	}
	image := models.Image{UserID: user.ID, ImageID: link, CreateTime: time.Now()}
	db.Save(&image)
}
