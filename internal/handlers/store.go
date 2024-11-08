package handlers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/fishmanDK/miet_project/internal/core"
	"github.com/gin-gonic/gin"
)

func (h *Handlers) GetStores(c *gin.Context) {
	is_admin, _ := c.Get("is_admin")
	stores, err := h.service.Store.GetStores()
	if err != nil {
		fmt.Println("fix mee h.GetStores", err)
	}

	// Рендерим шаблон и передаем данные о магазинах
	err = h.tmpls.ExecuteTemplate(c.Writer, "index.html", struct {
		Stores  []core.Store
		IsAdmin bool
	}{
		Stores:  stores,
		IsAdmin: is_admin.(bool),
	})

	if err != nil {
		fmt.Println("can't execute template", err)
		return
	}
}

func (h *Handlers) GetStore(c *gin.Context) {
	id := c.Param("id")
	user_id, _ := c.Get("user_id")

	is_admin, _ := c.Get("is_admin")

	// Преобразуем ID в нужный тип данных, например, в int
	storeID, err := strconv.Atoi(id)
	if err != nil {
		// Если не удается преобразовать ID, вернем ошибку
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid cassette ID"})
		return
	}
	cassettes, err := h.service.Cassettes.GetCassettesByStoreID(storeID)
	if err != nil {
		fmt.Println("fix mee h.GetStore", err)
	}

	// Рендерим шаблон и передаем данные о магазинах
	err = h.tmpls.ExecuteTemplate(c.Writer, "catalog.html", struct {
		Cassettes []core.Cassette
		StoreID   int
		IsAdmin   bool
		UserID int

	}{
		Cassettes: cassettes,
		StoreID:   storeID,
		IsAdmin:   is_admin.(bool),
		UserID: user_id.(int),

	})

	if err != nil {
		fmt.Println("can't execute template", err)
		return
	}
}

func (h *Handlers) DeleteStore(c *gin.Context) {
	// Логика для удаления кассеты
}

func (h *Handlers) CreateStore(c *gin.Context) {
	var input core.Store
	if err := c.BindJSON(&input); err != nil {
		h.log.Info(err.Error())
		c.JSON(http.StatusBadRequest, err)
		return
	}
	fmt.Println(input)
	id, err := h.service.Store.CreateStore(input)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	input.Id = id
	c.JSON(http.StatusOK, input)
}
