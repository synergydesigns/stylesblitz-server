CREATE TABLE IF NOT EXISTS vendors (
  id VARCHAR (25),
  userId VARCHAR (25),
  name VARCHAR(100),
  description VARCHAR,
  phone JSONB,
  profileImage VARCHAR (25),
  email VARCHAR(40),
  verified BOOLEAN DEFAULT false,
  created_at TIMESTAMP,
  updated_at TIMESTAMP,
  deleted_at TIMESTAMP,
  PRIMARY KEY (id),
  FOREIGN KEY (userId) REFERENCES users(id),
  FOREIGN KEY (profileImage) REFERENCES assets(id)
);