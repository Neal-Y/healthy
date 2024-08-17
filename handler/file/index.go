package file

import (
	"github.com/gin-gonic/gin"
	"healthy/service"
)

type File struct {
	service service.FileService
}

func NewFile(r *gin.RouterGroup, service service.FileService) *File {
	h := &File{
		service: service,
	}

	newRoute(h, r)

	return h
}

func newRoute(h *File, r *gin.RouterGroup) {
	r.POST("/upload_image", h.UploadHandler)
}
