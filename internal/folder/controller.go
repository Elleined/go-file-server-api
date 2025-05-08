package folder

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Controller interface {
	create(e *gin.Context)
	remove(e *gin.Context)
	RegisterRoutes(e *gin.Engine)
}

type ControllerImpl struct {
	service Service
}

func NewController(service Service) Controller {
	return &ControllerImpl{
		service: service,
	}
}

func (c ControllerImpl) RegisterRoutes(e *gin.Engine) {
	r := e.Group("/folders")
	{
		r.POST("/", c.create)
		r.DELETE("/:folder", c.remove)
	}
}

func (c ControllerImpl) create(e *gin.Context) {
	name := e.Query("folder")

	sanitizedName, err := c.service.create(name)
	if err != nil {
		e.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	e.JSON(http.StatusCreated, sanitizedName)
}

func (c ControllerImpl) remove(e *gin.Context) {
	name := e.Param("name")

	err := c.service.remove(name)
	if err != nil {
		e.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	e.JSON(http.StatusNoContent, nil)
}
