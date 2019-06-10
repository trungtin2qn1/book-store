package Order

import (
	"book-store/model/Order"

	"github.com/gin-gonic/gin"
)

//OrderRequest ...
type OrderRequest struct {
	Note string `json:"note,omitempty"`
}

//CreateOrder ...
func CreateOrder(c *gin.Context) {
	customerID := c.Param("customer_id")
	orderRequest := OrderRequest{}
	err := c.ShouldBindJSON(&orderRequest)
	if err != nil {
		c.JSON(400, gin.H{
			"message": "Data or data type is invalid",
		})
		return
	}
	res, err := Order.CreateOrder(orderRequest.Note, customerID)
	if err != nil {
		c.JSON(400, gin.H{
			"message": err,
		})
		return
	}
	c.JSON(200, res)
}

//GetAllOrders...
func GetAllOrders(c *gin.Context) {
	customerID := c.Param("customer_id")
	res, err := Order.GetAllOrders(customerID)
	if err != nil {
		c.JSON(400, gin.H{
			"message": err,
		})
		return
	}
	c.JSON(200, res)
}
