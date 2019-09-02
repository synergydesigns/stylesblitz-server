CREATE TABLE IF NOT EXISTS users (
  id VARCHAR(25) NOT NULL,
  firstname VARCHAR (255) NOT NULL,
  lastname VARCHAR (255) NOT NULL,
  username VARCHAR (100) NOT NULL,
  email VARCHAR (100) NOT NULL,
  password VARCHAR (100) NOT NULL,
  bio VARCHAR,
  phone VARCHAR (20),
  profile_image VARCHAR (25) NULL,
  wall_image VARCHAR (25) NULL,
  address_id INT,
  created_at TIMESTAMP,
  updated_at TIMESTAMP,
  deleted_at TIMESTAMP,
  PRIMARY KEY (id),
  UNIQUE(phone)
);
