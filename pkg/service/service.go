package service

import (
	"github.com/maxim-shestakov/final-yandex-project/internal/models"
	"github.com/maxim-shestakov/final-yandex-project/pkg/repo"
)

type Service struct {
	ServiceInterface
}

type ServiceInterface interface {
	GetTasks() ([]models.Task, error)
	CreateTask(task *models.Task) error
	UpdateTask(task *models.Task) error
	DeleteTask(id int64) error
}

func New(repo *repo.Repo) *Service {
	return &Service{
		ServiceInterface: NewService(repo),
	}
}
