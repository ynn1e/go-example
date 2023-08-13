CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    name TEXT NOT NULL,
    age INTEGER DEFAULT 0
);

INSERT INTO users (name, age) VALUES ('Alice', 28);
INSERT INTO users (name, age) VALUES ('Bob', 32);
