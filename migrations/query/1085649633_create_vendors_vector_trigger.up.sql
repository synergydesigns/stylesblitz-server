ALTER TABLE vendors
  ADD COLUMN tsv tsvector;

CREATE OR REPLACE FUNCTION 
  vendors_ts_trigger_func() RETURNS trigger AS $$
  begin
    new.tsv :=
      setweight(to_tsvector('pg_catalog.english', coalesce(new.name,'')), 'A') ||
      setweight(to_tsvector('pg_catalog.english', coalesce(new.description,'')), 'D');
    return new;
  end
$$ LANGUAGE plpgsql;

DROP
 TRIGGER IF EXISTS vendor_ts_vector ON services;
CREATE TRIGGER vendor_ts_vector BEFORE
INSERT
  OR
UPDATE
  ON vendors FOR EACH ROW EXECUTE PROCEDURE vendors_ts_trigger_func();
