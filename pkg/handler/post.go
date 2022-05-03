package handler

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/amrchnk/api-gateway/pkg/models"
	"github.com/gin-gonic/gin"
	"log"
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
	userId, exist := c.Get(userCtx)
	if !exist {
		newResponse(c, http.StatusBadRequest, "user id isn't found in current context!")
		return
	}

	account, err := h.Imp.GetAccountByUserId(c, userId.(int64))
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, "error getting user account")
		return
	}

	form, err := c.MultipartForm()
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	files := form.File["Files"]
	textData := form.Value["PostInfo"]

	filesInput := make([]models.File, 0, len(files))

	for _, file := range files {
		osFile, err := file.Open()
		if err != nil {
			newErrorResponse(c, http.StatusInternalServerError, err.Error())
			return
		}
		filesInput = append(filesInput, models.File{File: osFile, FileName: file.Filename})

		log.Println(file.Filename)
	}

	links, err := h.Imp.FilesUpload(fmt.Sprintf("user%d", userId), filesInput)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	var req models.CreatePostRequest
	req.Images = links

	err = json.Unmarshal([]byte(textData[0]), &req)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
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
		AccountId:   account.Id,
		Images:      images,
		Categories:  req.Categories,
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

	err = DeletePostImages(h, c, int64(postId))
	if err != nil {
		log.Printf("[ERROR]: %v", err)
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	msg, err := h.Imp.DeletePostById(c, int64(postId))
	if err != nil {
		log.Printf("[ERROR]: %v", err)
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	newResponse(c, http.StatusOK, msg)
}
func DeletePostImages(h *Handler, ctx context.Context, postId int64) error {
	images, err := h.Imp.GetImagesFromPost(ctx, postId)
	if err != nil {
		return err
	}
	var links []string
	for _, img := range images {
		links = append(links, img.Link)
	}
	err = h.Imp.DeleteFiles(links)
	if err != nil {
		return err
	}

	return nil
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
		Categories:  post.Categories,
	}

	c.JSON(http.StatusOK, resp)
}

/*func (h *Handler) updatePostById(c *gin.Context) {
	postId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		log.Printf("[ERROR]: %v", err)
		return
	}

	var request models.Post
	request.Id = int64(postId)
	if err := c.BindJSON(&request); err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid input body")
		return
	}

	msg, err := h.Imp.UpdatePost(c, request)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	newResponse(c, http.StatusOK, msg)
}*/

// @Summary Get user posts
// @Tags posts
// @Description Get all user post by user id
// @ID get-user-posts
// @Accept  json
// @Param id   path int64  true  "User ID"
// @Produce  json
// @Success 200 {object} models.GetAllUserPostsResponse
// @Failure 400 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Router /posts/users/:id [get]
func (h *Handler) getAllUserPosts(c *gin.Context) {
	userId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		log.Printf("[ERROR]: %v", err)
		return
	}

	posts, err := h.Imp.GetPostsByUserId(c, int64(userId))
	if err != nil {
		log.Printf("[ERROR]: %v", err)
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
	c.JSON(http.StatusOK, models.GetAllUserPostsResponse{
		Posts: postsResp,
	})
}

// @Summary Get all users posts
// @Tags posts
// @Description Get all users posts
// @ID get-users-posts
// @Param input body models.GetAllUsersPostsRequest true "params for partition"
// @Produce  json
// @Success 200 {array} models.GetAllUsersPosts
// @Failure 400 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Router /posts/users/ [get]
func (h *Handler) getAllUsersPosts(c *gin.Context) {
	var request models.GetAllUsersPostsRequest
	if err := c.BindJSON(&request); err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid input body")
		return
	}

	if request.Limit == 0 {
		request.Limit = 10
	}

	posts, err := h.Imp.GetAllUsersPosts(c, request)
	if err != nil {
		log.Printf("[ERROR]: %v", err)
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	for index := range posts {
		userInfo, err := h.Imp.GetUserById(c, posts[index].UserId)
		if err != nil {
			log.Printf("[ERROR]: %v", err)
			newErrorResponse(c, http.StatusInternalServerError, err.Error())
			return
		}

		posts[index].Username = userInfo.Username
	}

	c.JSON(http.StatusOK, models.GetAllUsersPostsResponse{
		Posts: posts,
	})
}
