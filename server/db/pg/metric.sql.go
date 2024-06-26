// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0
// source: metric.sql

package pg

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
)

const addMetric = `-- name: addMetric :exec
INSERT INTO metrics (member_id, date_recorded, unit, metric_value, metric_type)
VALUES ($1, $2, $3, $4, $5)
`

type addMetricParams struct {
	MemberID     int32
	DateRecorded pgtype.Date
	Unit         string
	MetricValue  int32
	MetricType   string
}

func (q *Queries) addMetric(ctx context.Context, arg addMetricParams) error {
	_, err := q.db.Exec(ctx, addMetric,
		arg.MemberID,
		arg.DateRecorded,
		arg.Unit,
		arg.MetricValue,
		arg.MetricType,
	)
	return err
}

const deleteMetric = `-- name: deleteMetric :exec
DELETE FROM metrics
WHERE metric_id=$1
`

func (q *Queries) deleteMetric(ctx context.Context, metricID int32) error {
	_, err := q.db.Exec(ctx, deleteMetric, metricID)
	return err
}

const getMetricsByType = `-- name: getMetricsByType :many
SELECT metric_id, member_id, date_recorded, unit, metric_value, metric_type FROM metrics
WHERE metric_type=$1
`

func (q *Queries) getMetricsByType(ctx context.Context, metricType string) ([]Metric, error) {
	rows, err := q.db.Query(ctx, getMetricsByType, metricType)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Metric
	for rows.Next() {
		var i Metric
		if err := rows.Scan(
			&i.MetricID,
			&i.MemberID,
			&i.DateRecorded,
			&i.Unit,
			&i.MetricValue,
			&i.MetricType,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const updateMetricValue = `-- name: updateMetricValue :exec
UPDATE metrics
SET metric_value=$2
WHERE metric_id=$1
`

type updateMetricValueParams struct {
	MetricID    int32
	MetricValue int32
}

func (q *Queries) updateMetricValue(ctx context.Context, arg updateMetricValueParams) error {
	_, err := q.db.Exec(ctx, updateMetricValue, arg.MetricID, arg.MetricValue)
	return err
}
