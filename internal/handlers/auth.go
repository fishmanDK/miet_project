package handlers

import (
	"fmt"
	"net/http"

	"github.com/fishmanDK/miet_project/internal/core"
	"github.com/gin-gonic/gin"
)

func (h *Handlers) RedirectSignIn(c *gin.Context) {
	c.Redirect(http.StatusPermanentRedirect, "/auth/sign-in")
}

func (h *Handlers) signInPage(c *gin.Context) {
	err := h.tmpls.ExecuteTemplate(c.Writer, "sign_in.html", nil)

	if err != nil {
		fmt.Println("can't execute template", err)
		return
	}
}

func (h *Handlers) signIn(c *gin.Context) {
	var input core.Client
	if err := c.BindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	tokens, err := h.service.Auth.Authentication(input)
	if err != nil {
		fmt.Println("can't execute template", err)
		return
	}

	c.JSON(http.StatusOK, tokens)
}

func (h *Handlers) signUpPage(c *gin.Context) {
	err := h.tmpls.ExecuteTemplate(c.Writer, "sign_up.html", nil)

	if err != nil {
		fmt.Println("can't execute template", err)
		return
	}
}

func (h *Handlers) signUp(c *gin.Context) {
	var input core.Client
	if err := c.BindJSON(&input); err != nil {
		fmt.Println(err)
		c.JSON(http.StatusBadRequest, err)
		return
	}

	_, err := h.service.Auth.CreateUser(input)
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusBadRequest, err)
		return
	}

	c.JSON(http.StatusOK, struct{
		Status string `json:"status"`
		Msg string `json:"msg"`
	}{
		Status: "ok",
		Msg: "user success created",
	})
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
