-- Initialize Yaro Wora Tourism Database
-- This script runs when the PostgreSQL container starts for the first time

-- Create the database if it doesn't exist (handled by docker-compose environment)
-- No additional initialization needed as GORM handles schema creation

-- Set timezone
SET timezone = 'Asia/Jakarta';

-- Create extensions that might be useful
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";
CREATE EXTENSION IF NOT EXISTS "citext";

-- Log initialization
SELECT 'Yaro Wora database initialized successfully!' as message;

-- Enforce at most one featured destination with a partial unique index
DO $$
BEGIN
    IF NOT EXISTS (
        SELECT 1
        FROM pg_indexes
        WHERE schemaname = 'public'
          AND indexname = 'uniq_destinations_featured_true'
    ) THEN
        EXECUTE 'CREATE UNIQUE INDEX uniq_destinations_featured_true ON destinations ((is_featured)) WHERE is_featured = true';
    END IF;
END $$;

