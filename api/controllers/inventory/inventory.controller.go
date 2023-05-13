package inventory

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/micro/outstream-agg/api/controllers"
	"net/http"
)

type Controller struct {
}

func (c *Controller) Index(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, "success")
}

func (c *Controller) Inventoryinfo(ctx *gin.Context) {

	var requestMap map[string]interface{}
	err := controllers.BindJSONMap(ctx, &requestMap)
	if err != nil {
		return
	}
	err = inventoryInfo(requestMap)
	if err != nil {
		_ = ctx.Error(err)
		return
	}
	fmt.Print(requestMap)

	ctx.JSON(http.StatusOK, requestMap)
}
