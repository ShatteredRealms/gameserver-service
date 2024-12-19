CREATE OR REPLACE FUNCTION timestamp_updated_column()
RETURNS TRIGGER AS $$
BEGIN
  NEW.updated_at = now();
  RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE TRIGGER dimensions_updated_at BEFORE UPDATE 
ON dimensions FOR EACH ROW EXECUTE PROCEDURE
timestamp_updated_column();

CREATE TRIGGER maps_updated_at BEFORE UPDATE 
ON maps FOR EACH ROW EXECUTE PROCEDURE
timestamp_updated_column();
