CREATE TABLE IF NOT EXISTS countries (
  id SERIAL,
  name VARCHAR (100),
  country_code NUMERIC (6), 
  iso_code VARCHAR (5),
  currency VARCHAR(10),
  currency_symbol CHAR(1),
  currency_code CHAR(5),
  longitude FLOAT(5),
  latitude FLOAT(5),
  PRIMARY KEY(id)
);
