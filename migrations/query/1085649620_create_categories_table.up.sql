CREATE TABLE IF NOT EXISTS categories(
  id INT PRIMARY KEY,
  name VARCHAR (100) NOT NULL,
  description VARCHAR(500),
  parent_id INT,
  vendor_id VARCHAR(25),
  FOREIGN KEY(parent_id) REFERENCES categories(id),
  FOREIGN KEY(vendor_id) REFERENCES vendors(id),
  UNIQUE(name, vendor_id)
);

-- vendors should be able to create custom categoresi
-- this category should be unique for every vendor, that means
-- if vendor a creates category c that vendor should not be allowwd to create another 
-- category c but vendor b should be able to create a c if vendor be dose not have a category c