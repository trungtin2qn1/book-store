package Cart

import (
	"book-store/database"
	"book-store/model/Cart"

	"github.com/gin-gonic/gin"
)

//DeleteCart ...
func DeleteCart(c *gin.Context) {
	customerID := c.Param("customer_id")
	cartID := c.Param("cart_id")
	err := Cart.DeleteProductFromCart(customerID, cartID)
	if err != nil {
		c.JSON(406, gin.H{
			"message": "Check input again",
		})
		return
	}

	c.JSON(200, gin.H{
		"message": "Success",
	})
}

//AddBookToCart ...
func AddBookToCart(c *gin.Context) {
	customerID := c.Param("customer_id")
	cart := database.Cart{}
	err := c.ShouldBindJSON(&cart)
	if err != nil {
		c.JSON(400, gin.H{
			"message": "Data or data type is invalid",
		})
		return
	}
	res, err := Cart.AddBookToCart(customerID, cart)
	if err != nil {
		c.JSON(406, gin.H{
			"message": "Check input again",
		})
		return
	}

	c.JSON(200, res)
}

//UpdateCart ...
func UpdateCart(c *gin.Context) {
	customerID := c.Param("customer_id")
	cartID := c.Param("cart_id")
	cart := database.Cart{}
	err := c.ShouldBindJSON(&cart)
	if err != nil {
		c.JSON(400, gin.H{
			"message": "Data or data type is invalid",
		})
		return
	}
	res, err := Cart.UpdateCartInfo(customerID, cartID, cart)
	if err != nil {
		c.JSON(406, gin.H{
			"message": "Check input again",
		})
		return
	}

	c.JSON(200, res)
}

//GetAllCartsInfo ...
func GetAllCartsInfo(c *gin.Context) {
	customerID := c.Param("customer_id")
	carts, _ := Cart.GetAllCartsInfo(customerID)
	c.JSON(200, carts)
}
