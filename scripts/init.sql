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
