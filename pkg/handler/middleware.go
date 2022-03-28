package handler

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *Handler) userIdentity(c *gin.Context) {
	session := sessions.Default(c)
	id := session.Get("id")
	if id == "" {
		newErrorResponse(c, http.StatusUnauthorized, "Empty auth cookie")
		return
	}
}