CREATE TABLE users(
    id serial PRIMARY KEY,
    username VARCHAR(30) UNIQUE NOT NULL,
    password VARCHAR(100) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP

);

-- CREATE TABLE token(
--     id serial PRIMARY KEY,
--     token VARCHAR(100),
--     username VARCHAR(30) UNIQUE
-- );

