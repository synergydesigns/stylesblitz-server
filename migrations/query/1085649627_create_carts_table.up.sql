CREATE TABLE IF NOT EXISTS carts (
  id VARCHAR(25) PRIMARY KEY NOT NULL,
  user_id VARCHAR(100),
  FOREIGN KEY (user_id) REFERENCES users(id),
  created_at TIMESTAMP,
  updated_at TIMESTAMP
);
