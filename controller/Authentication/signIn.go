package Authentication

import (
	"book-store/controller/SetUp"
	"book-store/model/Authentication"
	"book-store/model/Jwt"
	"donz-app-backend/model"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

//SignIn ...
func SignIn(c *gin.Context) {
	var authReq model.Authentication
	err := c.ShouldBind(&authReq)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Data or data type is invalid",
		})
		return
	}

	customer, err := Authentication.SignIn(authReq.Email, authReq.Password)

	if err != nil {
		c.JSON(SetUp.ErrorMap[fmt.Sprintf("%s", err)], gin.H{
			"message": fmt.Sprintf("%s", err),
		})
		return
	}

	token, err := Jwt.IssueToken(customer.ID.Hex(), customer.Email)
	customer.Token = token
	c.JSON(200, customer)
}
