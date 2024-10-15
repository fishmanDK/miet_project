package handlers

import (
	"net/http"

	"github.com/fishmanDK/miet_project/internal/core"
	"github.com/gin-gonic/gin"
)

func (h *Handlers) signIn(c *gin.Context) {
	
}

func (h *Handlers) signUp(c *gin.Context) {
	var input core.Client
	if err := c.BindJSON(&input); err != nil{
		c.JSON(http.StatusBadRequest, err)
		return 
	}

	id, err := h.service.Clients.CreateClient(input)
	if err != nil{
		c.JSON(http.StatusBadRequest, err)
		return 
	}

	input.Id = id
	c.JSON(http.StatusOK, input)
}

func (h *Handlers) GetUser(c *gin.Context) {
    // Логика для получения информации о пользователе
}

func (h *Handlers) DeleteUser(c *gin.Context) {
    // Логика для удаления пользователя
}

func (h *Handlers) UpdateUser(c *gin.Context) {
    // Логика для обновления информации о пользователе
}


