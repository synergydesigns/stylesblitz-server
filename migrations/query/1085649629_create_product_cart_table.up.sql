CREATE TABLE IF NOT EXISTS product_carts (
  id SERIAL PRIMARY KEY,
  vendor_id VARCHAR (25),
  product_id VARCHAR(25),
  shop_id VARCHAR(25),
  cart_id VARCHAR(25),
  quantity INT DEFAULT 1,
  FOREIGN KEY (vendor_id) REFERENCES vendors(id),
  FOREIGN KEY (product_id) REFERENCES products(id),
  FOREIGN KEY (cart_id) REFERENCES carts(id)
);
