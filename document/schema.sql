CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

-- Store git repository data
CREATE TABLE repository (
  id UUID PRIMARY KEY DEFAULT uuid_generate_v1(),
  name VARCHAR(100) NOT NULL, -- User's custom name for the repository
  url VARCHAR(255) NOT NULL, -- Repository URL (for cloning, etc.)
  created_at TIMESTAMP(3),
  updated_at TIMESTAMP(3),
  deleted_at TIMESTAMP(3) -- Use when delete repository for soft delete, handled automatically by GORM
);
CREATE INDEX deleted_at_index ON repository (deleted_at);

-- Store scan status historically
CREATE TABLE scan_history (
  id SERIAL PRIMARY KEY,
  repository_id UUID NOT NULL, -- ID of repository table
  scan_id UUID NOT NULL, -- Generated ID to keep track of a single scan; also ID of scan_result table
  status VARCHAR(11) NOT NULL, -- Scan status at the time
  created_at TIMESTAMP(6)
);
CREATE INDEX repository_id_index ON scan_history (repository_id);

-- Store scan result historically
CREATE TABLE scan_result (
  id UUID PRIMARY KEY, -- Generated ID from scan_history (scan_id)
  result TEXT, -- Scan result, either in JSON format or just an error message
  created_at TIMESTAMP(6)
);

-- Store error message if there's an unexpected error in consumer task
CREATE TABLE task_error_log (
  id SERIAL PRIMARY KEY,
  body TEXT, -- Message body received from message queue
  message VARCHAR(255), -- Error message
  created_at TIMESTAMP(6)
);
