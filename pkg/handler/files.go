package handler

import (
	"github.com/amrchnk/api-gateway/pkg/models"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func (h *Handler) uploadFileInCloudinary(c *gin.Context) {
	form, err := c.MultipartForm()
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	files := form.File["Files"]

	filesInput := make([]models.File, 0, len(files))

	for _, file := range files {
		osFile, err := file.Open()
		if err != nil {
			newErrorResponse(c, http.StatusInternalServerError, err.Error())
			return
		}
		filesInput = append(filesInput, models.File{File: osFile})

		log.Println(file.Filename)
	}

	links, err := h.Imp.FilesUpload(filesInput)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"images": links,
	})
}
