package handler

import (
	"encoding/json"
	"fmt"
	"github.com/amrchnk/api-gateway/pkg/models"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// @Summary GetFromCache User By Id
// @Tags users
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
// @Security Authorization
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
// @Tags users
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
// @Security Authorization
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
// @Tags users
// @Description Update user fields
// @ID update-user
// @Accept  json
// @Produce  json
// @Param input body models.UpdateUserRequest true "user fields to update"
// @Success 200 {string} string "message"
// @Failure 400 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /users/ [put]
// @Security Authorization
func (h *Handler) updateUser(c *gin.Context) {
	userId, exist := c.Get(userCtx)
	if !exist {
		newErrorResponse(c, http.StatusBadRequest, "user id isn't found in current context!")
		return
	}

	var request models.UpdateUserRequest
	var err error
	if err = c.ShouldBind(&request); err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid input body")
		return
	}

	form, err := c.MultipartForm()
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	var userChanges models.UpdateUserRequestTextData
	textData := form.Value["Json"]
	if textData != nil {
		err = json.Unmarshal([]byte(textData[0]), &userChanges)
		if err != nil {
			newErrorResponse(c, http.StatusBadRequest, err.Error())
			return
		}
	}

	avatar := form.File["Files"]
	var link string
	if avatar != nil {
		user, err := h.Imp.GetUserById(c, userId.(int64))
		if err != nil {
			newErrorResponse(c, http.StatusInternalServerError, err.Error())
			return
		}

		err = h.Imp.DeleteFile(user.ProfileImage)
		if err != nil {
			newErrorResponse(c, http.StatusInternalServerError, err.Error())
			return
		}

		osFile, err := avatar[0].Open()
		if err != nil {
			newErrorResponse(c, http.StatusInternalServerError, err.Error())
			return
		}
		link, err = h.Imp.UploadOneFile(fmt.Sprintf("design_app/avatars/user%d", userId), models.File{File: osFile, FileName: avatar[0].Filename})
		if err != nil {
			newErrorResponse(c, http.StatusInternalServerError, err.Error())
			return
		}
	}

	userChanges.ProfileImage, userChanges.Id = link, userId.(int64)

	if userChanges.Username == "" && userChanges.Login == "" && userChanges.Password == "" && userChanges.ProfileImage == "" {
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
// @Tags users
// @Description get all users
// @ID get-all-users
// @Accept  json
// @Produce  json
// @Success 200 {string} models.GetAllUsersResponse
// @Failure 400 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /users/ [get]
// @Security Authorization
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
