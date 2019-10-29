ALTER TABLE categories
  ADD COLUMN tsv tsvector;

CREATE OR REPLACE FUNCTION 
  categories_ts_trigger_func() RETURNS trigger AS $$
  begin
    new.tsv :=
      setweight(to_tsvector('pg_catalog.english', coalesce(new.name,'')), 'A') ||
      setweight(to_tsvector('pg_catalog.english', coalesce(new.description,'')), 'D');
    return new;
  end
$$ LANGUAGE plpgsql;

DROP
 TRIGGER IF EXISTS category_ts_vector ON services;
CREATE TRIGGER category_ts_vector BEFORE
INSERT
  OR
UPDATE
  ON categories FOR EACH ROW EXECUTE PROCEDURE categories_ts_trigger_func();

CREATE INDEX IF NOT EXISTS index_tsv_categories ON categories USING gin(tsv);
