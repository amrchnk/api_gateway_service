package handler

import (
	"fmt"
	"github.com/amrchnk/api-gateway/pkg/models"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

// @Summary SignUp
// @Tags auth
// @Description create account
// @ID create-user
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
		UserId:    id,
	})
}

// @Summary SignIn
// @Tags auth
// @Description login
// @ID login
// @Accept  json
// @Produce  json
// @Param input body models.SignInRequest true "credentials"
// @Success 200 {object} models.UserTokens "Success login"
// @Failure 400 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Router /auth/sign-in [post]
func (h *Handler) signIn(c *gin.Context) {
	var request models.SignInRequest

	if err := c.BindJSON(&request); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	user, err := h.Imp.SignIn(c, request.Login, request.Password)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	userTokens, err := h.Imp.CreateTokens(user.Id, user.RoleId)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	at := time.Unix(userTokens.AtExpires, 0)
	rt := time.Unix(userTokens.RtExpires, 0)
	now := time.Now()

	err = h.Imp.SetInCache(c, userTokens.AccessUuid, userTokens.AccessToken, at.Sub(now))
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	err = h.Imp.SetInCache(c, userTokens.RefreshUuid, userTokens.RefreshToken, rt.Sub(now))
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.Header(authorizationHeader, fmt.Sprintf("Bearer %v", userTokens.AccessToken))
	c.JSON(http.StatusOK, userTokens)
}

// @Summary LogOut
// @Tags auth
// @Description logout user
// @ID logout
// @Accept  json
// @Produce  json
// @Param input body models.SignOutRequest true "user tokens"
// @Success 200 {object} Response "Success login"
// @Failure 400 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Router /auth/logout [post]
// @Security ApiKeyAuth
func (h *Handler) logOut(c *gin.Context) {
	var request models.SignOutRequest

	if err := c.BindJSON(&request); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	accessClaims, err := h.Imp.ParseToken(request.AccessToken)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	del, err := h.Imp.DeleteFromCache(c, accessClaims.AccessUuid)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, "access token is invalid")
		return
	}

	refreshClaims, err := h.Imp.ParseRefreshToken(request.RefreshToken)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	del, err = h.Imp.DeleteFromCache(c, refreshClaims.RefreshUuid)
	if err != nil || del == 0 {
		newErrorResponse(c, http.StatusInternalServerError, "invalid refresh token")
		return
	}
	c.Writer.Header().Del(authorizationHeader)

	newResponse(c, http.StatusOK, "User successfully logged out")
}

// @Summary Refresh token
// @Tags auth
// @Description refresh access token
// @ID refresh
// @Accept  json
// @Produce  json
// @Param input body models.RefreshTokenRequest true "user refresh token"
// @Success 200 {object} models.UserTokens "Success refresh"
// @Failure 400 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Router /auth/refresh [post]
// @Security ApiKeyAuth
func (h *Handler) refreshAccessToken(c *gin.Context) {
	var request models.RefreshTokenRequest
	if err := c.BindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	refreshClaims, err := h.Imp.ParseRefreshToken(request.RefreshToken)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	user, err := h.Imp.GetUserById(c, refreshClaims.UserId)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	exist, err := h.Imp.DeleteFromCache(c, refreshClaims.RefreshUuid)
	if err != nil || exist == 0 {
		newErrorResponse(c, http.StatusUnauthorized, "invalid refresh token")
		return
	}
	fmt.Println(exist)

	userTokens, err := h.Imp.CreateTokens(user.Id, user.RoleId)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	at := time.Unix(userTokens.AtExpires, 0)
	rt := time.Unix(userTokens.RtExpires, 0)
	now := time.Now()

	err = h.Imp.SetInCache(c, userTokens.AccessUuid, userTokens.AccessToken, at.Sub(now))
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	err = h.Imp.SetInCache(c, userTokens.RefreshUuid, userTokens.RefreshToken, rt.Sub(now))
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.Header(authorizationHeader, fmt.Sprintf("Bearer %v", userTokens.AccessToken))
	c.JSON(http.StatusOK, userTokens)
}
