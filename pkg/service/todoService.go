package service

import (
	"github.com/maxim-shestakov/final-yandex-project/pkg/repo"
	"github.com/maxim-shestakov/final-yandex-project/internal/models"
)

type todoService struct {
	repo repo.RepoInterface
}

func NewService(r repo.RepoInterface) *todoService {
	return &todoService{
		repo: r,
	}
}

func (s *todoService) GetTasks() ([]models.Task, error) {
	return s.repo.GetTasks()
}

func (s *todoService) CreateTask(task *models.Task) error {
	return s.repo.CreateTask(task)
}

func (s *todoService) UpdateTask(task *models.Task) error {
	return s.repo.UpdateTask(task)
}

func (s *todoService) DeleteTask(id int64) error {
	return s.repo.DeleteTask(id)
}
