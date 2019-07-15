CREATE TABLE IF NOT EXISTS address (
  id INT PRIMARY KEY,
  country_id INT,
  state_id INT,
  zip_code NUMERIC(8),
  city VARCHAR (100),
  street VARCHAR (250),
  FOREIGN KEY (country_id) REFERENCES countries(id),
  FOREIGN KEY (state_id) REFERENCES states(id)
);
