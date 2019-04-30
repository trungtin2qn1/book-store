package app

import (
	"book-store/controller/Authentication"
	"book-store/controller/SetUp"
	"book-store/controller/customer"
	"book-store/database"
	"log"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

//SetupRouter ...
func SetupRouter() *gin.Engine {
	router := gin.Default()

	config := cors.DefaultConfig()
	config.AllowAllOrigins = true

	router.Use(cors.New(config))

	api := router.Group("/api")

	api.POST("/sign_up", Authentication.SignUp)

	auth := api.Group("/auth")
	{
		auth.GET("/customer/:customer_id", customer.GetCustomerByID)
	}

	router.NoRoute(func(c *gin.Context) {
		c.JSON(404, gin.H{"code": "404", "message": "Page not found"})
	})

	return router
}

//AppInit App Handler Start From Here
func AppInit() {
	SetUp.InitStuffs()
	database.InitMongo()
	router := SetupRouter()
	if err := router.Run(":2013"); err != nil {
		log.Fatal(err)
	}

}
