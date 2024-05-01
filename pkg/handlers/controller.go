package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/maxim-shestakov/final-yandex-project/pkg/service"
)

type Controller struct {
	service service.ServiceInterface
}

func NewController(service service.ServiceInterface) *Controller {
	return &Controller{
		service: service,
	}
}

func (ctrl *Controller) GetTasks(c *gin.Context) {
	//Todo
}

func (ctrl *Controller) CreateTask(c *gin.Context) {
	//Todo
}

func (ctrl *Controller) UpdateTask(c *gin.Context) {
	//Todo
}

func (ctrl *Controller) DeleteTask(c *gin.Context) {
	//Todo
}