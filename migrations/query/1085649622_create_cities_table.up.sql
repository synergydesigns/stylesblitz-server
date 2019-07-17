CREATE TABLE IF NOT EXISTS cities (
  id SERIAL,
  name VARCHAR (100),
  state_id INT,
  country_id INT,
  longitude FLOAT(5),
  latitude FLOAT(5),
  PRIMARY KEY(id),
  FOREIGN KEY(state_id) REFERENCES states(id),
  FOREIGN KEY(country_id) REFERENCES countries(id)
);