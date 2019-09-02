
CREATE TABLE IF NOT EXISTS user_assets (
  user_id VARCHAR(25) NOT NULL,
  asset_id VARCHAR(25) NOT NULL,
  FOREIGN KEY (asset_id) REFERENCES assets(id),
  FOREIGN KEY (user_id) REFERENCES users(id)
);