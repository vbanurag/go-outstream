package inventory

type InvenmtroyInfo struct {
	sku_id         int    `json:"sku_id" example:"3453" binding:"required"`
	vendor_id      string `json:"vendor_id" example:"123" binding:"required"`
	qty            int    `json:"qty" example:"23" binding:"required"`
	operation_type string `json:"type" example:"BLOCK" binding:"required"`
}
