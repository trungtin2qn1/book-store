package customer

import (
	"book-store/controller/SetUp"
	"book-store/model/customer"
	"book-store/utils"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

//GetCustomerByID ...
func GetCustomerByID(c *gin.Context) {
	interfaceCustomerID, e := c.Get("user_id")
	if e != true {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "Unauthorized",
		})
		return
	}
	customerID := utils.InterfaceToString(interfaceCustomerID)
	customer, err := customer.GetCustomerByID(customerID)
	if err != nil {
		c.JSON(SetUp.ErrorMap[fmt.Sprintf("%s", err)], gin.H{
			"message": fmt.Sprintf("%s", err),
		})
		return
	}
	c.JSON(200, customer)
}
