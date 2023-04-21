package handler

import (
	"fmt"
	"net/http"
	"strconv"
	"github.com/gin-gonic/gin"
)



type Interval struct {
	lte uint `json:"lte" binding:"required"`
	gte uint `json:"gte" binding:"required"` 
}

func (h *Handler) filter(c *gin.Context){
	// var input Interval
	
	// if err := c.BindJSON(&input); err!= nil {
	// 		newErrorResponse(c,http.StatusBadRequest, err.Error())
	// 		return
	// }

	// fmt.Println(input.gte)

	ltestr := c.Query("lte")
	gtestr := c.Query("gte")

	lte64, err := strconv.ParseUint(ltestr, 10, 32)

	gte64, err := strconv.ParseUint(gtestr, 10, 32)
    if err != nil {
        fmt.Println(err)
    }
    lte := int(lte64)
    gte := int(gte64)

	products,err := h.services.Filtering.FilteringProduct(c,lte,gte)

	if err!= nil {
		// fmt.Println("s")
		newErrorResponse(c,http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK,map[string]interface{}{
		"data":products,
	})
}