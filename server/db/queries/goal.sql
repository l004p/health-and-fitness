-- name: addGoal :exec
INSERT INTO goals (member_id, goal_date, title, unit, goal_value)
VALUES ($1, $2, $3, $4, $5);

-- name: deleteGoal :exec
DELETE FROM goals WHERE goal_id=$1;

-- name: updateGoalValue :exec
UPDATE goals
SET goal_value=$2
WHERE goal_id=$1;

-- name: updateGoalType :exec
UPDATE goals
SET title=$2, unit=$3 AND goal_value=$4 AND goal_type=$5
WHERE goal_id=$1;

-- name: getGoals :many
SELECT * FROM goals
WHERE member_id=$1;

-- name: getGoal :one
SELECT * FROM goals
WHERE goal_id=$1;