CREATE TABLE IF NOT EXISTS product_categories (
  id VARCHAR(25) PRIMARY KEY NOT NULL,
  name VARCHAR(100) NOT NULL,
  description VARCHAR(100),
  parent_id VARCHAR(25),
  asset_id VARCHAR(25),
  vendor_id VARCHAR(25),
  shop_id VARCHAR(25),
  -- FOREIGN KEY(asset_id) REFERENCES assets(id),
  FOREIGN KEY(vendor_id) REFERENCES vendors(id),
  FOREIGN KEY(shop_id) REFERENCES shops(id)
);
