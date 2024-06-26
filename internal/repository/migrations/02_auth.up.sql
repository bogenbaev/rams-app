CREATE TABLE IF NOT EXISTS users (
     id SERIAL PRIMARY KEY,
     full_name VARCHAR(255) DEFAULT '',
     login VARCHAR(20) NOT NULL CHECK (LENGTH(login) >= 3 AND LENGTH(login) <= 20),
     email VARCHAR(255) NOT NULL,
     password VARCHAR(255) NOT NULL,
     created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);