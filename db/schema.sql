CREATE TABLE IF NOT EXISTS orders (
  id            VARCHAR(36) PRIMARY KEY,
  customer_name VARCHAR(255) NOT NULL
);

CREATE TABLE IF NOT EXISTS order_items (
  id          VARCHAR(36) PRIMARY KEY,
  order_id    VARCHAR(36) NOT NULL,
  name        VARCHAR(255) NOT NULL,
  quantity    INT NOT NULL,
  unit_price  INT NOT NULL,
  FOREIGN KEY (order_id) REFERENCES orders(id)
);

CREATE TABLE IF NOT EXISTS tokens (
  id      VARCHAR(36) PRIMARY KEY,
  data    BLOB NOT NULL,
  version VARCHAR(255) NOT NULL
);
