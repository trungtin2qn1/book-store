package Customer

import (
	"book-store/controller/SetUp"
	"book-store/database"
	"book-store/model/Customer"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

//GetCustomerByID ...
func GetCustomerByID(c *gin.Context) {
	customerID := c.Param("customer_id")
	customer, err := Customer.GetCustomerByID(customerID)
	if err != nil {
		c.JSON(SetUp.ErrorMap[fmt.Sprintf("%s", err)], gin.H{
			"message": fmt.Sprintf("%s", err),
		})
		return
	}
	c.JSON(200, customer)
}

//UpdateCustomerInfo
func UpdateCustomerInfo(c *gin.Context) {
	customerID := c.Param("customer_id")
	newCustomer := database.Customer{}
	err := c.ShouldBindJSON(&newCustomer)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Data or data type is invalid",
		})
		return
	}
	customer, err := Customer.GetCustomerByID(customerID)
	if err != nil {
		c.JSON(SetUp.ErrorMap[fmt.Sprintf("%s", err)], gin.H{
			"message": fmt.Sprintf("%s", err),
		})
		return
	}
	err = Customer.UpdateCustomerInfo(&customer, &newCustomer)
	if err != nil {
		c.JSON(SetUp.ErrorMap[fmt.Sprintf("%s", err)], gin.H{
			"message": fmt.Sprintf("%s", err),
		})
		return
	}
	newCustomer.ID = customer.ID
	c.JSON(200, newCustomer)
}
