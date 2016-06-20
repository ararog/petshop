package petshop

import (
	"github.com/appleboy/gofight"
	"github.com/stretchr/testify/assert"
	"github.com/buger/jsonparser"
	"net/http"
 	"testing"
	"fmt"
)

func TestLogin(t *testing.T) {
	r := gofight.New()

  r.POST("/api/v1/auth/signin").
    SetJSON(gofight.D{
			"username": "rogerio.araujo@gmail.com",
		  "password": "1978@rpa",
    }).
    Run(GetServerEngine("test"), func(r gofight.HTTPResponse, rq gofight.HTTPRequest) {
      assert.Equal(t, http.StatusOK, r.Code)
    })
}

func TestMe(t *testing.T) {
	r := gofight.New()

	r.POST("/api/v1/auth/signin").
    SetJSON(gofight.D{
			"username": "rogerio.araujo@gmail.com",
		  "password": "1978@rpa",
    }).
    Run(GetServerEngine("test"), func(res gofight.HTTPResponse, rq gofight.HTTPRequest) {
			data := []byte(res.Body.String())
			token, _ := jsonparser.GetString(data, "token")

			r.GET("/api/v1/users/me").
				SetHeader(gofight.H{
			    "Authorization": fmt.Sprintf("Bearer %s", token),
			  }).
		    Run(GetServerEngine("test"), func(r gofight.HTTPResponse, rq gofight.HTTPRequest) {
		      assert.Equal(t, http.StatusOK, r.Code)
		    })
    })
}
