// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.16.0
// source: routes.sql

package hikedb

import (
	"context"
)

const createRoute = `-- name: CreateRoute :one
INSERT INTO routes (
  user_id,
  title,
  description,
  location,
  destination,
  height,
  level
) VALUES (
  $1, $2, $3, $4, $5, $6, $7
)
RETURNING id, user_id, title, description, location, destination, height, level, "createdAt", "updatedAt"
`

type CreateRouteParams struct {
	UserID      int32   `json:"user_id"`
	Title       string  `json:"title"`
	Description string  `json:"description"`
	Location    string  `json:"location"`
	Destination float64 `json:"destination"`
	Height      float64 `json:"height"`
	Level       string  `json:"level"`
}

func (q *Queries) CreateRoute(ctx context.Context, arg CreateRouteParams) (Route, error) {
	row := q.db.QueryRowContext(ctx, createRoute,
		arg.UserID,
		arg.Title,
		arg.Description,
		arg.Location,
		arg.Destination,
		arg.Height,
		arg.Level,
	)
	var i Route
	err := row.Scan(
		&i.ID,
		&i.UserID,
		&i.Title,
		&i.Description,
		&i.Location,
		&i.Destination,
		&i.Height,
		&i.Level,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const deleteRoute = `-- name: DeleteRoute :exec
DELETE FROM routes
WHERE id = $1
`

func (q *Queries) DeleteRoute(ctx context.Context, id int32) error {
	_, err := q.db.ExecContext(ctx, deleteRoute, id)
	return err
}

const getRoute = `-- name: GetRoute :one
SELECT id, user_id, title, description, location, destination, height, level, "createdAt", "updatedAt" FROM routes
WHERE id = $1 LIMIT 1
`

func (q *Queries) GetRoute(ctx context.Context, id int32) (Route, error) {
	row := q.db.QueryRowContext(ctx, getRoute, id)
	var i Route
	err := row.Scan(
		&i.ID,
		&i.UserID,
		&i.Title,
		&i.Description,
		&i.Location,
		&i.Destination,
		&i.Height,
		&i.Level,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const listRoutes = `-- name: ListRoutes :many
SELECT id, user_id, title, description, location, destination, height, level, "createdAt", "updatedAt" FROM routes
ORDER BY id
LIMIT $1
OFFSET $2
`

type ListRoutesParams struct {
	Limit  int32 `json:"limit"`
	Offset int32 `json:"offset"`
}

func (q *Queries) ListRoutes(ctx context.Context, arg ListRoutesParams) ([]Route, error) {
	rows, err := q.db.QueryContext(ctx, listRoutes, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Route{}
	for rows.Next() {
		var i Route
		if err := rows.Scan(
			&i.ID,
			&i.UserID,
			&i.Title,
			&i.Description,
			&i.Location,
			&i.Destination,
			&i.Height,
			&i.Level,
			&i.CreatedAt,
			&i.UpdatedAt,
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

const updateRoute = `-- name: UpdateRoute :one
UPDATE routes
SET user_id = $2,
    title = $3,
    description = $4,
    location = $5,
    destination = $6,
    height = $7,
    level = $8
WHERE id = $1
RETURNING id, user_id, title, description, location, destination, height, level, "createdAt", "updatedAt"
`

type UpdateRouteParams struct {
	ID          int32   `json:"id"`
	UserID      int32   `json:"user_id"`
	Title       string  `json:"title"`
	Description string  `json:"description"`
	Location    string  `json:"location"`
	Destination float64 `json:"destination"`
	Height      float64 `json:"height"`
	Level       string  `json:"level"`
}

func (q *Queries) UpdateRoute(ctx context.Context, arg UpdateRouteParams) (Route, error) {
	row := q.db.QueryRowContext(ctx, updateRoute,
		arg.ID,
		arg.UserID,
		arg.Title,
		arg.Description,
		arg.Location,
		arg.Destination,
		arg.Height,
		arg.Level,
	)
	var i Route
	err := row.Scan(
		&i.ID,
		&i.UserID,
		&i.Title,
		&i.Description,
		&i.Location,
		&i.Destination,
		&i.Height,
		&i.Level,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}
