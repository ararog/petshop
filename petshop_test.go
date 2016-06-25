package main

import (
	"os"
	"fmt"
	"net/http"
	"github.com/ararog/petshop/application"
  "github.com/ararog/petshop/models"
	"github.com/appleboy/gofight"
	"github.com/stretchr/testify/assert"
	"github.com/buger/jsonparser"
	"github.com/jinzhu/gorm"
 	"testing"
)

func TestMain(m *testing.M) {
	environment := os.Getenv("PETSHOP_ENV")
	if environment == "" {
		os.Setenv("PETSHOP_ENV", "test")
	}

	config := application.LoadConfig()

	db, err := gorm.Open(config.DB.Type, config.DB.ConnectionString)
	if err != nil {
		panic("failed to connect database")
	}

	db.AutoMigrate(&models.User{})
	db.Exec("DELETE from users;")
	db.Create(&models.User{Name: "Rogerio Araujo", Email: "rogerio.araujo@gmail.com", Password: "123456"})
	code := m.Run()
	os.Exit(code)
}

func TestLogin(t *testing.T) {
	r := gofight.New()

	config := application.LoadConfig()

  r.POST("/api/v1/auth/signin").
    SetJSON(gofight.D{
			"username": "rogerio.araujo@gmail.com",
		  "password": "123456",
    }).
    Run(GetServerEngine(config), func(r gofight.HTTPResponse, rq gofight.HTTPRequest) {
      assert.Equal(t, http.StatusOK, r.Code)
    })
}

func TestMe(t *testing.T) {
	r := gofight.New()

	config := application.LoadConfig()

	r.POST("/api/v1/auth/signin").
    SetJSON(gofight.D{
			"username": "rogerio.araujo@gmail.com",
		  "password": "123456",
    }).
    Run(GetServerEngine(config), func(res gofight.HTTPResponse, rq gofight.HTTPRequest) {
			data := []byte(res.Body.String())
			token, _ := jsonparser.GetString(data, "token")

			r.GET("/api/v1/users/me").
				SetHeader(gofight.H{
			    "Authorization": fmt.Sprintf("Bearer %s", token),
			  }).
		    Run(GetServerEngine(config), func(r gofight.HTTPResponse, rq gofight.HTTPRequest) {
		      assert.Equal(t, http.StatusOK, r.Code)
		    })
    })
}
