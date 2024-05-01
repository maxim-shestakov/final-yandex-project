package repo

import (
	"database/sql"

	_ "github.com/lib/pq"

	"github.com/maxim-shestakov/final-yandex-project/internal/models"
)

type PostgresRepo struct {
	db *sql.DB
}

func NewPostgresRepo(db *sql.DB) *PostgresRepo {
	return &PostgresRepo{
		db: db,
	}
}

func (r *PostgresRepo) GetTasks() ([]models.Task, error) {
	// TODO
	return nil, nil
}

func (r *PostgresRepo) CreateTask(task *models.Task) error {
	// TODO
	return nil
}

func (r *PostgresRepo) UpdateTask(task *models.Task) error {
	// TODO
	return nil
}

func (r *PostgresRepo) DeleteTask(id int64) error {
	// TODO
	return nil
}
