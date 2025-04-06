-- +goose Up
CREATE TABLE feeds (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(), -- Use UUID as primary key, automatically generate a value on insert
    created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP, -- Timestamp with timezone, cannot be null, defaults to the time of insertion
    updated_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP, -- Timestamp with timezone, cannot be null, defaults to the time of insertion (see note below)
    name VARCHAR(255) NOT NULL, -- String (up to 255 chars), must be unique and cannot be null
    url VARCHAR(255) UNIQUE NOT NULL, -- String (up to 255 chars), must be unique and cannot be null
    user_id VARCHAR(255) NOT NULL REFERENCES users (name) ON DELETE CASCADE -- String (up to 255 chars), must be unique and cannot be null
);
-- +goose Down
DROP TABLE feeds;