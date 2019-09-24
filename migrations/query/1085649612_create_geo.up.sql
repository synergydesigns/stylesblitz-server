CREATE EXTENSION IF NOT EXISTS postgis;

CREATE EXTENSION IF NOT EXISTS  postgis_topology;
-- Enable PostGIS Advanced 3D
-- and other geoprocessing algorithms
-- sfcgal not available with all distributions
-- CREATE EXTENSION IF NOT EXISTS postgis_sfcgal;
-- fuzzy matching needed for Tiger
CREATE EXTENSION IF NOT EXISTS  fuzzystrmatch;
-- rule based standardizer
-- CREATE EXTENSION IF NOT EXISTS address_standardizer;
-- example rule data set
-- CREATE EXTENSION IF NOT EXISTS  address_standardizer_data_us;
-- Enable US Tiger Geocoder
CREATE EXTENSION IF NOT EXISTS postgis_tiger_geocoder;

-- geometry table is faster but harder to query and the result tends to be less accurate
-- we will stick to geography until speed becomes an issue
-- https://medium.com/coord/postgis-performance-showdown-geometry-vs-geography-ec99967da4f0
-- ALTER TABLE address ADD COLUMN geom geometry(Point, 4326);
-- UPDATE address SET geom = ST_SetSRID(ST_MakePoint(longitude, latitude), 4326);