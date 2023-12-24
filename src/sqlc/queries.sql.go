// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.24.0
// source: queries.sql

package sqlc

import (
	"context"
	"database/sql"
)

const getTodosWithLimit = `-- name: GetTodosWithLimit :many
SELECT body
FROM todo
LIMIT $1
`

func (q *Queries) GetTodosWithLimit(ctx context.Context, dollar_1 sql.NullInt64) ([]string, error) {
	rows, err := q.db.QueryContext(ctx, getTodosWithLimit, dollar_1)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []string
	for rows.Next() {
		var body string
		if err := rows.Scan(&body); err != nil {
			return nil, err
		}
		items = append(items, body)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}