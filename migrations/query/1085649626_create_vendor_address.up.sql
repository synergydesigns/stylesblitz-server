CREATE TABLE IF NOT EXISTS vendor_address (
  vendor_id VARCHAR(25) NOT NULL,
  address_id SERIAL NOT NULL,
  FOREIGN KEY (vendor_id) REFERENCES vendors(id),
  FOREIGN KEY (address_id) REFERENCES address(id)
);