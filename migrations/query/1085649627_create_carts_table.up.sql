CREATE TYPE cart_types AS ENUM('product', 'service');

CREATE TABLE IF NOT EXISTS carts (
  id VARCHAR(25) PRIMARY KEY NOT NULL,
  user_id VARCHAR(25),
  vendor_id VARCHAR(25),
  type cart_types,
  type_id VARCHAR(25),
  FOREIGN KEY (user_id) REFERENCES users(id),
  quantity INT DEFAULT 1,
  created_at TIMESTAMP,
  updated_at TIMESTAMP
);
