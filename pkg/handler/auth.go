package handler

import (
	"fmt"
	"github.com/amrchnk/api-gateway/pkg/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

// @Summary SignUp
// @Tags auth
// @Description create account
// @ID create-account
// @Accept  json
// @Produce  json
// @Param input body models.SignUpRequest true "account info"
// @Success 200 {object} models.SignUpResponse
// @Failure 400 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Router /auth/sign-up [post]
func (h *Handler) signUp(c *gin.Context) {
	var request models.SignUpRequest

	if err := c.BindJSON(&request); err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid input body")
		return
	}

	user := models.User{
		Login:    request.Login,
		Password: request.Password,
		Username: request.Username,
	}
	id, err := h.Imp.SignUp(c, user)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	accountId, err := h.Imp.CreateAccount(c, id)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, fmt.Sprintf("Error while creating user account: %v", err.Error()))
		return
	}

	c.JSON(http.StatusOK, models.SignUpResponse{
		AccountId: accountId,
		UserId: id,
	})
}

// @Summary SignIn
// @Tags auth
// @Description login
// @ID login
// @Accept  json
// @Produce  json
// @Param input body models.SignInRequest true "credentials"
// @Success 200 {object} models.SignInResponse "Success login"
// @Failure 400 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Router /auth/sign-in [post]
func (h *Handler) signIn(c *gin.Context) {
	var request models.SignInRequest

	if err := c.BindJSON(&request); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	token, err := h.Imp.SignIn(c, request.Login, request.Password)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.Header(authorizationHeader, fmt.Sprintf("Bearer %v", token))
	c.JSON(http.StatusOK, models.SignInResponse{
		Token: token,
	})
}
