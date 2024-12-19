CREATE TABLE pending_connections (
  id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
  character_id UUID NOT NULL,
  server_name TEXT CHECK (server_name <> ''),
  created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);
