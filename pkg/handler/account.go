package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

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
	c.JSON(http.StatusOK, Response{
		Message: fmt.Sprintf("Account with id = %d was created", account.Id),
	})
}
