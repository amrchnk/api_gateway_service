package handler

import (
	"fmt"
	"github.com/amrchnk/api-gateway/pkg/models"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// @Summary Get account info
// @Tags account
// @Description get account info by user id
// @ID get-account
// @Accept  json
// @Produce  json
// @Param id   path int64  true  "User ID"
// @Success 200 {object} models.Account
// @Failure 400 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Router /account/:id [get]
func (h *Handler) getAccountByUserId(c *gin.Context) {
	userId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid account id param")
		return
	}

	account, err := h.Imp.GetAccountByUserId(c, int64(userId))
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, fmt.Sprintf("Error creating account: %v", err))
		return
	}

	c.Set(accountCtx, account.Id)
	c.JSON(http.StatusOK, models.Account{
		Id:        account.Id,
		UserId:    account.UserId,
		CreatedAt: account.CreatedAt,
	})
}

// @Summary Create account by user id
// @Tags account
// @Description create default account with user id from body
// @ID create-account
// @Accept  json
// @Produce  json
// @Param id   path int64  true  "User ID"
// @Success 200 {object} Response
// @Failure 400 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Router /account/:id [put]
func (h *Handler) createAccountByUserId(c *gin.Context) {
	userId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid account id param")
		return
	}

	accountId, err := h.Imp.CreateAccount(c, int64(userId))
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, fmt.Sprintf("Error creating account: %v", err.Error()))
		return
	}

	newResponse(c, http.StatusOK, fmt.Sprintf("Account with id = %d was created for user with id =%d", accountId, userId))
}