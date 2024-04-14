-- name: addMemberToRoster :exec
INSERT INTO member_trained (trainer_id, member_id)
VALUES ($1, $2);

-- name: removeMemberFromRoster :exec
DELETE FROM member_trained
WHERE trainer_id=$1 AND member_id=$2;

-- name: getAllMembersTraining :many
SELECT u.first_name, u.last_name, u.username
FROM users u 
INNER JOIN member_trained mt
ON u.user_id=mt.member_id
WHERE mt.trainer_id=$1;

-- name: addInterest :exec
INSERT INTO trainer_interests(trainer_id, interest)
VALUES($1, $2);

-- name: getTrainerInterests :many
SELECT interest
FROM trainer_interests
WHERE trainer_id=$1;

-- name: getTrainersByInterest :many
SELECT u.first_name, u.last_name, u.username
FROM trainer_interests ti
INNER JOIN users u
ON u.user_id=ti.trainer_id
WHERE ti.interest=$1;

-- name: updateInterest :exec
UPDATE trainer_interests
SET interest=$2
WHERE trainer_id=$1;

-- name: deleteInterest :exec
DELETE FROM trainer_interests
WHERE trainer_id=$1 AND interest=$2;

-- name: getUserByName :many
SELECT username, first_name, last_name
FROM users
WHERE first_name=$1 OR last_name=$2;

