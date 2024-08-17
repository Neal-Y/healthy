package route

import (
	"github.com/gin-gonic/gin"
	"healthy/handler/file"
	"healthy/handler/file/render"
	"healthy/infrastructure"
	"healthy/repository"
	"healthy/service"
)

func InitGinServer() (server *gin.Engine, err error) {
	server = GinRouter()
	err = server.Run("127.0.0.1:8080")
	return
}

func GinRouter() (server *gin.Engine) {
	server = gin.New()
	server.Use(gin.Logger())

	server.LoadHTMLGlob("template/*")
	render.FrontendPage(server)

	server.Static("/uploads", "./uploads")

	fileRepo := repository.NewFileRepository(infrastructure.Db)
	fileService := service.NewFileService(fileRepo)

	api := server.Group("/api")
	file.NewFile(api, fileService)

	return server
}
