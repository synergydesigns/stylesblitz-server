CREATE TABLE IF NOT EXISTS countries (
  id INT,
  name VARCHAR (100), -- name 
  countryCode VARCHAR (5), -- countryCode,
  isoCode VARCHAR (5), -- adminCodes1.ISO3166_2
  longitude Decimal(9,6), --
  latitude Decimal(9,6), --
  PRIMARY KEY(id)
);
