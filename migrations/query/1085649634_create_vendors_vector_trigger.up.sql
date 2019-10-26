CREATE
OR REPLACE VIEW autocomplete AS
SELECT
  vendors.name as query,
  vendors.tsv,
  'vendors' as type
FROM vendors
UNION
SELECT
  services.name as query,
  services.tsv,
  'services' as type
FROM services
UNION
SELECT
  categories.name as query,
  categories.tsv,
  'category' as type
FROM categories;