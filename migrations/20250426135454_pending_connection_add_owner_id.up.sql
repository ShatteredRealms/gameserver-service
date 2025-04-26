ALTER TABLE pending_connections ADD COLUMN owner_id UUID UNIQUE NOT NULL;
