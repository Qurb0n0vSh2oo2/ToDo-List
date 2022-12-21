package rcreatetasks

import (
	"context"
	"errors"
	"fmt"
	"topro/pkg/utils"

	"github.com/jackc/pgx/v4"
)

func (t *RTaskCreateProvider) RCreateTask(task utils.Task) (err error) {
	query := fmt.Sprintf(
		`INSERT INTO tasks(name, description, status, user_id)
		VALUES ('%v', '%v', '%v', %v)`,
		task.Name,
		task.Description,
		task.Status,
		task.UserID,
	)

	_, err = t.Postgres.Exec(context.Background(), query)
	if err == pgx.ErrNoRows {
		return errors.New("not found")
	}

	return
}
