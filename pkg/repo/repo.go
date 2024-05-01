package repo

import (
	"database/sql"
	"github.com/maxim-shestakov/final-yandex-project/internal/models"
)

type RepoInterface interface {
	GetTasks() ([]models.Task, error)
	CreateTask(task *models.Task) error
	UpdateTask(task *models.Task) error
	DeleteTask(id int64) error
}

type Repo struct {
	RepoInterface
}

func New(db *sql.DB) *Repo {
	return &Repo{
		RepoInterface: NewPostgresRepo(db),
	}
}
