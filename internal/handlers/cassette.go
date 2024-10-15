package handlers

import (
	"fmt"
	"net/http"

	"github.com/fishmanDK/miet_project/internal/core"
	"github.com/gin-gonic/gin"
)


func (h *Handlers) GetCassettes(c *gin.Context) {
    res, err := h.service.Cassettes.GetCassettes()
	if err != nil{
		c.JSON(http.StatusBadRequest, nil)
		return
	}

	type cassettes struct{
		cassettes []core.Cassette `json:"cassettes"`
	}

	c.JSON(http.StatusOK, cassettes{res})
}

func (h *Handlers) GetCassette(c *gin.Context) {
}


func (h *Handlers) DeleteCassette(c *gin.Context) {
    // Логика для удаления кассеты
}

func (h *Handlers) CreateCassette(c *gin.Context) {
	var input core.Cassette
	if err := c.BindJSON(&input); err != nil{
		h.log.Info(err.Error())
		c.JSON(http.StatusBadRequest, err)
		return 
	}

    id, err := h.service.Cassettes.CreateCassette(input)
	if err != nil{
		fmt.Println("---------", err)
		c.JSON(http.StatusBadRequest, err)
		return
	}

	c.JSON(http.StatusOK, id)
}

func (h *Handlers) CreateCassetteAvailability(c *gin.Context) {
	var input core.CassetteAvailability
	if err := c.BindJSON(&input); err != nil{
		h.log.Info(err.Error())
		c.JSON(http.StatusBadRequest, err)
		return 
	}

    err := h.service.Cassettes.CreateCassetteAvailability(input)
	if err != nil{
		fmt.Println("---------", err)
		c.JSON(http.StatusBadRequest, err)
		return
	}

	c.JSON(http.StatusOK, nil)
}