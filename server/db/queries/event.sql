-- name: addEvent :exec
INSERT INTO calender_events (room_id, event_type, max_attendees, event_status, event_title, start_time, end_time) 
VALUES ($1, $2, $3, 'UPCOMING', $4, $5, $6);

-- name: deleteEvent :exec
DELETE FROM calender_events
WHERE event_id=$1;

-- name: changeEventStatus :exec
UPDATE calender_events
SET event_status=$2
WHERE event_id=$1;

-- name: registerForEvent :exec
INSERT INTO attending_event (member_id, event_id)
VALUES ($1, $2);

-- name: cancelEventRegistration :exec
DELETE FROM attending_event
WHERE member_id=$1 AND event_id=$2;

-- name: removeEventLeader :exec
DELETE FROM leading_event
WHERE trainer_id=$1 AND event_id=$2;

-- name: addEventLeader :exec
INSERT INTO leading_event (trainer_id, event_id)
VALUES ($1, $2);

-- name: changeEventRoom :exec
UPDATE calender_events
SET room_id=$2
WHERE event_id=$1;

-- name: getNumRegistrees :one
SELECT COUNT(*)
FROM attending_event
WHERE event_id=$1;