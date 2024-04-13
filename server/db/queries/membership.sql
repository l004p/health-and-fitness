-- name: createMembership :exec
INSERT INTO memberships (member_id, title, membership_status)
VALUES ($1, $2, 'ACTIVE');

-- name: cancelMembership :exec
UPDATE memberships
SET membership_status='INACTIVE'
WHERE member_id=$1 AND title=$2;
