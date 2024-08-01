package service

import (
    "context"
    "github.com/jackc/pgx/v4/pgxpool"
    "errors"
)

var ErrNoRows = errors.New("no rows in result set")

type Task struct {
    ID        int    `json:"id"`
    Title     string `json:"title"`
    Completed bool   `json:"completed"`
}

type TaskService struct {
    db *pgxpool.Pool
}

func NewTaskService(db *pgxpool.Pool) *TaskService {
    return &TaskService{db: db}
}

func (s *TaskService) GetTasks(ctx context.Context) ([]Task, error) {
    rows, err := s.db.Query(ctx, "SELECT id, title, completed FROM tasks")
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    tasks := []Task{}
    for rows.Next() {
        var task Task
        if err := rows.Scan(&task.ID, &task.Title, &task.Completed); err != nil {
            return nil, err
        }
        tasks = append(tasks, task)
    }

    if len(tasks) == 0 {
        return nil, ErrNoRows
    }

    return tasks, nil
}

func (s *TaskService) AddTask(ctx context.Context, task *Task) error {
    return s.db.QueryRow(ctx, "INSERT INTO tasks (title, completed) VALUES ($1, $2) RETURNING id", task.Title, task.Completed).Scan(&task.ID)
}

func (s *TaskService) UpdateTask(ctx context.Context, task *Task) error {
    commandTag, err := s.db.Exec(ctx, "UPDATE tasks SET title=$1, completed=$2 WHERE id=$3", task.Title, task.Completed, task.ID)
    if err != nil {
        return err
    }
    if commandTag.RowsAffected() == 0 {
        return ErrNoRows
    }
    return nil
}

func (s *TaskService) DeleteTask(ctx context.Context, id int) error {
    commandTag, err := s.db.Exec(ctx, "DELETE FROM tasks WHERE id=$1", id)
    if err != nil {
        return err
    }
    if commandTag.RowsAffected() == 0 {
        return ErrNoRows
    }
    return nil
}
