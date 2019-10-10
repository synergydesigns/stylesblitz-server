CREATE TABLE IF NOT EXISTS service_carts (
  id SERIAL PRIMARY KEY,
  vendor_id VARCHAR (25),
  service_id INT,
  cart_id VARCHAR(25),
  quantity INT DEFAULT 1,
  FOREIGN KEY (vendor_id) REFERENCES vendors(id),
  FOREIGN KEY (service_id) REFERENCES services(id),
  FOREIGN KEY (cart_id) REFERENCES carts(id)
);
