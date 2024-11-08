package handlers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handlers) Index(c *gin.Context) {
	_, ok := c.Get("user_id")
	if !ok {
		fmt.Println(11111)
		err := h.tmpls.ExecuteTemplate(c.Writer, "init.html", nil)

		if err != nil {
			fmt.Println("can't execute template", err)
			return
		}
	} else {
		c.Redirect(http.StatusPermanentRedirect, "/store")
	}

}
