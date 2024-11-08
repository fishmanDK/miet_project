package handlers

import (
	"fmt"
	"net/http"

	"github.com/fishmanDK/miet_project/internal/core"
	"github.com/gin-gonic/gin"
)

func (h *Handlers) GetOrders(c *gin.Context) {
	userID := c.GetInt("user_id")
	if userID == 0{
		fmt.Println("user_id = 0")
		return 
	}

	orders, err := h.service.Orders.GetUserOrders(userID)
	if err != nil{
		fmt.Println(err)

		c.JSON(http.StatusBadRequest, nil)
		return
	}

	fmt.Println(orders)

	err = h.tmpls.ExecuteTemplate(c.Writer, "orders.html", struct {
		Orders []core.Order
	}{
		Orders: orders,
	})

}

func (h *Handlers) CreateOrder(c *gin.Context) {
	var input core.Order
	if err := c.BindJSON(&input); err != nil {
		fmt.Println(err)
		c.JSON(http.StatusBadRequest, err)
		return
	}

	id, err := h.service.Orders.CreateOrder(input)
	if err != nil {
		fmt.Println(err)
		return
	}

	c.JSON(http.StatusOK, struct {
		Status string `json:"status"`
		ID     int    `json:"id"`
	}{
		Status: "ok",
		ID:     id,
	})
}

func (h *Handlers) DeleteOrder(c *gin.Context) {
	var input core.DeleteOrder
	if err := c.BindJSON(&input); err != nil {
		fmt.Println(err)
		c.JSON(http.StatusBadRequest, err)
		return
	}

	err := h.service.Orders.DeleteOrder(input)
	if err != nil {
		fmt.Println(err)
		return
	}

	c.JSON(http.StatusOK, struct {
		Status string `json:"status"`
	}{
		Status: "ok",
	})
}
