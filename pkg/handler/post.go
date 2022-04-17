package handler

import (
	"fmt"
	"github.com/amrchnk/api-gateway/pkg/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

// @Summary Create post
// @Tags posts
// @Description create post with account id that written in context
// @ID create-post
// @Accept  json
// @Produce  json
// @Param input body models.CreatePostRequest true "post info"
// @Success 200 {object} Response
// @Failure 400 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Router /posts/ [post]
func (h *Handler) createPost(c *gin.Context) {
	accountId, exist := c.Get(accountCtx)
	if !exist {
		newResponse(c, http.StatusBadRequest, "account id isn't found in current context!")
		return
	}

	var req models.CreatePostRequest
	if err := c.BindJSON(&req); err != nil {
		newResponse(c, http.StatusBadRequest, "invalid input body")
		return
	}

	images := make([]models.Image, 0, len(req.Images))
	for i := range req.Images {
		image := models.Image{
			Link: req.Images[i],
		}
		images = append(images, image)
	}

	post := models.Post{
		Title:       req.Title,
		Description: req.Description,
		AccountId:   accountId.(int64),
		Images:      images,
	}

	postId, err := h.Imp.CreatePost(c, post)
	if err != nil {
		newResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	newResponse(c,http.StatusOK, fmt.Sprintf("Post with id = %d was created", postId))
}
