package handler

import (
	"fmt"
	"github.com/amrchnk/api-gateway/pkg/models"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// @Summary Create post
// @Tags posts
// @Description create post with account id that written in context
// @ID delete-post
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

	newResponse(c, http.StatusOK, fmt.Sprintf("Post with id = %d was created", postId))
}

// @Summary Delete post
// @Tags posts
// @Description delete post by post id
// @ID create-post
// @Accept  json
// @Produce  json
// @Param id   path int64  true  "Post ID"
// @Success 200 {object} Response
// @Failure 400 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Router /posts/:id [delete]
func (h *Handler) deletePostById(c *gin.Context) {
	postId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	msg, err := h.Imp.DeletePostById(c, int64(postId))
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	newResponse(c, http.StatusOK, msg)
}

// @Summary Get post
// @Tags posts
// @Description Get post by post id
// @ID get-post
// @Accept  json
// @Produce  json
// @Param id   path int64  true  "Post ID"
// @Success 200 {object} models.GetPostByIdResponse
// @Failure 400 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Router /posts/:id [get]
func (h *Handler) getPostById(c *gin.Context) {
	postId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	post, err := h.Imp.GetPostById(c, int64(postId))
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	imageLinks := make([]string, 0, len(post.Images))
	for _, image := range post.Images {
		imageLinks = append(imageLinks, image.Link)
	}

	resp := models.GetPostByIdResponse{
		Id:          post.Id,
		Title:       post.Title,
		Description: post.Description,
		CreatedAt:   post.CreatedAt,
		Images:      imageLinks,
	}

	c.JSON(http.StatusOK, resp)
}

// @Summary Get user posts
// @Tags posts
// @Description Get all user post by account id that place in context
// @ID get-user-posts
// @Accept  json
// @Produce  json
// @Success 200 {object} models.GetAllUserPostsResponse
// @Failure 400 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Router /posts/user [get]
func (h *Handler) getAllUserPosts(c *gin.Context) {
	accountId, _ := c.Get(accountCtx)

	posts, err := h.Imp.GetPostsByAccountId(c, accountId.(int64))
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	postsResp := make([]models.GetPostByIdResponse, 0, len(posts))
	if len(posts) != 0 {
		for _, post := range posts {
			imageLinks := make([]string, 0, len(post.Images))
			for _, image := range post.Images {
				imageLinks = append(imageLinks, image.Link)
			}
			postsResp = append(postsResp, models.GetPostByIdResponse{
				Id:          post.Id,
				Title:       post.Title,
				Description: post.Description,
				Images:      imageLinks,
				CreatedAt:   post.CreatedAt,
			})
		}
	}
	c.JSON(http.StatusOK,models.GetAllUserPostsResponse{
		Posts: postsResp,
	})
}
