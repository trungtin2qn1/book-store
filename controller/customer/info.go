package customer

import (
	"book-store/controller/SetUp"
	"book-store/model/customer"
	"fmt"

	"github.com/gin-gonic/gin"
)

//GetCustomerByID ...
func GetCustomerByID(c *gin.Context) {
	customerID := c.Param("customer_id")
	customer, err := customer.GetCustomerByID(customerID)
	if err != nil {
		c.JSON(SetUp.ErrorMap[fmt.Sprintf("%s", err)], gin.H{
			"message": fmt.Sprintf("%s", err),
		})
		return
	}
	c.JSON(200, customer)
}
