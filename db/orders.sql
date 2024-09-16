-- name: FindOrder :one
SELECT * FROM orders WHERE id = ?
LIMIT 1;

-- name: ListOrderItems :many
SELECT * FROM order_items
WHERE order_id = ?;

-- name: CreateOrder :one
INSERT INTO orders (
  id, customer_name
) VALUES (
  ?, ?
) RETURNING *;

-- name: CreateOrderItem :one
INSERT INTO order_items (
  id, order_id, name, quantity, unit_price
) VALUES (
  ?, ?, ?, ?, ?
) RETURNING *;
