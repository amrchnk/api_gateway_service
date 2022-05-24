package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

const (
	authorizationHeader = "Authorization"
	userCtx             = "userId"
	accountCtx          = "accountId"
	roleCtx             = "roleId"
)

func (h *Handler) userIdentity(c *gin.Context) {
	header := c.GetHeader(authorizationHeader) //получаем токен из заголовка
	if header == "" {
		newErrorResponse(c, http.StatusUnauthorized, "empty auth header")
		return
	}

	//проверяем токен на валидность
	headerParts := strings.Split(header, " ")
	if len(headerParts) != 2 || headerParts[0] != "Bearer" {
		newErrorResponse(c, http.StatusUnauthorized, "invalid auth header")
		return
	}

	if len(headerParts[1]) == 0 {
		newErrorResponse(c, http.StatusUnauthorized, "token is empty")
		return
	}

	claims, err := h.Imp.ParseToken(headerParts[1])
	if err != nil {
		newErrorResponse(c, http.StatusUnauthorized, err.Error())
		return
	}

	//проверяем существование токена в кэше
	_, err = h.Imp.GetFromCache(c, claims.AccessUuid)
	if err != nil {
		newErrorResponse(c, http.StatusUnauthorized, "token doesn't exist")
		return
	}

	//устанавливаем данные из payload в контекст запроса
	c.Set(userCtx, claims.UserId)
	c.Set(roleCtx, claims.RoleId)
}

func (h *Handler) adminIdentity(c *gin.Context) {
	id, ok := c.Get(roleCtx)
	if !ok {
		newErrorResponse(c, http.StatusBadRequest, "role id not found")
		return
	}

	idInt, ok := id.(int64)
	if !ok {
		newErrorResponse(c, http.StatusBadRequest, "role id is invalid")
		return
	}

	if idInt != 1 {
		newErrorResponse(c, http.StatusForbidden, "access denied for this role")
		return
	}
}
