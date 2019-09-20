CREATE TABLE IF NOT EXISTS service_cart (
  id INT PRIMARY KEY NOT NULL,
  vendor_id VARCHAR (25),
  service_id VARCHAR(25),
  cart_id VARCHAR(25),
  quantity INT DEFAULT 1,
  FOREIGN KEY (vendor_id) REFERENCES vendors(id),
  FOREIGN KEY (service_id) REFERENCES services(id),
  FOREIGN KEY (cart_id) REFERENCES carts(id),
);
