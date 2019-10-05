CREATE TABLE IF NOT EXISTS products (
  id VARCHAR(25) PRIMARY KEY NOT NULL,
  name VARCHAR(100) NOT NULL,
  category_id VARCHAR(25),
  vendor_id VARCHAR(25),
  brand_id VARCHAR (25),
  hot BOOLEAN,
  available INT,
  created_at TIMESTAMP,
  updated_at TIMESTAMP,
  FOREIGN KEY (vendor_id) REFERENCES vendors(id)
);
