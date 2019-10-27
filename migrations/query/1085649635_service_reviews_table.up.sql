CREATE TABLE IF NOT EXISTS service_reviews (
  id SERIAL PRIMARY KEY,
  user_id VARCHAR(25) NOT NULL,
  vendor_id VARCHAR(25) NOT NULL,
  service_id INT NOT NULL,
  text TEXT,
  rating VARCHAR,
  parent_id INT,
  created_at TIMESTAMP,
  updated_at TIMESTAMP,
  FOREIGN KEY(user_id) REFERENCES users(id),
  FOREIGN KEY(vendor_id) REFERENCES vendors(id),
  FOREIGN KEY(service_id) REFERENCES services(id)
);
