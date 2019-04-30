package Authentication

import (
	"book-store/controller/SetUp"
	"book-store/model/Authentication"
	"donz-app-backend/model"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

//SignUp ...
func SignUp(c *gin.Context) {
	var authReq model.Authentication
	err := c.ShouldBind(&authReq)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Data or data type is invalid",
		})
		return
	}

	customer, err := Authentication.SignUp(authReq.Email, authReq.Password)

	if err != nil {
		c.JSON(SetUp.ErrorMap[fmt.Sprintf("%s", err)], gin.H{
			"message": fmt.Sprintf("%s", err),
		})
		return
	}

	customer.Password = ""
	c.JSON(200, customer)
}
