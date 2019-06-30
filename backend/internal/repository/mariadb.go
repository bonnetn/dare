package repository

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/bonnetn/dare/backend/internal/entity"
	"database/sql"
	"fmt"
	"context"
	"github.com/bonnetn/dare/backend/internal/controller/uuid"
)

const (
	_get_all_query = "CALL get_all_tasks();"
	_get_query     = "CALL get_task(?);"
	_create_query  = "CALL create_task(?,?,?,?);"
	_update_query  = "CALL update_task(?,?,?,?);"
	_delete_query  = "CALL delete_task(?);"
)

type MariaDBRepository interface {
	Ping() error

	GetTasks(context.Context) ([]entity.Task, error)

	GetTask(context.Context, entity.UUID) (entity.Task, error)
	CreateTask(context.Context, entity.Task) (bool, error)
	UpdateTask(context.Context, entity.Task) (bool, error)
	DeleteTask(context.Context, entity.UUID) error

	Close() error
}

func NewMariadbRepository(config entity.DatabaseConfiguration, uuidController uuid.UUIDController) (MariaDBRepository, error) {
	dataSource := fmt.Sprintf("%s:%s@tcp(%s)/%s", config.Username, config.Password, config.Host, config.Database)
	db, err := sql.Open("mysql", dataSource)
	if err != nil {
		return nil, err
	}
	return &mariaDBRepository{
		db:             db,
		uuidController: uuidController,
	}, nil
}

type mariaDBRepository struct {
	db             *sql.DB
	uuidController uuid.UUIDController
}

func (r mariaDBRepository) Ping() error {
	return r.db.Ping()
}

func (r mariaDBRepository) GetTasks(ctx context.Context) ([]entity.Task, error) {
	rows, err := r.db.QueryContext(ctx, _get_all_query)
	if err != nil {
		return nil, fmt.Errorf("MariaDBRepository GetAll failed: %v", err)
	}
	defer rows.Close()

	var tasksSlice []entity.Task
	for rows.Next() {
		task, err := scanToTask(rows)
		if err != nil {
			return nil, fmt.Errorf("MariaDBRepository GetAll failed: %v", err)
		}
		tasksSlice = append(tasksSlice, task)
	}

	return tasksSlice, nil
}

func (r mariaDBRepository) GetTask(ctx context.Context, uuid entity.UUID) (entity.Task, error) {
	row := r.db.QueryRowContext(ctx, _get_query, uuid)
	task, err := scanToTask(row)

	if err == sql.ErrNoRows {
		return entity.NullTask, nil
	} else if err != nil {
		return entity.NullTask, fmt.Errorf("MariaDBRepository GetTask failed: %v", err)
	}
	return task, nil

}

func (r mariaDBRepository) CreateTask(ctx context.Context, task entity.Task) (bool, error) {
	row := r.db.QueryRowContext(
		ctx,
		_create_query,
		task.UUID(),
		task.Version(),
		task.Name(),
		task.Content(),
	)

	resultTask, err := scanToTask(row)
	if err != nil {
		return false, err
	}

	if resultTask != task {

		return false, nil
	}

	return true, nil
}

// updateTask updates a row to a specific version.
// This function will succeed in two cases:
// 1. Task is already inserted in the database with the specified version and the specified fields.
// 2. Task is at version N-1 in the database and it will be updated.
// Any other case will return an error.
func (r mariaDBRepository) UpdateTask(ctx context.Context, task entity.Task) (bool, error) {
	row := r.db.QueryRowContext(
		ctx,
		_update_query,
		task.UUID(),
		task.Version(),
		task.Name(),
		task.Content(),
	)

	resultTask, err := scanToTask(row)
	if err != nil {
		return false, err
	}

	if resultTask != task {
		return false, nil
	}

	return true, nil
}

func (r mariaDBRepository) DeleteTask(ctx context.Context, uuid entity.UUID) error {
	normalizedUUID, err := r.uuidController.Normalize(uuid)
	if err != nil {
		return err
	}

	_, err = r.db.ExecContext(ctx, _delete_query, normalizedUUID)
	if err != nil {
		return fmt.Errorf("MariaDBRepository Delete failed: %v", err)
	}

	return nil
}

func (r mariaDBRepository) Close() error {
	return r.db.Close()
}

func scanToTask(row scanner) (entity.Task, error) {
	var uuidStr, name, content string
	var version int64
	if err := row.Scan(&uuidStr, &name, &content, &version); err != nil {
		return entity.NullTask, err
	}

	return entity.NewTask(uuidStr, version, name, content), nil
}

type scanner interface {
	Scan(dest ...interface{}) error
}
