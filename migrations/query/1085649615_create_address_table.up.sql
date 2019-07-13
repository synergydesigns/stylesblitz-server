CREATE TABLE IF NOT EXISTS address (
  id INT PRIMARY KEY,
  countryId INT,
  stateId INT,
  zipCode NUMERIC(8),
  city VARCHAR (100),
  street VARCHAR (250),
  FOREIGN KEY (countryId) REFERENCES countries(id),
  FOREIGN KEY (stateId) REFERENCES states(id)
);
