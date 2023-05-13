package routes

import (
	"github.com/gin-gonic/gin"
	inventoryController "github.com/micro/outstream-agg/api/controllers/inventory"
)

func InventoryRoutes(router *gin.RouterGroup, controller *inventoryController.Controller) {
	routerInv := router.Group("/inventory")

	routerInv.GET("/", controller.Index)
	routerInv.POST("/", controller.Inventoryinfo)
}
