CREATE TABLE IF NOT EXISTS states (
  id INT PRIMARY KEY,
  name VARCHAR (100),
  regionCode VARCHAR (5),
  stateCode VARCHAR (5),
  countryId INTEGER,
  FOREIGN KEY (countryId) REFERENCES countries(id)
);