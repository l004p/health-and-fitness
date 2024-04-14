-- name: createBill :exec
INSERT INTO bills (user_id, bill_description, bill_status)
VALUES ($1, $2, 'PENDING');

-- name: payBill :exec
UPDATE bills
SET bill_status='COMPLETE'
WHERE bill_id=$1;

-- name: getAllUserBills :many
SELECT *
FROM bills
WHERE user_id=$1;

-- name: cancelBill :exec
UPDATE bills
SET bill_status='CANCELLED'
WHERE bill_id=$1;

-- name: getBillByID :one
SELECT * FROM bills WHERE bill_id=$1;