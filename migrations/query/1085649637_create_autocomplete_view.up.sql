CREATE
OR REPLACE VIEW autocomplete AS
SELECT
  id,
  vendors.name as query,
  vendors.tsv,
  'vendors' as type
FROM vendors
UNION
SELECT
  services.id::VARCHAR,
  services.name as query,
  services.tsv,
  'services' as type
FROM services
UNION
SELECT
  categories.id::VARCHAR,
  categories.name as query,
  categories.tsv,
  'categories' as type
FROM categories;