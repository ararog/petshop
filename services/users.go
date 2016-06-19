package services

import (
  "github.com/jinzhu/gorm"
  "github.com/ararog/petshop/models"
)

func Login(db *gorm.DB, email, password string) (error, models.User) {
  var user models.User
  err := db.First(&user, "email = ? AND password = ?", email, password).Error
  return err, user
}

func GetUserByEmail(db *gorm.DB, email string) (error, models.User) {
  var user models.User
  err := db.First(&user, "email = ?", email).Error
  return err, user
}

func GetUserById(db *gorm.DB, id int64) (error, models.User) {
  var user models.User
  err := db.First(&user, "id = ?", id).Error
  return err, user
}
