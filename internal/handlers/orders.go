package handlers

import (
	"fmt"
	"net/http"
	"strconv"

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

	for i := range orders{
		orders[i].OrderDate = orders[i].OrderDate[:10]
	}

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
	is_admin, _ := c.Get("is_admin")
	if !is_admin.(bool) {
		c.JSON(http.StatusForbidden, nil)
		return
	}
	
	var input core.DeleteOrder
	if err := c.BindJSON(&input); err != nil {
		fmt.Println(err)
		c.JSON(http.StatusBadRequest, err)
		return
	}

	err := h.service.Orders.DeleteOrder(input.UserID, input.CassetteID)
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


func (h *Handlers) GetOrdersForAdmin(c *gin.Context) {
	is_admin, _ := c.Get("is_admin")
	if !is_admin.(bool) {
		c.JSON(http.StatusForbidden, nil)
		return
	}

	// Получение параметров из query string
	storeID, err := strconv.Atoi(c.Query("store_id"))
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid store_id"})
		return
	}

	cassetteID, err := strconv.Atoi(c.Query("cassette_id"))
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid cassette_id"})
		return
	}

	reservations, err := h.service.Orders.GetOrdersForAdmin(cassetteID, storeID)
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get reservations"})
		return
	}

	for i := range reservations{
		reservations[i].ReservationDate = reservations[i].ReservationDate[:10]
	}

	c.JSON(http.StatusOK, reservations)
}
