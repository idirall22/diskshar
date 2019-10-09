DROP TABLE IF EXISTS users;

CREATE TABLE IF NOT EXISTS users(
    id SERIAL PRIMARY KEY,
    username VARCHAR UNIQUE NOT NULL,
    first_name VARCHAR,
    last_name VARCHAR,
    Email VARCHAR UNIQUE NOT NULL,
    password VARCHAR NOT NULL,
    avatar VARCHAR,
    created_at TIMESTAMP with TIME ZONE DEFAULT now(),
    deleted_at TIMESTAMP DEFAULT NULL
);