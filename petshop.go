package main

import (
  "github.com/ararog/petshop/resources"
  "github.com/ararog/petshop/models"
  "github.com/ararog/petshop/services"
	"github.com/ararog/petshop/config"
  "github.com/appleboy/gin-jwt"
  "github.com/gin-gonic/gin"
	"github.com/BurntSushi/toml"
  "github.com/jinzhu/gorm"
  _ "github.com/jinzhu/gorm/dialects/postgres"
  "time"
)

func GetServerEngine(environment string) *gin.Engine {

	db, err := gorm.Open("postgres", databaseName)
  if err != nil {
    panic("failed to connect database")
  }

  db.AutoMigrate(&models.User{})

  userResource := &resources.UserResource{DB: db}

  authMiddleware := &jwt.GinJWTMiddleware{
    Realm: "test zone",
    Key: []byte("petshop2016"),
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
  v1 := router.Group("/api/v1")
  v1.POST("/auth/signin", authMiddleware.LoginHandler)

  users := v1.Group("/users")
  users.Use(authMiddleware.MiddlewareFunc())
  {
      users.GET("/me", userResource.User)
  }

	return router
}

func main() {
  GetServerEngine("production").Run(":8080")
}
