ALTER TABLE
  services
ADD
  COLUMN description VARCHAR,
ADD
  COLUMN tsv tsvector;

CREATE
  OR REPLACE FUNCTION ts_vector_services_trigger() RETURNS trigger AS $$ BEGIN new.tsv := setweight(
    to_tsvector('pg_catalog.english', coalesce(new.name, '')),
    'A'
  ) || setweight(
    to_tsvector(
      'pg_catalog.english',
      coalesce(new.description, '')
    ),
    'D'
  );
RETURN new;
END;
$$ LANGUAGE plpgsql;
DROP TRIGGER IF EXISTS services_ts_vector ON services;
CREATE TRIGGER services_ts_vector BEFORE
INSERT
  OR
UPDATE
  ON services FOR EACH ROW EXECUTE PROCEDURE ts_vector_services_trigger();