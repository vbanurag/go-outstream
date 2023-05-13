package routes

import (
	"github.com/gin-gonic/gin"
	inventoryController "github.com/micro/outstream-agg/api/controllers/inventory"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func ApplicationV1Router(router *gin.Engine) {
	router_v1 := router.Group("/api/v1")
	{
		{
			router_v1.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
		}
		InventoryRoutes(router_v1, &inventoryController.Controller{})
	}
}
