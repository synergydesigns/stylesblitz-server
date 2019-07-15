CREATE TABLE IF NOT EXISTS states (
  id INT PRIMARY KEY,
  name VARCHAR (100),  -- name
  stateCode VARCHAR (5), --
  countryId INTEGER, -- id.fk
  longitude Decimal(9,6), --
  latitude Decimal(9,6), --
  FOREIGN KEY (countryId) REFERENCES countries(id) -- 
);