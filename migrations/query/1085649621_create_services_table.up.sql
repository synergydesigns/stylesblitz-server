DO $$
  BEGIN
    IF NOT EXISTS (SELECT 1 FROM pg_type WHERE typname = 'duration_types') THEN
        CREATE TYPE duration_types AS enum(
          'days',
          'hours',
          'mins'
        );
  END IF;
END $$;

CREATE TABLE IF NOT EXISTS services (
  id SERIAL PRIMARY KEY,
  name VARCHAR (100) NOT NULL,
  duration NUMERIC(3) NOT NULL,
  duration_type duration_types,
  price NUMERIC(6, 4) DEFAULT 0,
  trending BOOLEAN DEFAULT false,
  category_id INT,
  vendor_id VARCHAR(25),
  created_at TIMESTAMP,
  updated_at TIMESTAMP,
  deleted_at TIMESTAMP,
  FOREIGN KEY(category_id) REFERENCES categories(id),
  FOREIGN KEY(vendor_id) REFERENCES vendors(id),
  UNIQUE(name, vendor_id)
);