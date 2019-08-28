CREATE TABLE IF NOT EXISTS vendor_assets (
  vendor_id VARCHAR(25) NOT NULL,
  asset_id VARCHAR(25) NOT NULL,
  FOREIGN KEY (asset_id) REFERENCES assets(id),
  FOREIGN KEY (vendor_id) REFERENCES vendors(id)
);