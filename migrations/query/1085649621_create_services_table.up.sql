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
  id INT PRIMARY KEY,
  name VARCHAR (100) NOT NULL,
  duration NUMERIC(3) NOT NULL,
  durationType duration_types,
  price NUMERIC(6, 4) DEFAULT 0,
  trending BOOLEAN DEFAULT false,
  categoryId INT,
  vendorId VARCHAR(25),
  FOREIGN KEY(categoryId) REFERENCES categories(id),
  FOREIGN KEY(vendorId) REFERENCES vendors(id)
);

-- duration needs to have a type (days, hour, mins)