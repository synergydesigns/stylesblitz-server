-- stored function to update address geo location.
CREATE OR REPLACE FUNCTION fn_address_table_geo_update_event() RETURNS trigger AS $fn_address_table_geo_update_event$
 	BEGIN  
     	UPDATE address SET 
     	geog = ST_SetSRID(ST_MakePoint(NEW.longitude,NEW.latitude), 4326)::geography;
	
     	RAISE NOTICE 'UPDATING geo data for %, [%,%]',
		NEW.id,
		NEW.latitude,
		NEW.longitude; 
	   	RETURN NULL; -- result is ignored since this is an AFTER trigger
 	END;
$fn_address_table_geo_update_event$ LANGUAGE plpgsql;


-- trigger for updating geo if either longitude of latitude when and address is created
DROP 
  TRIGGER IF EXISTS fn_address_table_geo_inserted ON address;
CREATE TRIGGER fn_address_table_geo_inserted 
AFTER 
  INSERT ON address FOR EACH ROW EXECUTE PROCEDURE fn_address_table_geo_update_event();


-- trigger for updating geo if either longitude of latitude of an address changes
DROP 
  TRIGGER IF EXISTS fn_address_table_geo_update_event ON address;
CREATE TRIGGER fn_address_table_geo_updated 
AFTER 
UPDATE
  OF latitude, 
  longitude ON address FOR EACH ROW EXECUTE PROCEDURE fn_address_table_geo_update_event();
