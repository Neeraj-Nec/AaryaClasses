package userM

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"arooClasses/pkg/comman/userAuthModels"
)

type UserManagement struct {
}

func (usm *UserManagement) LoginHandler(c *gin.Context) {
	login := userAuthModels.Login{}
	if err := c.BindJSON(&login); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	
	c.JSON(http.StatusOK, gin.H{"message": "Data received successfully"})
}

func (usm *UserManagement) Register(c *gin.Context) {
	c.JSON(http.StatusOK, "")
}
