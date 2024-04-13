-- name: addMetric :exec
INSERT INTO metrics (member_id, date_recorded, unit, metric_value, metric_type)
VALUES ($1, $2, $3, $4, $5);

-- name: deleteMetric :exec
DELETE FROM metrics
WHERE metric_id=$1;


-- name: updateMetricValue :exec
UPDATE metrics
SET metric_value=$2
WHERE metric_id=$1;


-- name: getMetricsByType :many
SELECT * FROM metrics
WHERE metric_type=$1;