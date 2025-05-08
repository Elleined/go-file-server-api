package file

import (
	"github.com/gin-gonic/gin"
	"mime/multipart"
	"net/http"
)

type Controller interface {
	upload(ctx *gin.Context)
	read(ctx *gin.Context)
	delete(ctx *gin.Context)
	RegisterRoutes(ctx *gin.Engine)
}

type ControllerImpl struct {
	service Service
}

func NewController(service Service) Controller {
	return &ControllerImpl{
		service: service,
	}
}

func (c ControllerImpl) RegisterRoutes(ctx *gin.Engine) {
	r := ctx.Group("/folders/:folder/files")
	{
		r.POST("", c.upload)
		r.GET("/:file", c.read)
		r.DELETE("/:file", c.delete)
	}
}

func (c ControllerImpl) upload(ctx *gin.Context) {
	// Getting file from request
	file, header, err := ctx.Request.FormFile("file")
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "can't get the file from request" + err.Error(),
		})
		return
	}
	defer func(file multipart.File) {
		err := file.Close()
		if err != nil {
			return
		}
	}(file)

	// Uploading file to local machine
	folder := ctx.Param("folder")
	fileName, err := c.service.upload(folder, file, *header)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "can't upload file " + err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, fileName)
}

func (c ControllerImpl) read(ctx *gin.Context) {
	folder := ctx.Param("folder")
	file := ctx.Param("file")

	if err := c.service.read(folder, file); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "can't read file " + err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "successfully read " + file,
	})
}

func (c ControllerImpl) delete(ctx *gin.Context) {
	folder := ctx.Param("folder")
	file := ctx.Param("file")

	if err := c.service.delete(folder, file); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "can't delete file " + err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "successfully deleted file " + file,
	})
}
