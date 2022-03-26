package handler

import (
	"encoding/json"
	"github.com/amrchnk/api-gateway/pkg/models"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *Handler) signUp(c *gin.Context) {
	var input models.User

	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid input body")
		return
	}

	id, err := h.authService.SignUp(c,input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	input.Id = id
	userSession, err := json.Marshal(input)
	if err != nil {
		return
	}
	session := sessions.Default(c)
	session.Set("UserSession",userSession)
	session.Save()

	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

func (h *Handler) signIn(c *gin.Context) {
	var input models.User

	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	userSession, err := h.authService.SignIn(c,input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	session := sessions.Default(c)
	session.Set("UserSession", userSession)
	session.Save()

	c.JSON(http.StatusOK, map[string]interface{}{
		"session": userSession,
	})
}
