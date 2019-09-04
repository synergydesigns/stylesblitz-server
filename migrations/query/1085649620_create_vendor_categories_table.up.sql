CREATE TABLE IF NOT EXISTS categories(
  id SERIAL PRIMARY KEY,
  name VARCHAR (100) NOT NULL,
  description VARCHAR(500),
  vendor_id VARCHAR(25),
  FOREIGN KEY(vendor_id) REFERENCES vendors(id),
  UNIQUE(name, vendor_id)
);