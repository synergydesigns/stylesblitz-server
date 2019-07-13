CREATE TABLE IF NOT EXISTS assets (
  id VARCHAR(25) PRIMARY KEY,
  title VARCHAR(100),
  description VARCHAR(100),
  caption VARCHAR(100),
  alt VARCHAR(100),
  mediaType VARCHAR(100),
  mimeType VARCHAR(100),
  width VARCHAR(100),
  height VARCHAR(100),
  filename VARCHAR(100),
  size NUMERIC,
  created_at TIMESTAMP,
  updated_at TIMESTAMP,
  deleted_at TIMESTAMP
);