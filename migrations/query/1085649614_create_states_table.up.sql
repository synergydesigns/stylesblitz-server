CREATE TABLE IF NOT EXISTS states (
  id SERIAL PRIMARY KEY,
  name VARCHAR (100),  -- name
  state_code VARCHAR (5), --
  country_id INTEGER, -- id.fk
  longitude FLOAT(5),
  latitude FLOAT(5),
  FOREIGN KEY (country_id) REFERENCES countries(id) -- 
);