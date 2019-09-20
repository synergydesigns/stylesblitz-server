CREATE TABLE IF NOT EXISTS products (
  id INT PRIMARY KEY NOT NULL,
  category_id VARCHAR (100),
  brand_id VARCHAR (100),
  hot BOOLEAN,
  available INT,
  created_at TIMESTAMP,
  updated_at TIMESTAMP,
);
