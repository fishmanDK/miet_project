package handlers

import (
	"net/http"

	"github.com/fishmanDK/miet_project/internal/core"
	"github.com/gin-gonic/gin"
)

func (h *Handlers) GetStore(c *gin.Context) {
    // Логика для получения списка кассет
}

func (h *Handlers) DeleteStore(c *gin.Context) {
    // Логика для удаления кассеты
}

func (h *Handlers) CreateStore(c *gin.Context) {
    var input core.Store
	if err := c.BindJSON(&input); err != nil{
		h.log.Info(err.Error())
		c.JSON(http.StatusBadRequest, err)
		return 
	}
    id, err := h.service.Store.CreateStore(input)
	if err != nil{
		c.JSON(http.StatusBadRequest, err)
		return
	}

	input.Id = id
	c.JSON(http.StatusOK, input)
}