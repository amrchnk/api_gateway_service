package handler

import (
	"github.com/amrchnk/api-gateway/pkg/models"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// @Summary GetFromCache User By Id
// @Tags user
// @Description get user by id
// @ID get-user-by-id
// @Accept  json
// @Produce  json
// @Param id   path int64  true  "User ID"
// @Success 200 {object} models.User
// @Failure 400 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /users/:id [get]
func (h *Handler) getUserById(c *gin.Context) {

	userId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid user id param")
		return
	}

	user, err := h.Imp.GetUserById(c, int64(userId))
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
// @Param id   path int64  true  "User ID"
// @Success 200 {string} string "message"
// @Failure 400 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /users/:id [delete]
func (h *Handler) deleteUserById(c *gin.Context) {
	userId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid user id param")
		return
	}

	_, err = h.Imp.DeleteAccount(c, int64(userId))
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, "can't delete account: "+err.Error())
		return
	}

	msg, err := h.Imp.DeleteUserById(c, int64(userId))
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	newResponse(c, http.StatusOK, msg)
}

// @Summary Update User fields
// @Tags user
// @Description Update user fields
// @ID update-user
// @Accept  json
// @Produce  json
// @Param input body models.UpdateUserResponse true "user fields to update"
// @Success 200 {string} string "message"
// @Failure 400 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /users/ [put]
func (h *Handler) updateUser(c *gin.Context) {
	var userChanges models.UpdateUserResponse

	if err := c.BindJSON(&userChanges); err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid input body")
		return
	}
	if userChanges.Username == "" && userChanges.Login == "" && userChanges.Password == "" {
		newErrorResponse(c, http.StatusBadRequest, "at least one parameter must be passed to change")
		return
	}
	msg, err := h.Imp.UpdateUser(c, userChanges)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	newResponse(c, http.StatusOK, msg)
}

// @Summary GetFromCache all users
// @Tags user
// @Description get all users
// @ID get-all-users
// @Accept  json
// @Produce  json
// @Success 200 {string} models.GetAllUsersResponse
// @Failure 400 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /users/ [get]
func (h *Handler) getAllUsers(c *gin.Context) {
	users, err := h.Imp.GetAllUsers(c)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, models.GetAllUsersResponse{
		Data: users,
	})
}
