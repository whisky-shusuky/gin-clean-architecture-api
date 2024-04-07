package router

import (
	"gin-clean-architecture-api/pkg/interfaces/handler"

	"github.com/gin-gonic/gin"
)

func SetupRouter(shopHandler handler.ShopHandler) *gin.Engine {
	router := gin.Default()

	v1 := router.Group("/api/v1")
	{
		shops := v1.Group("/shops")
		{
			shops.GET("", shopHandler.GetShops)
			shops.GET("/:id", shopHandler.GetShop)
			shops.POST("", shopHandler.CreateShop)
			shops.PUT("/:id", shopHandler.UpdateShop)
			shops.DELETE("/:id", shopHandler.DeleteShop)
		}
	}

	return router
}
