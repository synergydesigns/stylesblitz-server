CREATE TABLE IF NOT EXISTS vendors (
  id VARCHAR (25),
  user_id VARCHAR (25),
  name VARCHAR(100),
  description VARCHAR,
  phone JSONB,
  profile_image VARCHAR (25),
  email VARCHAR(40),
  verified BOOLEAN DEFAULT false,
  created_at TIMESTAMP,
  updated_at TIMESTAMP,
  deleted_at TIMESTAMP,
  PRIMARY KEY (id),
  FOREIGN KEY (user_id) REFERENCES users(id),
  FOREIGN KEY (profile_image) REFERENCES assets(id)
);