package handlers

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func (h *Handlers) Authentication(c *gin.Context) {
	// Попытка получить токен из заголовка Authorization
	header := c.GetHeader("Authorization")
	var token string

	// Если заголовок Authorization пустой, пытаемся найти access_token в cookies
	if header == "" {
		// Ищем access_token в cookies
		cookie, err := c.Cookie("access_token")
		if err != nil {
			// Если токен не найден в cookies, перенаправляем на страницу авторизации
			c.Redirect(http.StatusPermanentRedirect, "/auth/sign-in")
			NewErrorResponse(c, http.StatusUnauthorized, "пользователь не авторизирован")
			return
		}
		token = cookie // Используем токен из cookie
	} else {
		// Разбираем заголовок Authorization
		headerParts := strings.Split(header, " ")
		if len(headerParts) != 2 {
			c.Redirect(http.StatusPermanentRedirect, "/auth/sign-in")
			NewErrorResponse(c, http.StatusUnauthorized, "ошибка авторизации")
			return
		}
		token = headerParts[1] // Используем токен из заголовка
	}

	// Проверяем токен
	res, err := h.service.Auth.ParseToken(token)
	if err != nil {
		c.Redirect(http.StatusPermanentRedirect, "/auth/sign-in")
		NewErrorResponse(c, http.StatusUnauthorized, "ошибка парсинга токена")
		return
	}

	// Устанавливаем user_id в контекст
	c.Set("user_id", res.ID)
	if res.Role == "admin"{
		c.Set("is_admin", true)
	} else {
		c.Set("is_admin", false)
	}
	c.Next()
}
