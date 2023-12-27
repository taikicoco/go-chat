-- +goose Up
CREATE TABLE IF NOT EXISTS users (
    id uint PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
);
