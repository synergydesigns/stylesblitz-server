CREATE
OR REPLACE FUNCTION fn_address_table_geo_update_event() RETURNS TRIGGER AS $fn_address_table_geo_update_event$ 
BEGIN
  NEW.geog := ST_SetSRID(ST_MakePoint(NEW.longitude, NEW.latitude), 4326)::geography;RAISE NOTICE 'UPDATING geo data for %, [%,%]',
  NEW.id,
  NEW.latitude,
  NEW.longitude;
RETURN NEW;END;$fn_address_table_geo_update_event$ LANGUAGE plpgsql;


  -- trigger for updating geo if either longitude of latitude when and address is created
DROP 
  TRIGGER IF EXISTS fn_address_table_geo_inserted ON address;
CREATE TRIGGER fn_address_table_geo_inserted 
BEFORE 
  INSERT ON address FOR EACH ROW EXECUTE PROCEDURE fn_address_table_geo_update_event();


-- trigger for updating geo if either longitude of latitude of an address changes
DROP 
  TRIGGER IF EXISTS fn_address_table_geo_update_event ON address;
CREATE TRIGGER fn_address_table_geo_updated 
BEFORE 
UPDATE
  OF latitude, 
  longitude ON address FOR EACH ROW EXECUTE PROCEDURE fn_address_table_geo_update_event();
