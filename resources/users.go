package resources

import (
	"github.com/gin-gonic/gin"
    "github.com/jinzhu/gorm"
    "github.com/ararog/petshop/models"
)

type UserResource struct {
	DB *gorm.DB
}

func (u *UserResource) User(c *gin.Context) {

	user, ok := c.MustGet("user").(models.User)
	if ok {
    	c.JSON(200, user)
	} else {
		json := gin.H{ "message": "User not found" }
		c.JSON(404, json)
	}
}
