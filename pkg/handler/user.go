package handler

import (
	"github.com/amrchnk/api-gateway/pkg/models"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// @Summary Get User By Id
// @Tags user
// @Description get user by id
// @ID get-user-by-id
// @Accept  json
// @Produce  json
// @Success 200 {object} models.User
// @Failure 400 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /users/:id [get]
func (h *Handler) getUserById(c *gin.Context) {

	userId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid list id param")
		return
	}

	user, err := h.Imp.GetUserById(c,int64(userId))
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, user)
}

// @Summary Delete User By Id
// @Tags user
// @Description delete user by id
// @ID delete-user-by-id
// @Accept  json
// @Produce  json
// @Success 200 {string} string "message"
// @Failure 400 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /users/:id [delete]
func (h *Handler) deleteUserById(c *gin.Context) {

	userId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid list id param")
		return
	}

	msg, err := h.Imp.DeleteUserById(c,int64(userId))
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, msg)
}

type getAllUsersResponse struct {
	Data []models.User `json:"data"`
}

// @Summary Get all users
// @Tags user
// @Description get all users
// @ID get-all-users
// @Accept  json
// @Produce  json
// @Success 200 {string} getAllUsersResponse
// @Failure 400 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /users/ [get]
func (h *Handler) getAllUsers(c *gin.Context){
	users,err:=h.Imp.GetAllUsers(c)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK,getAllUsersResponse{
		Data: users,
	})
}