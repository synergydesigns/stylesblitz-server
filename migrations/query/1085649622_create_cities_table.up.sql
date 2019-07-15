CREATE TABLE IF NOT EXISTS cities (
  id INT,
  name VARCHAR (100),
  state_id INT,
  country_id INT,
  longitude Decimal(9,6), --
  latitude Decimal(9,6), --
  PRIMARY KEY(id),
  FOREIGN KEY(state_id) REFERENCES states(id),
  FOREIGN KEY(country_id) REFERENCES countries(id)
);