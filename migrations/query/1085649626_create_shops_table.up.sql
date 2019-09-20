CREATE TABLE IF NOT EXISTS shops (
  id VARCHAR(25) PRIMARY KEY NOT NULL,
  name VARCHAR (100),
  vendor_id VARCHAR (100), 
  FOREIGN KEY (vendor_id) REFERENCES vendors(id),
  created_at TIMESTAMP,
  updated_at TIMESTAMP,
);
