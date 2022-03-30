package handler

import (
	"encoding/json"
	"github.com/amrchnk/api-gateway/pkg/models"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"net/http"
)

// @Summary SignUp
// @Tags auth
// @Description create account
// @ID create-account
// @Accept  json
// @Produce  json
// @Param input body models.User true "account info"
// @Success 200 {integer} integer 1
// @Failure 400 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Router /auth/sign-up [post]
func (h *Handler) signUp(c *gin.Context) {
	var input models.User

	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid input body")
		return
	}

	id, err := h.Imp.SignUp(c,input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	input.Id = id
	userSession, err := json.Marshal(input)
	if err != nil {
		return
	}
	c.SetCookie("UserSession", string(userSession), 100, "/", "localhost:8000", true, false)
	session := sessions.Default(c)
	session.Set("UserSession",userSession)
	session.Save()

	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

type signInInput struct {
	Login string `json:"login" binding:"required"`
	Password string `json:"password" binding:"required"`
}

// @Summary SignIn
// @Tags auth
// @Description login
// @ID login
// @Accept  json
// @Produce  json
// @Param input body signInInput true "credentials"
// @Success 200 {string} string "token"
// @Failure 400 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Router /auth/sign-in [post]
func (h *Handler) signIn(c *gin.Context) {
	var input signInInput

	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	userSession, err := h.Imp.SignIn(c,input.Login,input.Password)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.SetCookie("userSession", string(userSession), 100, "/", "localhost", true, false)

	session := sessions.Default(c)
	session.Set("userSession", userSession)
	session.Save()

	c.JSON(http.StatusOK, map[string]interface{}{
		"session": userSession,
	})
}
