CREATE TABLE IF NOT EXISTS categories(
  id INT PRIMARY KEY,
  name VARCHAR (100) NOT NULL,
  description VARCHAR(500),
  parentId INT,
  vendorId VARCHAR(25),
  FOREIGN KEY(parentId) REFERENCES categories(id),
  FOREIGN KEY(vendorId) REFERENCES vendors(id),
  UNIQUE(name, vendorId)
);

-- vendors should be able to create custom categoresi
-- this category should be unique for every vendor, that means
-- if vendor a creates category c that vendor should not be allowwd to create another 
-- category c but vendor b should be able to create a c if vendor be dose not have a ctegory
-- c