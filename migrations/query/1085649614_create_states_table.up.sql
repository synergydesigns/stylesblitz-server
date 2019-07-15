CREATE TABLE IF NOT EXISTS states (
  id INT PRIMARY KEY,
  name VARCHAR (100),  -- name
  state_code VARCHAR (5), --
  country_id INTEGER, -- id.fk
  longitude Decimal(9,6), --
  latitude Decimal(9,6), --
  FOREIGN KEY (country_id) REFERENCES countries(id) -- 
);