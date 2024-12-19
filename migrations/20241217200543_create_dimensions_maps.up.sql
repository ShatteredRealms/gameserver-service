CREATE EXTENSION IF NOT EXISTS pgcrypto;
CREATE EXTENSION IF NOT EXISTS btree_gist;

CREATE TABLE dimensions (
  id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
  name TEXT CHECK (name <> ''),
  version TEXT CHECK (version <> ''),
  location TEXT CHECK (location <> ''),
  created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  deleted_at TIMESTAMP,
  EXCLUDE USING GIST (
    name WITH =
  ) WHERE (deleted_at IS NULL)
);

CREATE TABLE maps (
  id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
  name TEXT CHECK (name <> ''),
  map_path TEXT CHECK (map_path <> ''),
  created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  deleted_at TIMESTAMP,
  EXCLUDE USING GIST (
    name WITH =
  ) WHERE (deleted_at IS NULL)
);

CREATE TABLE dimension_maps (
  dimension_id UUID NOT NULL,
  map_id UUID NOT NULL,
  CONSTRAINT fk_dimension_id FOREIGN KEY (dimension_id) REFERENCES dimensions(id),
  CONSTRAINT fk_map_id FOREIGN KEY (map_id) REFERENCES maps(id),
  CONSTRAINT pk_dimension_map PRIMARY KEY (dimension_id, map_id)
);

