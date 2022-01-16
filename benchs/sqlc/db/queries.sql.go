// Code generated by sqlc. DO NOT EDIT.
// source: queries.sql

package db

import (
	"context"
)

const createModel = `-- name: CreateModel :one
INSERT INTO models (NAME, title, fax, web, age, "right", counter)
VALUES ($1, $2, $3, $4, $5, $6, $7)
RETURNING id, name, title, fax, web, age, "right", counter
`

type CreateModelParams struct {
	Name    string
	Title   string
	Fax     string
	Web     string
	Age     int32
	Right   bool
	Counter int64
}

func (q *Queries) CreateModel(ctx context.Context, arg CreateModelParams) (Model, error) {
	row := q.db.QueryRowContext(ctx, createModel,
		arg.Name,
		arg.Title,
		arg.Fax,
		arg.Web,
		arg.Age,
		arg.Right,
		arg.Counter,
	)
	var i Model
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Title,
		&i.Fax,
		&i.Web,
		&i.Age,
		&i.Right,
		&i.Counter,
	)
	return i, err
}

const getModel = `-- name: GetModel :one
SELECT id, name, title, fax, web, age, "right", counter
FROM models
WHERE id = $1
`

func (q *Queries) GetModel(ctx context.Context, id int32) (Model, error) {
	row := q.db.QueryRowContext(ctx, getModel, id)
	var i Model
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Title,
		&i.Fax,
		&i.Web,
		&i.Age,
		&i.Right,
		&i.Counter,
	)
	return i, err
}

const listModels = `-- name: ListModels :many
SELECT id, name, title, fax, web, age, "right", counter
FROM models
WHERE ID > $1
ORDER BY ID
LIMIT $2
`

type ListModelsParams struct {
	ID    int32
	Limit int32
}

func (q *Queries) ListModels(ctx context.Context, arg ListModelsParams) ([]Model, error) {
	rows, err := q.db.QueryContext(ctx, listModels, arg.ID, arg.Limit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Model
	for rows.Next() {
		var i Model
		if err := rows.Scan(
			&i.ID,
			&i.Name,
			&i.Title,
			&i.Fax,
			&i.Web,
			&i.Age,
			&i.Right,
			&i.Counter,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const updateModel = `-- name: UpdateModel :exec
UPDATE models
SET name = $1,
    title = $2,
    fax = $3,
    web = $4,
    age = $5,
    "right" = $6,
    counter = $7
WHERE id = $8
`

type UpdateModelParams struct {
	Name    string
	Title   string
	Fax     string
	Web     string
	Age     int32
	Right   bool
	Counter int64
	ID      int32
}

func (q *Queries) UpdateModel(ctx context.Context, arg UpdateModelParams) error {
	_, err := q.db.ExecContext(ctx, updateModel,
		arg.Name,
		arg.Title,
		arg.Fax,
		arg.Web,
		arg.Age,
		arg.Right,
		arg.Counter,
		arg.ID,
	)
	return err
}