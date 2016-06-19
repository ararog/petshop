package main

import (
  "github.com/ararog/petshop/resources"
  "github.com/ararog/petshop/models"
  "github.com/ararog/petshop/services"
  "github.com/appleboy/gin-jwt"
  "github.com/gin-gonic/gin"
  "github.com/jinzhu/gorm"
  _ "github.com/jinzhu/gorm/dialects/sqlite"
  "time"
)

func main() {

  db, err := gorm.Open("sqlite3", "test.db")
  if err != nil {
    panic("failed to connect database")
  }

  db.AutoMigrate(&models.User{})
  db.Create(&models.User{Name: "Rogerio Araujo", Email: "rogerio.araujo@gmail.com", Password: ""})

  userResource := &resources.UserResource{DB: db}

  authMiddleware := &jwt.GinJWTMiddleware{
    Realm: "test zone",
    Key: []byte("petshop42016"),
    Timeout: time.Hour,
    MaxRefresh: time.Hour * 24,
    Authenticator: func(userId string, password string, c *gin.Context) (string, bool) {
      err, _ := services.Login(db, userId, password)
      if err == nil {
        return userId, true
      }

      return userId, false
    },
    Authorizator: func(userId string, c *gin.Context) bool {
      err, user := services.GetUserByEmail(db, userId)
      if err == nil {
        c.Set("user", user)
        return true
      }

      return false
    },
  }

  router := gin.Default()
  router.Use(gin.Logger())
  v1 := router.Group("/v1")
  v1.POST("/auth/login", authMiddleware.LoginHandler)

  users := v1.Group("/users")
  users.Use(authMiddleware.MiddlewareFunc())
  {
      users.GET("/me", userResource.User)
  }

  router.Run(":8080")
}
