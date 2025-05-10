-- Create the schema if it doesn't exist
CREATE SCHEMA IF NOT EXISTS users;

-- Create the users table
CREATE TABLE IF NOT EXISTS users.users (
    id BIGSERIAL PRIMARY KEY,
    name VARCHAR(255),
    phone VARCHAR(20),
    email VARCHAR(255) UNIQUE NOT NULL,
    role VARCHAR(255) NOT NULL,
    password_hash TEXT NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    is_deleted BOOLEAN NOT NULL DEFAULT FALSE
);

-- Create the sessions table
CREATE TABLE IF NOT EXISTS users.sessions (
    id BIGSERIAL PRIMARY KEY,
    user_id BIGINT NOT NULL REFERENCES users.users(id) ON DELETE CASCADE,
    refresh_token TEXT NOT NULL UNIQUE,
    expires_at TIMESTAMP NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);

-- Add an index for faster lookups on email in the users table
CREATE INDEX IF NOT EXISTS idx_users_email ON users.users (email);

-- Add an index for faster lookups on refresh_token in the sessions table
CREATE INDEX IF NOT EXISTS idx_sessions_refresh_token ON users.sessions (refresh_token);