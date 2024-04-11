-- name: getAllTrainers :many
SELECT u.first_name, u.last_name, u.username
FROM users u 
INNER JOIN member_trained mt
ON u.user_id=mt.trainer_id
WHERE mt.member_id=$1;