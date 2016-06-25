package main

import (
	"fmt"
	"time"
	"github.com/ararog/petshop/application"
	"github.com/ararog/petshop/resources"
	"github.com/ararog/petshop/services"
	"github.com/appleboy/gin-jwt"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func GetServerEngine(environment string, config application.Config) *gin.Engine {

	db, err := gorm.Open(config.DB.Type, config.DB.ConnectionString)
  if err != nil {
    panic("failed to connect database")
  }

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
	environment := application.Environment()
	config := application.LoadConfig()
	GetServerEngine(environment, config).Run(fmt.Sprintf(":%d", config.Server.Port))
}
