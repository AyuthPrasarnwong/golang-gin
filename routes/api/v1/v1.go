package v1

import (
	"app/controllers/api/v1"
	//"app/middlewares/jwt"
	"github.com/gin-gonic/gin"
)

func Setup(r *gin.Engine) {
	apiV1 := r.Group("v1")

	apiV1.GET("/orders",  v1.GetAllOrders)
	apiV1.GET("/orders/:id",  v1.FindOrderByID)
	//apiV1.POST("/checkout",  v1.CheckoutOrder)
	apiV1.GET("/terminals/:id",  v1.FindTerminalByID)
}