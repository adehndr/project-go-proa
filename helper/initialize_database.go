package helper

import (
	"context"
	"database/sql"
)

func InitializeDatabase(db *sql.DB,ctx context.Context) {
	query := "CREATE TABLE IF NOT EXISTS task_table (id serial PRIMARY KEY, task_detail VARCHAR(4096), assignee VARCHAR(512), deadline DATE, is_finished BOOLEAN, created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP );"
	_, err := db.ExecContext(ctx,query)
	if err != nil {
		panic(err)
	}
}