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
// @Accept  mpfd
// @Produce  json
// @Param input formData models.CreatePostRequest true "post info"
// @Param input formData file true "post files"
// @Success 200 {object} Response
// @Failure 400 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Router /posts/ [post]
// @Security Authorization
func (h *Handler) createPost(c *gin.Context) {
	userId, exist := c.Get(userCtx)
	if !exist {
		newErrorResponse(c, http.StatusBadRequest, "user id isn't found in current context!")
		return
	}

	account, err := h.Imp.GetAccountByUserId(c, userId.(int64))
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, "error getting user account")
		return
	}
	var request models.CreatePostRequest
	if err = c.ShouldBind(&request); err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid input body")
		return
	}

	form, err := c.MultipartForm()
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	files := form.File["Files"]
	textData := form.Value["PostInfo"]
	var post models.Post

	err = json.Unmarshal([]byte(textData[0]), &post)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

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

	images := make([]models.Image, 0, len(links))
	for i := range links {
		image := models.Image{
			Link: links[i],
		}
		images = append(images, image)
	}

	post.Images, post.AccountId = images, account.Id

	postId, err := h.Imp.CreatePost(c, post)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
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
// @Security Authorization
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
	if len(images) == 0 {
		return nil
	}
	for _, img := range images {
		links = append(links, img.Link)
	}
	err = h.Imp.DeleteFiles(links)
	if err != nil {
		return err
	}

	return nil
}

// @Summary get post by id
// @Tags posts
// @Description get post by post id
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

	resp := models.GetPostByIdResponse{
		Id:          post.Id,
		Title:       post.Title,
		Description: post.Description,
		CreatedAt:   post.CreatedAt,
		UpdatedAt:   post.UpdatedAt,
		Images:      post.Images,
		Categories:  post.Categories,
		UserId:      post.UserId,
	}

	userInfo, err := h.Imp.GetUserById(c, resp.UserId)
	if err != nil {
		log.Printf("[ERROR]: %v", err)
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	resp.Username, resp.ProfileImage = userInfo.Username, userInfo.ProfileImage

	c.JSON(http.StatusOK, resp)
}

// @Summary Update post
// @Tags posts
// @Description Update post by post id
// @ID update-post
// @Accept  json
// @Produce  json
// @Param id   path int64  true  "Post ID"
// @Param input body models.UpdatePostRequest true "post update info"
// @Success 200 {object} Response
// @Failure 400 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Router /posts/:id [put]
// @Security Authorization
func (h *Handler) updatePostById(c *gin.Context) {
	userId, exist := c.Get(userCtx)
	if !exist {
		newErrorResponse(c, http.StatusBadRequest, "user id isn't found in current context!")
		return
	}

	postId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		log.Printf("[ERROR]: %v", err)
		return
	}

	var request models.UpdatePostRequest
	if err = c.ShouldBind(&request); err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid input body:"+err.Error())
		return
	}

	form, err := c.MultipartForm()
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	var postChanges models.UpdatePostRequestTextData
	textData := form.Value["Json"]
	if textData != nil {
		err = json.Unmarshal([]byte(textData[0]), &postChanges)
		if err != nil {
			newErrorResponse(c, http.StatusBadRequest, err.Error())
			return
		}
	}

	files := form.File["Files"]
	if files != nil {
		err = DeletePostImages(h, c, int64(postId))
		if err != nil {
			log.Printf("[ERROR]: %v", err)
			newErrorResponse(c, http.StatusInternalServerError, err.Error())
			return
		}

		filesInput := make([]models.File, 0, len(files))
		for _, file := range files {
			osFile, err := file.Open()
			if err != nil {
				newErrorResponse(c, http.StatusInternalServerError, err.Error())
				return
			}
			filesInput = append(filesInput, models.File{File: osFile, FileName: file.Filename})
		}

		links, err := h.Imp.FilesUpload(fmt.Sprintf("user%d", userId), filesInput)
		if err != nil {
			newErrorResponse(c, http.StatusInternalServerError, err.Error())
			return
		}
		postChanges.Images = links
	}

	postChanges.Id = int64(postId)
	msg, err := h.Imp.UpdatePost(c, postChanges)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	newResponse(c, http.StatusOK, msg)
}

// @Summary GetFromCache user posts
// @Tags posts
// @Description GetFromCache all user post by user id
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
				UpdatedAt:   post.UpdatedAt,
			})
		}
	}
	c.JSON(http.StatusOK, models.GetAllUserPostsResponse{
		Posts: postsResp,
	})
}

// @Summary GetFromCache all users posts
// @Tags posts
// @Description GetFromCache all users posts
// @ID get-users-posts
// @Param input body models.GetAllUsersPostsRequest true "params for partition"
// @Produce  json
// @Success 200 {object} models.GetAllUsersPostsResponse
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

		posts[index].Username, posts[index].ProfileImage = userInfo.Username, userInfo.ProfileImage
	}

	c.JSON(http.StatusOK, models.GetAllUsersPostsResponse{
		Posts: posts,
	})
}
