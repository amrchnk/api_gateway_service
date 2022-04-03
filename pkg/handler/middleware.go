package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

const (
	authorizationHeader = "Authorization"
	userCtx             = "userId"
	roleCtx             = "roleId"
)

func (h *Handler) userIdentity(c *gin.Context) {
	header := c.GetHeader(authorizationHeader)
	if header == "" {
		newErrorResponse(c, http.StatusUnauthorized, "empty auth header")
		return
	}

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

	if err!=nil{
		return
	}

	c.Set(userCtx, claims.UserId)
	c.Set(roleCtx, claims.RoleId)
}

func (h *Handler)AdminIdentity(c *gin.Context){
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

	if idInt!=1 {
		newErrorResponse(c, http.StatusForbidden, "access denied for this role")
		return
	}

}