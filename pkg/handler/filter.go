package handler

import (
	"net/http"
	"github.com/gin-gonic/gin"
)



type Interval struct {
	lte uint `json:"lte" binding:"required"`
	gte uint `json:"gte" binding:"required"` 
}

func (h *Handler) filter(c *gin.Context){
	var input Interval

	if err := c.BindJSON(&input); err!= nil {
		newErrorResponse(c,http.StatusBadRequest, err.Error())
		return
	}

	products,err := h.services.Filtering.FilteringProduct(c,input.lte,input.gte)


	if err!= nil {
		newErrorResponse(c,http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK,map[string]interface{}{
		"data":products,
	})
}