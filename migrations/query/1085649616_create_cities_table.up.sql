CREATE TABLE IF EXISTS cities (
  id INT,
  name VARCHAR (100),
  stateId INT,
  countryId INT,
  longitude Decimal(9,6), --
  latitude Decimal(9,6), --
  PRIMARY KEY(id),
  FOREIGN KEY(stateId) REFERENCES states(id),
  FOREIGN KEY(stateId) REFERENCES countries(id)
)