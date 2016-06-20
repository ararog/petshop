package main

import (
	"github.com/jinzhu/gorm"
  "github.com/ararog/petshop/models"
)

func seed() {
	db, err := gorm.Open("sqlite3", "production.db")
  if err != nil {
    panic("failed to connect database")
  }

  db.AutoMigrate(&models.User{})
	db.Create(&models.User{Name: "Rogerio Araujo", Email: "rogerio.araujo@gmail.com", Password: "123456"})
}
