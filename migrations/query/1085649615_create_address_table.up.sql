CREATE TABLE IF NOT EXISTS address (
  id SERIAL PRIMARY KEY,
  country_id INT,
  state_id INT,
  zipcode NUMERIC(8),
  city VARCHAR (100),
  street VARCHAR (250),
  longitude FLOAT(5),
  latitude FLOAT(5),
  geog geography(POINT, 4326)
);
