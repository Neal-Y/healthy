package file

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *File) UploadHandler(c *gin.Context) {
	file, err := c.FormFile("image")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Error uploading file"})
		return
	}

	response, err := h.service.ProcessAndAnalyzeImage(file)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"response": response})
}
