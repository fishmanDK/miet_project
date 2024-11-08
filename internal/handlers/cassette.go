package handlers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/fishmanDK/miet_project/internal/core"
	"github.com/gin-gonic/gin"
)

func (h *Handlers) GetCassettes(c *gin.Context) {
	res, err := h.service.Cassettes.GetCassettes()
	if err != nil {
		c.JSON(http.StatusBadRequest, nil)
		return
	}

	type cassettes struct {
		cassettes []core.Cassette `json:"cassettes"`
	}

	c.JSON(http.StatusOK, cassettes{res})
}

func (h *Handlers) GetCassetteDetails(c *gin.Context) {
	cassetteID, err := strconv.Atoi(c.Param("id"))
	userID  := c.GetInt("user_id")
	if err != nil {
		fmt.Println(err)
		// Если не удается преобразовать ID, вернем ошибку
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid cassette ID"})
		return
	}

	res, err := h.service.Cassettes.GetCassetteDetails(cassetteID, userID)
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusBadRequest, nil)
		return
	}

	c.JSON(http.StatusOK, res)
}

func (h *Handlers) GetCassette(c *gin.Context) {
	cassetteID, err := strconv.Atoi(c.Param("id"))

	cassette, err := h.service.Cassettes.GetCassette(cassetteID)
	if err != nil{
		fmt.Println(err)

		c.JSON(http.StatusBadRequest, nil)
		return
	}

	cassette.Year = cassette.Year[:10]

	fmt.Println("-------------------------", cassette.Year)

	err = h.tmpls.ExecuteTemplate(c.Writer, "cassette.html", struct {
		Cassette core.Cassette
	}{
		Cassette: cassette,
	})
}

func (h *Handlers) DeleteCassette(c *gin.Context) {
	id := c.Param("id")

	cassetteID, err := strconv.Atoi(id)
	if err != nil {
		fmt.Println(err)
		// Если не удается преобразовать ID, вернем ошибку
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid cassette ID"})
		return
	}

	err = h.service.Cassettes.DeleteCasseteByID(cassetteID)
	if err != nil {
		fmt.Println(err)
		// Если не удается преобразовать ID, вернем ошибку
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid delete cassete by id"})
		return
	}

	c.JSON(http.StatusOK, nil)
}

func (h *Handlers) CreateCassette(c *gin.Context) {
	var input core.CreateCassetteReq

	if err := c.BindJSON(&input); err != nil {
		h.log.Info(err.Error())
		c.JSON(http.StatusBadRequest, err)
		return
	}

	id, err := h.service.Cassettes.CreateCassette(input)
	if err != nil {
		fmt.Println("---------", err)
		c.JSON(http.StatusBadRequest, err)
		return
	}

	c.JSON(http.StatusOK, struct {
		ID int `json:"id"`
	}{ID: id})
}

func (h *Handlers) CreateCassetteAvailability(c *gin.Context) {
	var input core.CassetteAvailability
	if err := c.BindJSON(&input); err != nil {
		h.log.Info(err.Error())
		c.JSON(http.StatusBadRequest, err)
		return
	}

	err := h.service.Cassettes.CreateCassetteAvailability(input)
	if err != nil {
		fmt.Println("---------", err)
		c.JSON(http.StatusBadRequest, err)
		return
	}

	c.JSON(http.StatusOK, nil)
}
